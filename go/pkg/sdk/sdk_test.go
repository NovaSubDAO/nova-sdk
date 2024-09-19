package sdk

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/NovaSubDAO/nova-sdk/go/pkg/constants"
	ethereumContracts "github.com/NovaSubDAO/nova-sdk/go/pkg/sdk/ethereum/abis"
	optimismContracts "github.com/NovaSubDAO/nova-sdk/go/pkg/sdk/optimism/abis"
	"github.com/NovaSubDAO/nova-sdk/go/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	rpcEndpoint    string
	chainId        int64
	vault          string
	vaultAddress   common.Address
	vaultDecimals  int64
	initialAdapter common.Address
	stable         constants.Stablecoin
	stableAddress  common.Address
	savings        common.Address
	testAddress    common.Address
	privateKeyHex  string
}

var testCases = []testCase{
	{
		rpcEndpoint:    os.Getenv("ETH_RPC_ENDPOINT"),
		chainId:        1,
		vault:          constants.ConfigDetails[1].VaultAddress,
		vaultAddress:   common.HexToAddress(constants.ConfigDetails[1].VaultAddress),
		vaultDecimals:  constants.ConfigDetails[1].VaultDecimals,
		initialAdapter: getActiveAdapter(1),
		stable:         constants.DAI,
		stableAddress:  common.HexToAddress(constants.StablecoinDetails[1][constants.DAI].Address),
		savings:        common.HexToAddress(constants.ConfigDetails[1].Savings),
		testAddress:    common.HexToAddress(os.Getenv("TEST_ADDRESS")),
		privateKeyHex:  os.Getenv("TEST_PRIVATE_KEY"),
	},
	{
		rpcEndpoint:    os.Getenv("OPT_RPC_ENDPOINT"),
		chainId:        10,
		vault:          constants.ConfigDetails[10].VaultAddress,
		vaultAddress:   common.HexToAddress(constants.ConfigDetails[10].VaultAddress),
		vaultDecimals:  constants.ConfigDetails[10].VaultDecimals,
		initialAdapter: getActiveAdapter(10),
		stable:         constants.USDC,
		stableAddress:  common.HexToAddress(constants.StablecoinDetails[10][constants.USDC].Address),
		savings:        common.HexToAddress(constants.ConfigDetails[10].Savings),
		testAddress:    common.HexToAddress(os.Getenv("TEST_ADDRESS")),
		privateKeyHex:  os.Getenv("TEST_PRIVATE_KEY"),
	},
}

func TestSdkConfig(t *testing.T) {
	for _, tc := range testCases {
		t.Run("ChainID: "+strconv.FormatInt(tc.chainId, 10), func(t *testing.T) {
			novaSdk, err := NewNovaSDK(tc.rpcEndpoint, tc.chainId)
			assert.NoError(t, err)

			cfg := novaSdk.GetConfig()
			assert.Equal(t, tc.chainId, cfg.ChainId)
		})
	}
}

func TestSdkConfigDetails(t *testing.T) {
	for _, tc := range testCases {
		t.Run("ChainID: "+strconv.FormatInt(tc.chainId, 10), func(t *testing.T) {
			novaSdk, err := NewNovaSDK(tc.rpcEndpoint, tc.chainId)
			assert.NoError(t, err)

			cfg := novaSdk.GetConfig()

			vault := constants.ConfigDetails[tc.chainId].VaultAddress
			assert.Equal(t, vault, cfg.VaultAddress)

			savings := constants.ConfigDetails[tc.chainId].Savings
			assert.Equal(t, savings, cfg.Savings)

			vaultDecimals := constants.ConfigDetails[tc.chainId].VaultDecimals
			assert.Equal(t, vaultDecimals, cfg.VaultDecimals)
		})
	}
}

func TestSdkCreateDepositTx(t *testing.T) {
	for _, tc := range testCases {
		t.Run("ChainID: "+strconv.FormatInt(tc.chainId, 10), func(t *testing.T) {
			if tc.chainId == 1 {
				t.Skip("Skipping test for chain ID 1")
				return
			}
			novaSdk, err := NewNovaSDK(tc.rpcEndpoint, tc.chainId)
			assert.NoError(t, err)

			client, err := ethclient.Dial(tc.rpcEndpoint)
			if err != nil {
				log.Fatalf("Failed to connect to the client: %v", err)
			}
			privateKey, err := crypto.HexToECDSA(tc.privateKeyHex)
			if err != nil {
				log.Fatalf("Failed to parse private key: %v", err)
			}

			mockAmount := big.NewInt(1e6)
			mockReferral := big.NewInt(123)

			fmt.Println()
			fmt.Println("------------------------------ Increasing allowance... ------------------------------")
			printUserEthBalance(tc.testAddress, client)
			_, err = increaseAllowance(tc.testAddress, tc.vaultAddress, tc.stableAddress, mockAmount, privateKey, tc.rpcEndpoint, tc.chainId)
			if err != nil {
				log.Fatalf("TestSdkCreateDepositTx: increaseAllowance: %v", err)
			}
			fmt.Println("******************************** Allowance increased ********************************")

			initialBalanceStable, err := getAssetUserBalance(tc.testAddress, tc.stableAddress, client, tc.chainId)
			if err != nil {
				log.Fatalf("Failed to get user's stable balance: %v", err)
			}
			fmt.Println("Initial stable balance: ", initialBalanceStable)
			initialBalanceSavings, err := getAssetUserBalance(tc.testAddress, tc.savings, client, tc.chainId)
			if err != nil {
				log.Fatalf("Failed to get user's savings balance: %v", err)
			}

			fmt.Println("Initial savings balance: ", initialBalanceSavings)

			fmt.Println()
			fmt.Println("-------------------------- Creating deposit transaction... --------------------------")
			printUserEthBalance(tc.testAddress, client)
			txJSON, err := novaSdk.CreateDepositTransaction(tc.stable, tc.testAddress, mockAmount, mockReferral)
			if err != nil {
				log.Fatalf("Failed to call 'CreateDepositTransaction' function: %s", err)
			}
			fmt.Println("- Transaction created")

			var tx types.Transaction
			err = json.Unmarshal([]byte(txJSON), &tx)
			if err != nil {
				log.Fatalf("Failed to unmarshal transaction: %v", err)
			}

			signedTx, err := utils.SignTransaction(&tx, big.NewInt(tc.chainId), privateKey)
			if err != nil {
				log.Fatalf("TestSdkCreateDepositTx: failed to sign transaction:%v", err)
			}
			fmt.Println("-- Transaction signed")

			err = utils.SendTransaction(signedTx, tc.rpcEndpoint)
			if err != nil {
				log.Fatalf("TestSdkCreateDepositTx: Failed to send transaction: %v", err)
			}
			fmt.Println("---- Transaction sent")
			fmt.Println("********************************** Deposit success **********************************")

			newBalanceStable, _ := getAssetUserBalance(tc.testAddress, tc.stableAddress, client, tc.chainId)
			expectedBalanceStable := new(big.Int).Sub(initialBalanceStable, mockAmount)

			newBalanceSavings, err := getAssetUserBalance(tc.testAddress, tc.savings, client, tc.chainId)
			if err != nil {
				log.Fatalf("Failed to get user's savings balance: %v", err)
			}
			receivedSavings := new(big.Int).Sub(newBalanceSavings, initialBalanceSavings)
			fmt.Println("Received savings: ", receivedSavings)

			_, _, executedPrice, _ := novaSdk.GetSlippage(tc.stable, mockAmount)
			mockAmountFloat := new(big.Float).SetInt(mockAmount)
			executedPriceFloat := new(big.Float).SetFloat64(executedPrice)
			reciprocalPrice := new(big.Float).Quo(big.NewFloat(1), executedPriceFloat)
			expectedReceivedSavingsFloat := new(big.Float).Mul(mockAmountFloat, reciprocalPrice)

			savingsDecimals := tc.vaultDecimals
			stableDecimals := constants.StablecoinDetails[tc.chainId][tc.stable].Decimals
			decimalsDifference := savingsDecimals - int64(stableDecimals)
			scaleFactor := new(big.Float).SetFloat64(math.Pow(10, float64(decimalsDifference)))
			scaledExpectedReceivedSavings := new(big.Float).Mul(expectedReceivedSavingsFloat, scaleFactor)

			roundedExpectedReceivedSavings := new(big.Int)
			scaledExpectedReceivedSavings.Int(roundedExpectedReceivedSavings)

			expectedReceivedSavings6Decimals := scaleByFactor(roundedExpectedReceivedSavings, scaleFactor)
			receivedSavings6decimals := scaleByFactor(receivedSavings, scaleFactor)

			assert.Equal(t, expectedBalanceStable, newBalanceStable, "Stable balance should be equal to the initial balance minus the mock amount")
			assert.Equal(t, expectedReceivedSavings6Decimals, receivedSavings6decimals, "Wrong sDai received amount")
		})
	}
}

func TestSdkCreateWithdrawTx(t *testing.T) {
	for _, tc := range testCases {
		t.Run("ChainID: "+strconv.FormatInt(tc.chainId, 10), func(t *testing.T) {
			if tc.chainId == 1 {
				t.Skip("Skipping test for chain ID 1")
				return
			}
			novaSdk, err := NewNovaSDK(tc.rpcEndpoint, tc.chainId)
			assert.NoError(t, err)

			client, err := ethclient.Dial(tc.rpcEndpoint)
			if err != nil {
				log.Fatalf("Failed to connect to the client: %v", err)
			}
			privateKey, err := crypto.HexToECDSA(tc.privateKeyHex)
			if err != nil {
				log.Fatalf("Failed to parse private key: %v", err)
			}

			mockAmount := big.NewInt(6e17)
			mockReferral := big.NewInt(123)

			initialBalanceSavings, err := getAssetUserBalance(tc.testAddress, tc.savings, client, tc.chainId)
			if err != nil {
				log.Fatalf("Failed to get user's savings balance: %v", err)
			}
			fmt.Println("Initial balance savings: ", initialBalanceSavings)
			initialBalanceStable, err := getAssetUserBalance(tc.testAddress, tc.stableAddress, client, tc.chainId)
			if err != nil {
				log.Fatalf("Failed to get user's savings balance: %v", err)
			}
			fmt.Println("Initial balance stable: ", initialBalanceStable)

			fmt.Println()
			fmt.Println("------------------------------ Increasing allowance... ------------------------------")
			printUserEthBalance(tc.testAddress, client)
			_, err = increaseAllowance(tc.testAddress, tc.vaultAddress, tc.savings, mockAmount, privateKey, tc.rpcEndpoint, tc.chainId)
			if err != nil {
				log.Fatalf("Failed to increase allowance: %v", err)
			}
			fmt.Println("******************************** Allowance increased ********************************")

			fmt.Println()
			fmt.Println("------------------------- Creating withdraw transaction... --------------------------")
			printUserEthBalance(tc.testAddress, client)
			txJSON, err := novaSdk.CreateWithdrawTransaction(tc.stable, tc.testAddress, mockAmount, mockReferral)
			if err != nil {
				log.Fatalf("Failed to call 'CreateWithdrawTransaction' function: %s", err)
			}
			fmt.Println("- Transaction created")

			var tx types.Transaction
			err = json.Unmarshal([]byte(txJSON), &tx)
			if err != nil {
				log.Fatalf("Failed to unmarshal transaction: %v", err)
			}

			signedTx, err := utils.SignTransaction(&tx, big.NewInt(tc.chainId), privateKey)
			if err != nil {
				log.Fatalf("TestSdkCreateWithdrawTx: failed to sign transaction: %v", err)
			}
			fmt.Println("-- Transaction signed")

			err = utils.SendTransaction(signedTx, tc.rpcEndpoint)
			if err != nil {
				log.Fatalf("TestSdkCreateWithdrawTx: failed to send transaction: %v", err)
			}
			fmt.Println("---- Transaction sent")
			fmt.Println("********************************** Withdraw success **********************************")

			newBalanceSavings, err := getAssetUserBalance(tc.testAddress, tc.savings, client, tc.chainId)
			if err != nil {
				log.Fatalf("Failed to get user's savings balance: %v", err)
			}
			fmt.Println("New balance savings: ", newBalanceSavings)
			newBalanceStable, err := getAssetUserBalance(tc.testAddress, tc.stableAddress, client, tc.chainId)
			if err != nil {
				log.Fatalf("Failed to get user's savings balance: %v", err)
			}
			fmt.Println("New balance stable: ", newBalanceStable)

			expectedBalanceSavings := new(big.Int).Sub(initialBalanceSavings, mockAmount)
			receivedStable := new(big.Int).Sub(newBalanceStable, initialBalanceStable)

			price, _ := novaSdk.GetPrice(tc.stable)

			savingsDecimals := tc.vaultDecimals
			stableDecimals := constants.StablecoinDetails[tc.chainId][tc.stable].Decimals

			decimalsDifference := savingsDecimals - int64(stableDecimals)
			decimalMultiplier := new(big.Int).Exp(big.NewInt(10), big.NewInt(decimalsDifference), nil)

			priceAdjustedToDecimals := new(big.Int).Mul(price, decimalMultiplier)

			priceTimesMockAmount := new(big.Int).Mul(priceAdjustedToDecimals, mockAmount)

			fixedDecimalDivisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(30), nil)
			expectedReceivedStable := new(big.Int).Div(priceTimesMockAmount, fixedDecimalDivisor)

			receivedStableFloat := new(big.Float).SetInt(receivedStable)
			expectedReceivedStableFloat := new(big.Float).SetInt(expectedReceivedStable)

			stableDecimalsInt := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(stableDecimals)), nil)
			stableDecimalsFloat := new(big.Float).SetInt(stableDecimalsInt)
			receivedStableInCorrectDecimals := new(big.Float).Quo(receivedStableFloat, stableDecimalsFloat)
			expectedReceivedStableInCorrectDecimals := new(big.Float).Quo(expectedReceivedStableFloat, stableDecimalsFloat)

			tolerance := 1e-3
			receivedStableFloat64, _ := receivedStableInCorrectDecimals.Float64()
			expectedStableFloat64, _ := expectedReceivedStableInCorrectDecimals.Float64()

			assert.InDelta(t, expectedStableFloat64, receivedStableFloat64, tolerance, "Wrong stable received amount")
			assert.Equal(t, newBalanceSavings, expectedBalanceSavings, "Savings balance should be equal to the initial balance minus the mock amount")
		})
	}
}

func TestReplaceAdapter(t *testing.T) {
	for _, tc := range testCases {
		t.Run("ChainID: "+strconv.FormatInt(tc.chainId, 10), func(t *testing.T) {
			if tc.chainId == 1 {
				t.Skip("Skipping test for chain ID 1")
				return
			}

			privateKey, err := crypto.HexToECDSA(tc.privateKeyHex)
			if err != nil {
				log.Fatalf("Failed to parse private key: %v", err)
			}

			newAdapter := common.HexToAddress("0xe20aB0BaaF8d581f48FECC84dC32be076C223860")

			activeAdapter := getActiveAdapter(tc.chainId)
			assert.Equal(t, activeAdapter, tc.initialAdapter, "The active adapter should match the initial adapter before invoking the 'replaceAdapter' function")

			replaceAdapter(tc.testAddress, tc.vaultAddress, tc.stableAddress, newAdapter, privateKey, tc.rpcEndpoint, tc.chainId)
			activeAdapter = getActiveAdapter(tc.chainId)

			assert.Equal(t, activeAdapter, newAdapter, "The active adapter should match the new adapter after invoking the 'replaceAdapter' function")

			replaceAdapter(tc.testAddress, tc.vaultAddress, tc.stableAddress, tc.initialAdapter, privateKey, tc.rpcEndpoint, tc.chainId)
			activeAdapter = getActiveAdapter(tc.chainId)

			assert.Equal(t, activeAdapter, tc.initialAdapter, "The active adapter should correspond to the initial adapter before invoking the 'TestSdkCreateDepositTx' and 'TestSdkCreateWithdrawTx' functions")
			TestSdkCreateDepositTx(t)
			TestSdkCreateWithdrawTx(t)
		})
	}
}

func TestSdkGetSavingsPrice(t *testing.T) {
	for _, tc := range testCases {
		t.Run("ChainID: "+strconv.FormatInt(tc.chainId, 10), func(t *testing.T) {
			novaSdk, err := NewNovaSDK(tc.rpcEndpoint, tc.chainId)
			assert.NoError(t, err)

			price, err := novaSdk.GetSDaiPrice()
			assert.NoError(t, err)

			lowerBound := big.NewInt(1e18)
			upperBound := big.NewInt(2e18)

			assert.True(t, price.Cmp(lowerBound) >= 0, "Price should be at least 1e18")
			assert.True(t, price.Cmp(upperBound) <= 0, "Price should be at most 2e18")
		})
	}
}

func TestSdkGetSlippage(t *testing.T) {
	for _, tc := range testCases {
		t.Run("ChainID: "+strconv.FormatInt(tc.chainId, 10), func(t *testing.T) {
			novaSdk, err := NewNovaSDK(tc.rpcEndpoint, tc.chainId)
			assert.NoError(t, err)

			stableDecimals := constants.StablecoinDetails[tc.chainId][tc.stable].Decimals
			base := big.NewInt(10)
			one := new(big.Int).Exp(base, big.NewInt(int64(stableDecimals)), nil)

			slippageAmount := new(big.Int).Mul(big.NewInt(1000), one)
			slippage, expectedPrice, executedPrice, err := novaSdk.GetSlippage(tc.stable, slippageAmount)
			assert.NoError(t, err)

			slippageLowerBound := 0.0
			slippageUpperBound := 1.0

			assert.True(t, slippage >= slippageLowerBound, "slippage should be at least 0")
			assert.True(t, slippage <= slippageUpperBound, "slippage should be at most 1")

			priceLowerBound := 1.0
			priceUpperBound := 2.0

			assert.True(t, expectedPrice >= priceLowerBound, "expectedPrice should be at least 1e18")
			assert.True(t, expectedPrice <= priceUpperBound, "expectedPrice should be at most 2e18")

			assert.True(t, executedPrice >= priceLowerBound, "executedPrice should be at least 1e18")
			assert.True(t, executedPrice <= priceUpperBound, "executedPrice should be at most 2e18")
		})
	}
}

func TestSdkGetPosition(t *testing.T) {
	for _, tc := range testCases {
		t.Run("ChainID: "+strconv.FormatInt(tc.chainId, 10), func(t *testing.T) {
			novaSdk, err := NewNovaSDK(tc.rpcEndpoint, tc.chainId)
			assert.NoError(t, err)

			position, err := novaSdk.GetPosition(tc.stable, tc.testAddress)
			assert.NoError(t, err)

			lowerBound := big.NewInt(0)

			assert.True(t, position.Cmp(lowerBound) >= 0, "Position should be at least 0")
		})
	}
}

func TestSdkGetSavingsTotalValue(t *testing.T) {
	for _, tc := range testCases {
		t.Run("ChainID: "+strconv.FormatInt(tc.chainId, 10), func(t *testing.T) {
			novaSdk, err := NewNovaSDK(tc.rpcEndpoint, tc.chainId)
			assert.NoError(t, err)

			position, err := novaSdk.GetSDaiTotalValue()
			assert.NoError(t, err)

			lowerBound := big.NewInt(0)

			assert.True(t, position.Cmp(lowerBound) >= 0, "Position should be at least 0")
		})
	}
}

func TestSdkGetSupportedStablecoins(t *testing.T) {
	for _, tc := range testCases {
		t.Run("ChainID: "+strconv.FormatInt(tc.chainId, 10), func(t *testing.T) {
			novaSdk, err := NewNovaSDK(tc.rpcEndpoint, tc.chainId)
			assert.NoError(t, err)

			stables, err := novaSdk.GetSupportedStablecoins()
			assert.NoError(t, err)

			assert.Len(t, stables, 1, "The length of the stablecoins array should be 1")
			assert.Equal(t, tc.stable, stables[0], "The first stablecoin is not correct")
		})
	}
}

func replaceAdapter(from common.Address, vaultAddress common.Address, stable common.Address, adapter common.Address, privateKey *ecdsa.PrivateKey, rpcEndpoint string, chainId int64) {
	novaVaultAbi, err := abi.JSON(strings.NewReader(optimismContracts.NovaVaultABI))
	if err != nil {
		log.Fatalf("failed to parse contract ABI: %v", err)
	}

	data, err := novaVaultAbi.Pack("replaceAdapter", stable, adapter)
	if err != nil {
		log.Fatalf("ABI pack failed: %v", err)
	}

	tx, err := utils.CreateTransaction(from, vaultAddress, data, rpcEndpoint)
	if err != nil {
		log.Fatalf("Replace adapter failed: failed to create transaction: %v", err)
	}

	fmt.Println("- Transaction created")

	signedTx, err := utils.SignTransaction(tx, big.NewInt(chainId), privateKey)
	if err != nil {
		log.Fatalf("Replace adapter failed: failed to sign transaction: %v", err)
	}
	fmt.Println("-- Transaction signed")

	err = utils.SendTransaction(signedTx, rpcEndpoint)
	if err != nil {
		log.Fatalf("Replace adapter failed: failed to send transaction: %v", err)
	}
	fmt.Println("---- Transaction sent")
}

func printUserEthBalance(from common.Address, client *ethclient.Client) {
	balance, err := client.BalanceAt(context.Background(), from, nil)
	if err != nil {
		log.Fatalf("Failed to get the balance: %v", err)
	}
	ethBalance := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))
	fmt.Printf("Eth balance of %s: %f ETH\n", from.Hex(), ethBalance)
}

func getAssetUserBalance(from common.Address, asset common.Address, client *ethclient.Client, chainId int64) (*big.Int, error) {
	var balanceAsset *big.Int

	if chainId == 1 {
		if asset == common.HexToAddress(constants.StablecoinDetails[chainId][constants.DAI].Address) {
			assetContract, err := ethereumContracts.NewDaiCaller(asset, client)
			if err != nil {
				return nil, fmt.Errorf("failed to load Dai contract: %w", err)
			}
			balanceAsset, err = assetContract.BalanceOf(nil, from)
			if err != nil {
				return nil, fmt.Errorf("failed to get Dai user's balance: %v", err)
			}
		} else if asset == common.HexToAddress(constants.ConfigDetails[chainId].Savings) {
			assetContract, err := ethereumContracts.NewSavingsDaiCaller(asset, client)
			if err != nil {
				return nil, fmt.Errorf("failed to load SavingsDai contract: %w", err)
			}
			balanceAsset, err = assetContract.BalanceOf(nil, from)
			if err != nil {
				return nil, fmt.Errorf("failed to get savings user's balance: %v", err)
			}
		} else {
			return nil, fmt.Errorf("asset address not compatible")
		}
	} else if chainId == 10 {
		if asset == common.HexToAddress(constants.StablecoinDetails[chainId][constants.USDC].Address) {
			assetContract, err := optimismContracts.NewFiatTokenV22Caller(asset, client)
			if err != nil {
				return nil, fmt.Errorf("failed to load FiatTokenV22 contract: %w", err)
			}
			balanceAsset, err = assetContract.BalanceOf(nil, from)
			if err != nil {
				return nil, fmt.Errorf("failed to get stable user's balance: %v", err)
			}

		} else if asset == common.HexToAddress(constants.ConfigDetails[chainId].Savings) {
			assetContract, err := optimismContracts.NewSavingsDaiCaller(asset, client)
			if err != nil {
				return nil, fmt.Errorf("failed to load SavingsDai contract: %w", err)
			}
			balanceAsset, err = assetContract.BalanceOf(nil, from)
			if err != nil {
				return nil, fmt.Errorf("failed to get savings user's balance: %v", err)
			}
		} else {
			return nil, fmt.Errorf("asset address not compatible")
		}
	}
	return balanceAsset, nil
}

func increaseAllowance(from common.Address, spender common.Address, token common.Address, amount *big.Int, privateKey *ecdsa.PrivateKey, rpcEndpoint string, chainId int64) (*types.Transaction, error) {
	var asset abi.ABI
	var err error

	if chainId == 1 {
		if token == common.HexToAddress(constants.StablecoinDetails[chainId][constants.DAI].Address) {
			asset, err = abi.JSON(strings.NewReader(ethereumContracts.DaiABI))
			if err != nil {
				log.Fatalf("Failed to parse Dai ABI: %v", err)
			}
		} else if token == common.HexToAddress(constants.ConfigDetails[chainId].Savings) {
			asset, err = abi.JSON(strings.NewReader(ethereumContracts.SavingsDaiABI))
			if err != nil {
				log.Fatalf("Failed to parse savings ABI: %v", err)
			}
		} else {
			return nil, fmt.Errorf("asset address not compatible")
		}
	} else if chainId == 10 {
		if token == common.HexToAddress(constants.StablecoinDetails[chainId][constants.USDC].Address) {
			asset, err = abi.JSON(strings.NewReader(optimismContracts.FiatTokenV22ABI))
			if err != nil {
				log.Fatalf("Failed to parse FiatTokenV22 ABI: %v", err)
			}
		} else if token == common.HexToAddress(constants.ConfigDetails[chainId].Savings) {
			asset, err = abi.JSON(strings.NewReader(optimismContracts.SavingsDaiABI))
			if err != nil {
				log.Fatalf("Failed to parse savings ABI: %v", err)
			}
		} else {
			return nil, fmt.Errorf("asset address not compatible")
		}
	}

	data, err := asset.Pack("approve", spender, amount)
	if err != nil {
		log.Fatalf("Failed to pack data for transaction: %v", err)
	}

	tx, err := utils.CreateTransaction(from, token, data, rpcEndpoint)
	if err != nil {
		log.Fatalf("Increase allowance failed: failed to create transaction: %v", err)
	}
	fmt.Println("- Transaction created")

	signedTx, err := utils.SignTransaction(tx, big.NewInt(chainId), privateKey)
	if err != nil {
		log.Fatalf("Increase allowance failed: failed to sign transaction: %v", err)
	}
	fmt.Println("-- Transaction signed")

	err = utils.SendTransaction(signedTx, rpcEndpoint)
	if err != nil {
		log.Fatalf("Increase allowance failed: failed to send transaction: %v", err)
	}
	fmt.Println("---- Transaction sent")

	return signedTx, nil
}

func getActiveAdapter(chainId int64) common.Address {
	var novaVaultInstance *optimismContracts.NovaVault
	var vaultAddress common.Address
	var stableAddress common.Address

	if chainId == 1 {
		return common.Address{}
	}

	if chainId == 10 {
		rpcEndpoint := os.Getenv("OPT_RPC_ENDPOINT")

		client, err := ethclient.Dial(rpcEndpoint)
		if err != nil {
			log.Fatalf("Failed to connect to the client: %v", err)
		}

		vaultAddress = common.HexToAddress(constants.ConfigDetails[10].VaultAddress)
		stableAddress = common.HexToAddress(constants.StablecoinDetails[10][constants.USDC].Address)

		novaVaultInstance, err = optimismContracts.NewNovaVault(vaultAddress, client)
		if err != nil {
			log.Fatalf("Failed to create novaVaultInstance on Optimism: %v", err)
		}
	}

	activeAdapter, err := novaVaultInstance.NovaAdapters(&bind.CallOpts{Context: context.Background()}, stableAddress)
	if err != nil {
		log.Fatalf("Failed to get active adapter: %v", err)
	}

	return activeAdapter
}

func scaleByFactor(amount *big.Int, scaleFactor *big.Float) *big.Int {
	scaledAmount := new(big.Float).SetInt(amount)
	scaledAmount.Quo(scaledAmount, scaleFactor)

	result := new(big.Int)
	scaledAmount.Int(result)

	return result
}
