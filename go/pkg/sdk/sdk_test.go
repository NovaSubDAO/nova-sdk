package sdk

import (
	"encoding/json"
	"math/big"
	"os"
	"strconv"
	"testing"

	"github.com/NovaSubDAO/nova-sdk/go/pkg/constants"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestSdkConfig(t *testing.T) {
	rpcEndpoint := os.Getenv("RPC_ENDPOINT")
	chainIdStr := os.Getenv("CHAIN_ID")
	chainId, err := strconv.ParseInt(chainIdStr, 10, 64)
	assert.NoError(t, err)

	novaSdk, err := NewNovaSDK(rpcEndpoint, chainId)
	assert.NoError(t, err)

	assert.Equal(t, chainId, novaSdk.Config.ChainId)
}

func TestSdkConfigDetails(t *testing.T) {
	rpcEndpoint := os.Getenv("RPC_ENDPOINT")
	chainIdStr := os.Getenv("CHAIN_ID")
	chainId, err := strconv.ParseInt(chainIdStr, 10, 64)
	assert.NoError(t, err)

	novaSdk, err := NewNovaSDK(rpcEndpoint, chainId)
	assert.NoError(t, err)

	vaultAddress := constants.ConfigDetails[chainId].VaultAddress
	assert.Equal(t, vaultAddress, novaSdk.Config.VaultAddress)

	sDai := constants.ConfigDetails[chainId].SDai
	assert.Equal(t, sDai, novaSdk.Config.SDai)

	vaultDecimals := constants.ConfigDetails[chainId].VaultDecimals
	assert.Equal(t, vaultDecimals, novaSdk.Config.VaultDecimals)
}

// func TestSdkCreateDepositTx(t *testing.T) {
// 	rpcEndpoint := os.Getenv("RPC_ENDPOINT")
// 	chainIdStr := os.Getenv("CHAIN_ID")
// 	chainId, err := strconv.ParseInt(chainIdStr, 10, 64)
// 	assert.NoError(t, err)

// 	novaSdk, err := NewNovaSDK(rpcEndpoint, chainId)
// 	assert.NoError(t, err)

// 	mockAmount := big.NewInt(1e3)
// 	mockReferral := big.NewInt(123)
// 	mockAddress := common.HexToAddress("0x83F20F44975D03b1b09e64809B757c47f942BEeA")
// 	txJSON, err := novaSdk.SdkDomain.CreateDepositTransaction(constants.USDC, mockAddress, mockAmount, mockReferral)
// 	assert.NoError(t, err)

// 	assert.Equal(t, txJSON, "")

// 	expectedTxJSON := `{"type":"0x0","nonce":"0x0","to":"0x36a2f7fb07c102415afe2461a9a43377970e081c","gas":"0x1e8480","maxPriorityFeePerGas":null,"maxFeePerGas":null,"value":"0x0","input":"0xd2d0e0660000000000000000000000000b2c639c533813f4aa9d7837caf62653d097ff8500000000000000000000000000000000000000000000000000000000000003e8000000000000000000000000000000000000000000000000000000000000007b","v":"0x0","r":"0x0","s":"0x0"}`

// 	var expectedTx map[string]interface{}
// 	err = json.Unmarshal([]byte(expectedTxJSON), &expectedTx)
// 	assert.NoError(t, err)

// 	var tx map[string]interface{}
// 	err = json.Unmarshal([]byte(txJSON), &tx)
// 	assert.NoError(t, err)

// 	// Remove keys that should not be compared
// 	delete(tx, "gasPrice")
// 	delete(tx, "hash")

// 	assert.Equal(t, tx, expectedTx)
// }

func TestSdkCreateWithdrawTx(t *testing.T) {
	rpcEndpoint := os.Getenv("RPC_ENDPOINT")
	chainIdStr := os.Getenv("CHAIN_ID")
	chainId, err := strconv.ParseInt(chainIdStr, 10, 64)
	assert.NoError(t, err)

	novaSdk, err := NewNovaSDK(rpcEndpoint, chainId)
	assert.NoError(t, err)

	mockAmount := big.NewInt(1e3)
	mockReferral := big.NewInt(123)
	mockAddress := common.HexToAddress("0x83F20F44975D03b1b09e64809B757c47f942BEeA")
	txJSON, err := novaSdk.SdkDomain.CreateWithdrawTransaction(constants.USDC, mockAddress, mockAmount, mockReferral)
	assert.NoError(t, err)

	expectedTxJSON := `{"type":"0x0","nonce":"0x0","to":"0x36a2f7fb07c102415afe2461a9a43377970e081c","gas":"0x1e8480","maxPriorityFeePerGas":null,"maxFeePerGas":null,"value":"0x0","input":"0xf3fef3a30000000000000000000000000b2c639c533813f4aa9d7837caf62653d097ff8500000000000000000000000000000000000000000000000000000000000003e8","v":"0x0","r":"0x0","s":"0x0"}`

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
}

func TestSdkGetPrice(t *testing.T) {
	rpcEndpoint := os.Getenv("RPC_ENDPOINT")
	chainIdStr := os.Getenv("CHAIN_ID")
	chainId, err := strconv.ParseInt(chainIdStr, 10, 64)
	assert.NoError(t, err)

	novaSdk, err := NewNovaSDK(rpcEndpoint, chainId)
	assert.NoError(t, err)

	price, err := novaSdk.SdkDomain.GetPrice(constants.USDC)
	assert.NoError(t, err)

	// TODO: fix GetPrice function and return price with stablecoin decimals
	lowerBound := big.NewInt(1e12)
	upperBound := big.NewInt(2e12)

	assert.True(t, price.Cmp(lowerBound) >= 0, "Price should be at least 1e18")
	assert.True(t, price.Cmp(upperBound) <= 0, "Price should be at most 2e18")
}

func TestSdkGetSDaiPrice(t *testing.T) {
	rpcEndpoint := os.Getenv("RPC_ENDPOINT")
	chainIdStr := os.Getenv("CHAIN_ID")
	chainId, err := strconv.ParseInt(chainIdStr, 10, 64)
	assert.NoError(t, err)

	novaSdk, err := NewNovaSDK(rpcEndpoint, chainId)
	assert.NoError(t, err)

	price, err := novaSdk.SdkDomain.GetSDaiPrice()
	assert.NoError(t, err)

	lowerBound := big.NewInt(1e18)
	upperBound := big.NewInt(2e18)

	assert.True(t, price.Cmp(lowerBound) >= 0, "Price should be at least 1e18")
	assert.True(t, price.Cmp(upperBound) <= 0, "Price should be at most 2e18")
}

func TestSdkGetSlippage(t *testing.T) {
	rpcEndpoint := os.Getenv("RPC_ENDPOINT")
	chainIdStr := os.Getenv("CHAIN_ID")
	chainId, err := strconv.ParseInt(chainIdStr, 10, 64)
	assert.NoError(t, err)

	novaSdk, err := NewNovaSDK(rpcEndpoint, chainId)
	assert.NoError(t, err)

	slippageAmount := big.NewInt(1e10)
	slippage, expectedPrice, executedPrice, err := novaSdk.SdkDomain.GetSlippage(constants.USDC, slippageAmount)
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
}

func TestSdkGetPosition(t *testing.T) {
	rpcEndpoint := os.Getenv("RPC_ENDPOINT")
	chainIdStr := os.Getenv("CHAIN_ID")
	chainId, err := strconv.ParseInt(chainIdStr, 10, 64)
	assert.NoError(t, err)

	novaSdk, err := NewNovaSDK(rpcEndpoint, chainId)
	assert.NoError(t, err)

	mockAddress := common.HexToAddress("0x83F20F44975D03b1b09e64809B757c47f942BEeA")
	position, err := novaSdk.SdkDomain.GetPosition(constants.USDC, mockAddress)
	assert.NoError(t, err)

	lowerBound := big.NewInt(0)

	assert.True(t, position.Cmp(lowerBound) >= 0, "Position should be at least 0")
}

func TestSdkGetSDaiTotalValue(t *testing.T) {
	rpcEndpoint := os.Getenv("RPC_ENDPOINT")
	chainIdStr := os.Getenv("CHAIN_ID")
	chainId, err := strconv.ParseInt(chainIdStr, 10, 64)
	assert.NoError(t, err)

	novaSdk, err := NewNovaSDK(rpcEndpoint, chainId)
	assert.NoError(t, err)

	position, err := novaSdk.SdkDomain.GetSDaiTotalValue()
	assert.NoError(t, err)

	lowerBound := big.NewInt(0)

	assert.True(t, position.Cmp(lowerBound) >= 0, "Position should be at least 0")
}

func TestSdkGetSupportedStablecoins(t *testing.T) {
	rpcEndpoint := os.Getenv("RPC_ENDPOINT")
	chainIdStr := os.Getenv("CHAIN_ID")
	chainId, err := strconv.ParseInt(chainIdStr, 10, 64)
	assert.NoError(t, err)

	novaSdk, err := NewNovaSDK(rpcEndpoint, chainId)
	assert.NoError(t, err)

	stables, err := novaSdk.SdkDomain.GetSupportedStablecoins()
	assert.NoError(t, err)

	assert.Len(t, stables, 1, "The length of the stablecoins array should be 1")
	assert.Equal(t, constants.USDC, stables[0], "The first stablecoin should be USDC")
}
