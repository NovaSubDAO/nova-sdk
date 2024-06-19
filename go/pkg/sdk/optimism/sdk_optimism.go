package optimism

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/NovaSubDAO/nova-sdk/go/pkg/config"
	"github.com/NovaSubDAO/nova-sdk/go/pkg/constants"
	optimismContracts "github.com/NovaSubDAO/nova-sdk/go/pkg/sdk/optimism/abis"
	"github.com/NovaSubDAO/nova-sdk/go/pkg/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type SdkOptimism struct {
	Config   *config.Config
	Contract *optimismContracts.SavingsDaiCaller
}

func NewSdkOptimism(cfg *config.Config) (*SdkOptimism, error) {
	client, err := ethclient.Dial(cfg.RpcEndpoint)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to Optimism client: %w", err)
	}

	contract, err := optimismContracts.NewSavingsDaiCaller(common.HexToAddress(cfg.SDai), client)
	if err != nil {
		return nil, fmt.Errorf("Failed to instantiate contract caller: %w", err)
	}

	return &SdkOptimism{Config: cfg, Contract: contract}, nil
}

func (sdk *SdkOptimism) isStablecoinSupported(stable constants.Stablecoin) bool {
	if stablecoins, ok := constants.StablecoinDetails[sdk.Config.ChainId]; ok {
		_, exists := stablecoins[stable]
		return exists
	}
	return false
}

func (sdk *SdkOptimism) getPriceFromInput(input optimismContracts.IMixedRouteQuoterV1QuoteExactInputSingleV3Params) (*big.Int, error) {
	client, err := ethclient.Dial(sdk.Config.RpcEndpoint)
	if err != nil {
		return nil, fmt.Errorf("Error loading client: %w", err)
	}

	quoterAddress := common.HexToAddress("0xa4ac92a0F54f1a447c55a4082c90742F5E76Df62")
	quoter, err := optimismContracts.NewMixedRouteQuoterV1(quoterAddress, client)
	if err != nil {
		return nil, fmt.Errorf("Failed to load MixedRouteQuoterV1 contract: %w", err)
	}

	var output []interface{}
	rawCaller := &optimismContracts.MixedRouteQuoterV1Raw{Contract: quoter}
	err = rawCaller.Call(nil, &output, "quoteExactInputSingleV3", input)
	if err != nil {
		return nil, fmt.Errorf("Failed to call quoteExactInputSingleV3 function: %w", err)
	}
	amountOut := output[0].(*big.Int)

	return amountOut, nil
}

func (sdk *SdkOptimism) GetPrice(stable constants.Stablecoin) (*big.Int, error) {
	ok := sdk.isStablecoinSupported(stable)
	if !ok {
		return big.NewInt(0), fmt.Errorf("stablecoin %s is not supported", stable)
	}

	// Packing the input arguments
	stableAddress := constants.StablecoinDetails[sdk.Config.ChainId][stable].Address
	stableDecimals := constants.StablecoinDetails[sdk.Config.ChainId][stable].Decimals
	tokenIn := common.HexToAddress(stableAddress)
	tokenOut := common.HexToAddress(sdk.Config.SDai)

	base := big.NewInt(10)
	one := new(big.Int).Exp(base, big.NewInt(int64(stableDecimals)), nil)

	input := optimismContracts.IMixedRouteQuoterV1QuoteExactInputSingleV3Params{
		TokenIn:           tokenIn,
		TokenOut:          tokenOut,
		AmountIn:          one,
		TickSpacing:       big.NewInt(1),
		SqrtPriceLimitX96: big.NewInt(0),
	}

	result, err := sdk.getPriceFromInput(input)
	if err != nil {
		return nil, fmt.Errorf("Failed to get price from input params: %w", err)
	}

	resultFloat := new(big.Float).SetInt(result)
	factorFloat := new(big.Float).SetInt(big.NewInt(1e18))
	decimalsFactorFloat := new(big.Float).SetInt(one)
	price := new(big.Float).Quo(factorFloat, resultFloat)
	price.Mul(price, decimalsFactorFloat)

	priceInt := new(big.Int)
	price.Int(priceInt)

	return priceInt, nil
}

func (sdk *SdkOptimism) GetPosition(stable constants.Stablecoin, address common.Address) (*big.Int, error) {
	ok := sdk.isStablecoinSupported(stable)
	if !ok {
		return big.NewInt(0), fmt.Errorf("stablecoin %s is not supported", stable)
	}

	balance, err := sdk.Contract.BalanceOf(nil, address)
	if err != nil {
		return big.NewInt(0), err
	}

	price, err := sdk.GetPrice(stable)
	if err != nil {
		return big.NewInt(0), err
	}

	value := new(big.Int).Mul(balance, price)
	factor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(sdk.Config.VaultDecimals)), nil)
	valueNormalized := new(big.Int).Div(value, factor)

	return valueNormalized, nil
}

func (sdk *SdkOptimism) GetSDaiPrice() (*big.Int, error) {
	// Packing the input arguments
	client, err := ethclient.Dial(sdk.Config.RpcEndpoint)
	if err != nil {
		return nil, fmt.Errorf("Error loading client: %w", err)
	}

	oracleAddress := common.HexToAddress("0x33a3aB524A43E69f30bFd9Ae97d1Ec679FF00B64")
	oracle, err := optimismContracts.NewDSRAuthOracleCaller(oracleAddress, client)
	if err != nil {
		return nil, fmt.Errorf("Failed to load DSRAuthOracle contract: %w", err)
	}

	result, err := oracle.GetPotData(nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to call GetPotData function: %w", err)
	}

	blockTimestamp := big.NewInt(time.Now().Unix())
	chi := new(big.Int).Div(new(big.Int).Mul(utils.Rpow(result.Dsr, new(big.Int).Sub(blockTimestamp, result.Rho)), result.Chi), utils.RAY)

	factor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(sdk.Config.VaultDecimals)), nil)
	price := new(big.Int).Div(new(big.Int).Mul(factor, chi), utils.RAY)

	return price, nil
}

func (sdk *SdkOptimism) GetSDaiTotalValue() (*big.Int, error) {
	totalSupply, err := sdk.Contract.TotalSupply(nil)
	if err != nil {
		return big.NewInt(0), err
	}

	price, err := sdk.GetSDaiPrice()
	if err != nil {
		return big.NewInt(0), err
	}

	totalValue := new(big.Int).Mul(totalSupply, price)
	factor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(sdk.Config.VaultDecimals)), nil)
	totalValueNormalized := new(big.Int).Div(totalValue, factor)

	return totalValueNormalized, nil
}

func (sdk *SdkOptimism) GetSupportedStablecoins() ([]constants.Stablecoin, error) {
	if stablecoins, ok := constants.StablecoinDetails[sdk.Config.ChainId]; ok {
		var supportedStablecoins []constants.Stablecoin
		for sc := range stablecoins {
			supportedStablecoins = append(supportedStablecoins, sc)
		}
		return supportedStablecoins, nil
	}
	return nil, fmt.Errorf("no stablecoins found for chain ID %d", sdk.Config.ChainId)
}

func (sdk *SdkOptimism) GetSlippage(stable constants.Stablecoin, amount *big.Int) (float64, float64, float64, error) {
	ok := sdk.isStablecoinSupported(stable)
	if !ok {
		return 0, 0, 0, fmt.Errorf("stablecoin %s is not supported", stable)
	}

	// Packing the input arguments
	stableAddress := constants.StablecoinDetails[sdk.Config.ChainId][stable].Address
	stableDecimals := constants.StablecoinDetails[sdk.Config.ChainId][stable].Decimals
	sDaiDecimals := sdk.Config.VaultDecimals
	tokenIn := common.HexToAddress(stableAddress)
	tokenOut := common.HexToAddress(sdk.Config.SDai)

	base := big.NewInt(10)
	one := new(big.Int).Exp(base, big.NewInt(int64(stableDecimals)), nil)

	input := optimismContracts.IMixedRouteQuoterV1QuoteExactInputSingleV3Params{
		TokenIn:           tokenIn,
		TokenOut:          tokenOut,
		AmountIn:          one,
		TickSpacing:       big.NewInt(1),
		SqrtPriceLimitX96: big.NewInt(0),
	}
	resultOne, err := sdk.getPriceFromInput(input)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("Failed to get price from input params: %w", err)
	}

	if resultOne.Cmp(big.NewInt(0)) == 0 {
		return 0, 0, 0, fmt.Errorf("result of one unit of token is zero")
	}

	stableTenPow := new(big.Int).Exp(base, big.NewInt(int64(stableDecimals)), nil)
	sDaiTenPow := new(big.Int).Exp(base, big.NewInt(int64(sDaiDecimals)), nil)
	stableTenPowFloat := new(big.Float).SetInt(stableTenPow)
	sDaiTenPowFloat := new(big.Float).SetInt(sDaiTenPow)

	oneFloat := new(big.Float).SetInt(one)
	oneFloat.Quo(oneFloat, stableTenPowFloat)

	resultOneFloat := new(big.Float).SetInt(resultOne)
	resultOneFloat.Quo(resultOneFloat, sDaiTenPowFloat)

	expectedPrice := new(big.Float).Quo(oneFloat, resultOneFloat)

	input.AmountIn = amount
	resultAmount, err := sdk.getPriceFromInput(input)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("Failed to get price from input params: %w", err)
	}

	if amount.Cmp(big.NewInt(0)) == 0 || resultAmount.Cmp(big.NewInt(0)) == 0 {
		return 0, 0, 0, fmt.Errorf("amount or result for the specified amount is zero")
	}

	amountFloat := new(big.Float).SetInt(amount)
	amountFloat.Quo(amountFloat, stableTenPowFloat)

	resultAmountFloat := new(big.Float).SetInt(resultAmount)
	resultAmountFloat.Quo(resultAmountFloat, sDaiTenPowFloat)

	executedPrice := new(big.Float).Quo(amountFloat, resultAmountFloat)

	num := new(big.Float).Sub(executedPrice, expectedPrice)
	den := expectedPrice

	percentageChange, _ := new(big.Float).Quo(num, den).Float64()
	percentageChange *= 100

	expectedPriceFloat, _ := expectedPrice.Float64()
	executedPriceFloat, _ := executedPrice.Float64()

	return percentageChange, expectedPriceFloat, executedPriceFloat, nil
}

func (sdk *SdkOptimism) CreateDepositTransaction(stable constants.Stablecoin, fromAddress common.Address, amount *big.Int, referral *big.Int) (string, error) {
	ok := sdk.isStablecoinSupported(stable)
	if !ok {
		return "", fmt.Errorf("stablecoin %s is not supported", stable)
	}

	stableAddress := common.HexToAddress(constants.StablecoinDetails[sdk.Config.ChainId][stable].Address)
	vaultAddress := common.HexToAddress(sdk.Config.VaultAddress)

	client, err := ethclient.Dial(sdk.Config.RpcEndpoint)
	if err != nil {
		return "", fmt.Errorf("Failed to connect to the Optimism client: %v", err)
	}

	stableContract, err := optimismContracts.NewFiatTokenV22Caller(stableAddress, client)
	if err != nil {
		return "", fmt.Errorf("Failed to load FiatTokenV22 contract: %w", err)
	}

	result, err := stableContract.Allowance(nil, fromAddress, vaultAddress)
	if err != nil {
		return "", fmt.Errorf("Failed to call Allowance function: %w", err)
	}

	if amount.Cmp(result) > 0 {
		return "", fmt.Errorf("Allowance is too low. First call approve function on USDC contract.")
	}

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", fmt.Errorf("Failed to get nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("Failed to suggest gas price: %v", err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(optimismContracts.NovaVaultMetaData.ABI))
	if err != nil {
		return "", fmt.Errorf("Failed to parse contract ABI: %w", err)
	}

	referralUint16 := uint16(referral.Uint64())
	data, err := contractAbi.Pack("deposit", stableAddress, amount, referralUint16)
	if err != nil {
		return "", fmt.Errorf("ABI pack failed: %v", err)
	}

	// Estimating the gas needed for the transaction
	msg := ethereum.CallMsg{From: fromAddress, To: &vaultAddress, GasPrice: gasPrice, Value: big.NewInt(0), Data: data}
	gasLimit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		log.Printf("Gas estimation failed, using fallback gas limit: %v", err)
		gasLimit = 2000000
	}

	tx := types.NewTransaction(nonce, vaultAddress, big.NewInt(0), gasLimit, gasPrice, data)

	txJSON, err := json.Marshal(tx)
	if err != nil {
		return "", fmt.Errorf("Failed to marshal transaction: %w", err)
	}

	return string(txJSON), nil
}

func (sdk *SdkOptimism) CreateWithdrawTransaction(stable constants.Stablecoin, fromAddress common.Address, amount *big.Int, referral *big.Int) (string, error) {
	ok := sdk.isStablecoinSupported(stable)
	if !ok {
		return "", fmt.Errorf("stablecoin %s is not supported", stable)
	}

	stableAddress := common.HexToAddress(constants.StablecoinDetails[sdk.Config.ChainId][stable].Address)
	vaultAddress := common.HexToAddress(sdk.Config.VaultAddress)
	sDaiAddress := common.HexToAddress(sdk.Config.SDai)

	client, err := ethclient.Dial(sdk.Config.RpcEndpoint)
	if err != nil {
		return "", fmt.Errorf("Failed to connect to the Ethereum client: %v", err)
	}

	sDaiContract, err := optimismContracts.NewSavingsDaiCaller(sDaiAddress, client)
	if err != nil {
		return "", fmt.Errorf("Failed to load SavingsDai contract: %w", err)
	}

	result, err := sDaiContract.Allowance(nil, fromAddress, vaultAddress)
	if err != nil {
		return "", fmt.Errorf("Failed to call Allowance function: %w", err)
	}

	if amount.Cmp(result) > 0 {
		return "", fmt.Errorf("Allowance is too low. First call approve function on sDai contract.")
	}

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", fmt.Errorf("Failed to get nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("Failed to suggest gas price: %v", err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(optimismContracts.NovaVaultMetaData.ABI))
	if err != nil {
		return "", fmt.Errorf("Failed to parse contract ABI: %w", err)
	}

	data, err := contractAbi.Pack("withdraw", stableAddress, amount)
	if err != nil {
		return "", fmt.Errorf("ABI pack failed: %v", err)
	}

	// Estimating the gas needed for the transaction
	msg := ethereum.CallMsg{From: fromAddress, To: &vaultAddress, GasPrice: gasPrice, Value: big.NewInt(0), Data: data}
	gasLimit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		log.Printf("Gas estimation failed, using fallback gas limit: %v", err)
		gasLimit = 2000000
	}

	tx := types.NewTransaction(nonce, vaultAddress, big.NewInt(0), gasLimit, gasPrice, data)

	txJSON, err := json.Marshal(tx)
	if err != nil {
		return "", fmt.Errorf("Failed to marshal transaction: %w", err)
	}

	return string(txJSON), nil
}
