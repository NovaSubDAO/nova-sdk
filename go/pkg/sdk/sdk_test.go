package sdk

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	rpcEndpoint   string
	chainId       int64
	vaultAddress  string
	stable        constants.Stablecoin
	mockAddress   common.Address
	privateKeyHex string
}

var testCases = []testCase{
	{
		rpcEndpoint:   os.Getenv("ETH_RPC_ENDPOINT"),
		chainId:       1,
		vaultAddress:  constants.ConfigDetails[1].VaultAddress,
		stable:        constants.DAI,
		mockAddress:   common.HexToAddress(os.Getenv("TEST_ADDRESS")),
		privateKeyHex: os.Getenv("TEST_PRIVATE_KEY"),
	},
	{
		rpcEndpoint:   os.Getenv("OPT_RPC_ENDPOINT_ANVIL"),
		chainId:       10,
		vaultAddress:  constants.ConfigDetails[10].VaultAddress,
		stable:        constants.USDC,
		mockAddress:   common.HexToAddress(os.Getenv("TEST_ADDRESS")),
		privateKeyHex: os.Getenv("TEST_PRIVATE_KEY"),
	},
}

func printUserBalance(from common.Address, client *ethclient.Client) {
	balance, err := client.BalanceAt(context.Background(), from, nil)
	if err != nil {
		log.Fatalf("Failed to get the balance: %v", err)
	}
	ethBalance := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))
	fmt.Printf("Eth balance of %s: %f ETH\n", from.Hex(), ethBalance)
}

func increaseAllowance(from common.Address, spender common.Address, token common.Address, amount *big.Int, isDeposit bool, client *ethclient.Client, privateKey *ecdsa.PrivateKey, rpcEndpoint string, chainId int64) (*types.Transaction, error) {
	var asset abi.ABI
	var err error

	if chainId == 1 {
		if isDeposit {
			asset, err = abi.JSON(strings.NewReader(ethereumContracts.DaiABI))
			if err != nil {
				log.Fatalf("Failed to parse Dai ABI: %v", err)
			}
		} else {
			asset, err = abi.JSON(strings.NewReader(ethereumContracts.SavingsDaiABI))
			if err != nil {
				log.Fatalf("Failed to parse sDai ABI: %v", err)
			}
		}
	} else if chainId == 10 {
		if isDeposit {
			asset, err = abi.JSON(strings.NewReader(optimismContracts.FiatTokenV22ABI))
			if err != nil {
				log.Fatalf("Failed to parse FiatTokenV22 ABI: %v", err)
			}
		} else {
			asset, err = abi.JSON(strings.NewReader(optimismContracts.SavingsDaiABI))
			if err != nil {
				log.Fatalf("Failed to parse sDai ABI: %v", err)
			}
		}
	}

	data, err := asset.Pack("approve", spender, amount)
	if err != nil {
		log.Fatalf("Failed to pack data for transaction: %v", err)
	}

	tx, err := utils.CreateTransaction(from, token, data, rpcEndpoint)
	if err != nil {
		log.Fatalf("Increase allowance failed: %v", err)
	}
	fmt.Println("- Transaction created")

	signedTx, err := utils.SignTransaction(tx, big.NewInt(chainId), privateKey)
	if err != nil {
		log.Fatalf("Increase allowance failed: %v", err)
	}
	fmt.Println("-- Transaction signed")

	err = utils.SendTransaction(signedTx, rpcEndpoint)
	if err != nil {
		log.Fatalf("Increase allowance failed: %v", err)
	}
	fmt.Println("---- Transaction sent")

	return signedTx, nil
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

			vaultAddress := constants.ConfigDetails[tc.chainId].VaultAddress
			assert.Equal(t, vaultAddress, cfg.VaultAddress)

			sDai := constants.ConfigDetails[tc.chainId].SDai
			assert.Equal(t, sDai, cfg.SDai)

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

			mockAmount := big.NewInt(1e4)
			mockReferral := big.NewInt(123)
			spender := common.HexToAddress(tc.vaultAddress)
			client, err := ethclient.Dial(tc.rpcEndpoint)
			if err != nil {
				log.Fatalf("Failed to connect to the Optimism client: %v", err)
			}
			privateKey, err := crypto.HexToECDSA(tc.privateKeyHex)
			if err != nil {
				log.Fatalf("Failed to parse private key: %v", err)
			}

			stableAddress := common.HexToAddress(constants.StablecoinDetails[tc.chainId][tc.stable].Address)

			fmt.Println()
			fmt.Println("------------------------------ Increasing allowance... ------------------------------")
			printUserBalance(tc.mockAddress, client)
			_, err = increaseAllowance(tc.mockAddress, spender, stableAddress, mockAmount, true, client, privateKey, tc.rpcEndpoint, tc.chainId)
			if err != nil {
				log.Fatalf("Failed to increase allowance: %v", err)
			}
			fmt.Println("******************************** Allowance increased ********************************")

			fmt.Println()
			fmt.Println("-------------------------- Creating deposit transaction... --------------------------")
			printUserBalance(tc.mockAddress, client)
			txJSON, err := novaSdk.CreateDepositTransaction(tc.stable, tc.mockAddress, mockAmount, mockReferral)
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
				log.Fatalf("TestSdkCreateDepositTx: %v", err)
			}
			fmt.Println("-- Transaction signed")

			err = utils.SendTransaction(signedTx, tc.rpcEndpoint)
			if err != nil {
				log.Fatalf("TestSdkCreateDepositTx: %v", err)
			}
			fmt.Println("---- Transaction sent")
			fmt.Println("********************************** Deposit success **********************************")
		})
	}
}

// func TestSdkCreateWithdrawTx(t *testing.T) {
// 	for _, tc := range testCases {
// 		t.Run("ChainID: "+strconv.FormatInt(tc.chainId, 10), func(t *testing.T) {
// 			novaSdk, err := NewNovaSDK(tc.rpcEndpoint, tc.chainId)
// 			assert.NoError(t, err)

// 			mockAmount := big.NewInt(1811735460794544158)
// 			mockReferral := big.NewInt(123)
// 			txJSON, err := novaSdk.CreateWithdrawTransaction(tc.stable, tc.mockAddress, mockAmount, mockReferral)
// 			expectedErrorMessage := "Allowance is too low. First call approve function on sDai contract."
// 			assert.Equal(t, expectedErrorMessage, err.Error())

// 			var expectedTx map[string]interface{}
// 			err = json.Unmarshal([]byte(expectedTxJSON), &expectedTx)
// 			assert.NoError(t, err)

// 			var tx map[string]interface{}
// 			err = json.Unmarshal([]byte(txJSON), &tx)
// 			assert.NoError(t, err)

// 			// Remove keys that should not be compared
// 			delete(tx, "gasPrice")
// 			delete(tx, "hash")

// 			assert.Equal(t, tx, expectedTx)
// 		})
// 	}
// }

// func TestSdkGetPrice(t *testing.T) {
// 	for _, tc := range testCases {
// 		t.Run("ChainID: "+strconv.FormatInt(tc.chainId, 10), func(t *testing.T) {
// 			novaSdk, err := NewNovaSDK(tc.rpcEndpoint, tc.chainId)
// 			assert.NoError(t, err)

// 			price, err := novaSdk.GetPrice(tc.stable)
// 			assert.NoError(t, err)

// 			stableDecimals := constants.StablecoinDetails[tc.chainId][tc.stable].Decimals
// 			base := big.NewInt(10)
// 			one := new(big.Int).Exp(base, big.NewInt(int64(stableDecimals)), nil)

// 			lowerBound := new(big.Int).Mul(big.NewInt(1), one)
// 			upperBound := new(big.Int).Mul(big.NewInt(2), one)

// 			assert.True(t, price.Cmp(lowerBound) >= 0, "Price should be at least 1e18")
// 			assert.True(t, price.Cmp(upperBound) <= 0, "Price should be at most 2e18")
// 		})
// 	}
// }

// func TestSdkGetSDaiPrice(t *testing.T) {
// 	for _, tc := range testCases {
// 		t.Run("ChainID: "+strconv.FormatInt(tc.chainId, 10), func(t *testing.T) {
// 			novaSdk, err := NewNovaSDK(tc.rpcEndpoint, tc.chainId)
// 			assert.NoError(t, err)

// 			price, err := novaSdk.GetSDaiPrice()
// 			assert.NoError(t, err)

// 			lowerBound := big.NewInt(1e18)
// 			upperBound := big.NewInt(2e18)

// 			assert.True(t, price.Cmp(lowerBound) >= 0, "Price should be at least 1e18")
// 			assert.True(t, price.Cmp(upperBound) <= 0, "Price should be at most 2e18")
// 		})
// 	}
// }

// func TestSdkGetSlippage(t *testing.T) {
// 	for _, tc := range testCases {
// 		t.Run("ChainID: "+strconv.FormatInt(tc.chainId, 10), func(t *testing.T) {
// 			novaSdk, err := NewNovaSDK(tc.rpcEndpoint, tc.chainId)
// 			assert.NoError(t, err)

// 			stableDecimals := constants.StablecoinDetails[tc.chainId][tc.stable].Decimals
// 			base := big.NewInt(10)
// 			one := new(big.Int).Exp(base, big.NewInt(int64(stableDecimals)), nil)

// 			slippageAmount := new(big.Int).Mul(big.NewInt(1000), one)
// 			slippage, expectedPrice, executedPrice, err := novaSdk.GetSlippage(tc.stable, slippageAmount)
// 			assert.NoError(t, err)

// 			slippageLowerBound := 0.0
// 			slippageUpperBound := 1.0

// 			assert.True(t, slippage >= slippageLowerBound, "slippage should be at least 0")
// 			assert.True(t, slippage <= slippageUpperBound, "slippage should be at most 1")

// 			priceLowerBound := 1.0
// 			priceUpperBound := 2.0

// 			assert.True(t, expectedPrice >= priceLowerBound, "expectedPrice should be at least 1e18")
// 			assert.True(t, expectedPrice <= priceUpperBound, "expectedPrice should be at most 2e18")

// 			assert.True(t, executedPrice >= priceLowerBound, "executedPrice should be at least 1e18")
// 			assert.True(t, executedPrice <= priceUpperBound, "executedPrice should be at most 2e18")
// 		})
// 	}
// }

// func TestSdkGetPosition(t *testing.T) {
// 	for _, tc := range testCases {
// 		t.Run("ChainID: "+strconv.FormatInt(tc.chainId, 10), func(t *testing.T) {
// 			novaSdk, err := NewNovaSDK(tc.rpcEndpoint, tc.chainId)
// 			assert.NoError(t, err)

// 			position, err := novaSdk.GetPosition(tc.stable, tc.mockAddress)
// 			assert.NoError(t, err)

// 			lowerBound := big.NewInt(0)

// 			assert.True(t, position.Cmp(lowerBound) >= 0, "Position should be at least 0")
// 		})
// 	}
// }

// func TestSdkGetSDaiTotalValue(t *testing.T) {
// 	for _, tc := range testCases {
// 		t.Run("ChainID: "+strconv.FormatInt(tc.chainId, 10), func(t *testing.T) {
// 			novaSdk, err := NewNovaSDK(tc.rpcEndpoint, tc.chainId)
// 			assert.NoError(t, err)

// 			position, err := novaSdk.GetSDaiTotalValue()
// 			assert.NoError(t, err)

// 			lowerBound := big.NewInt(0)

// 			assert.True(t, position.Cmp(lowerBound) >= 0, "Position should be at least 0")
// 		})
// 	}
// }

// func TestSdkGetSupportedStablecoins(t *testing.T) {
// 	for _, tc := range testCases {
// 		t.Run("ChainID: "+strconv.FormatInt(tc.chainId, 10), func(t *testing.T) {
// 			novaSdk, err := NewNovaSDK(tc.rpcEndpoint, tc.chainId)
// 			assert.NoError(t, err)

// 			stables, err := novaSdk.GetSupportedStablecoins()
// 			assert.NoError(t, err)

// 			assert.Len(t, stables, 1, "The length of the stablecoins array should be 1")
// 			assert.Equal(t, tc.stable, stables[0], "The first stablecoin is not correct")
// 		})
// 	}
// }
