package sdk

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"testing"

	"github.com/NovaSubDAO/nova-sdk/go/pkg/constants"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	rpcEndpoint   string
	chainId       int64
	stable        constants.Stablecoin
	mockAddress   common.Address
	privateKeyHex string
}

var testCases = []testCase{
	{
		rpcEndpoint:   os.Getenv("ETH_RPC_ENDPOINT"),
		chainId:       1,
		stable:        constants.DAI,
		mockAddress:   common.HexToAddress(os.Getenv("CLIENT_ADDRESS")),
		privateKeyHex: os.Getenv("PRIVATE_KEY"),
	},
	{
		rpcEndpoint:   os.Getenv("OPT_RPC_ENDPOINT_ANVIL"),
		chainId:       10,
		stable:        constants.USDC,
		mockAddress:   common.HexToAddress(os.Getenv("CLIENT_ADDRESS")),
		privateKeyHex: os.Getenv("PRIVATE_KEY"),
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
			novaSdk, err := NewNovaSDK(tc.rpcEndpoint, tc.chainId)
			assert.NoError(t, err)

			mockAmount := big.NewInt(20000)
			mockReferral := big.NewInt(123)

			txJSON, err := novaSdk.CreateDepositTransaction(tc.stable, tc.mockAddress, mockAmount, mockReferral)
			expectedErrorMessage := fmt.Sprintf("Allowance is too low. First call approve function on %s contract.", tc.stable)
			assert.Equal(t, expectedErrorMessage, err.Error())

			var expectedTxJSON string
			if tc.chainId == 1 {
				expectedTxJSON = `{"type":"0x0","nonce":"0x2d","to":"0x83f20f44975d03b1b09e64809b757c47f942beea","gas":"0x1e8480","maxPriorityFeePerGas":null,"maxFeePerGas":null,"value":"0x0","input":"0x6e553f6500000000000000000000000000000000000000000000000000000000000003e800000000000000000000000047ac0fb4f2d84898e4d9e7b4dab3c24507a6d503","v":"0x0","r":"0x0","s":"0x0"}`
			}
			if tc.chainId == 10 {
				expectedTxJSON = `{"type":"0x0","nonce":"0x2e854c","to":"0x36a2f7fb07c102415afe2461a9a43377970e081c","gas":"0x1e8480","maxPriorityFeePerGas":null,"maxFeePerGas":null,"value":"0x0","input":"0xd2d0e0660000000000000000000000000b2c639c533813f4aa9d7837caf62653d097ff8500000000000000000000000000000000000000000000000000000000000003e8000000000000000000000000000000000000000000000000000000000000007b","v":"0x0","r":"0x0","s":"0x0"}`
			}

			var expectedTx map[string]interface{}
			err = json.Unmarshal([]byte(expectedTxJSON), &expectedTx)
			assert.NoError(t, err)

			var tx map[string]interface{}
			err = json.Unmarshal([]byte(txJSON), &tx)
			assert.NoError(t, err)

			// Remove keys that should not be compared
			delete(tx, "gasPrice")
			delete(tx, "hash")

			assert.Equal(t, tx, expectedTx)
		})
	}
}

func TestSdkCreateWithdrawTx(t *testing.T) {
	for _, tc := range testCases {
		t.Run("ChainID: "+strconv.FormatInt(tc.chainId, 10), func(t *testing.T) {
			novaSdk, err := NewNovaSDK(tc.rpcEndpoint, tc.chainId)
			assert.NoError(t, err)

			mockAmount := big.NewInt(1811735460794544158)
			mockReferral := big.NewInt(123)
			txJSON, err := novaSdk.CreateWithdrawTransaction(tc.stable, tc.mockAddress, mockAmount, mockReferral)
			expectedErrorMessage := "Allowance is too low. First call approve function on sDai contract."
			assert.Equal(t, expectedErrorMessage, err.Error())

			var expectedTxJSON string
			if tc.chainId == 1 {
				expectedTxJSON = `{"type":"0x0","nonce":"0x1","to":"0x83f20f44975d03b1b09e64809b757c47f942beea","gas":"0x1e8480","maxPriorityFeePerGas":null,"maxFeePerGas":null,"value":"0x0","input":"0xb460af9400000000000000000000000000000000000000000000000000000000000003e800000000000000000000000083f20f44975d03b1b09e64809b757c47f942beea00000000000000000000000083f20f44975d03b1b09e64809b757c47f942beea","v":"0x0","r":"0x0","s":"0x0"}`
			}
			if tc.chainId == 10 {
				expectedTxJSON = `{"type":"0x0","nonce":"0x0","to":"0x36a2f7fb07c102415afe2461a9a43377970e081c","gas":"0x1e8480","maxPriorityFeePerGas":null,"maxFeePerGas":null,"value":"0x0","input":"0xf3fef3a30000000000000000000000000b2c639c533813f4aa9d7837caf62653d097ff8500000000000000000000000000000000000000000000000000000000000003e8","v":"0x0","r":"0x0","s":"0x0"}`
			}

			var expectedTx map[string]interface{}
			err = json.Unmarshal([]byte(expectedTxJSON), &expectedTx)
			assert.NoError(t, err)

			var tx map[string]interface{}
			err = json.Unmarshal([]byte(txJSON), &tx)
			assert.NoError(t, err)

			// Remove keys that should not be compared
			delete(tx, "gasPrice")
			delete(tx, "hash")

			assert.Equal(t, tx, expectedTx)
		})
	}
}

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
