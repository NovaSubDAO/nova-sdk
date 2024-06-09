package optimism

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/NovaSubDAO/nova-sdk/go/pkg/config"
	"github.com/NovaSubDAO/nova-sdk/go/pkg/contracts"
	optimismContracts "github.com/NovaSubDAO/nova-sdk/go/pkg/sdk/optimism/abis"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type SdkOptimism struct {
	Config   *config.Config
	Contract *contracts.ContractsCaller
}

func NewSdkOptimism(cfg *config.Config) (*SdkOptimism, error) {
	client, err := ethclient.Dial(cfg.RpcEndpoint)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to Optimism client: %w", err)
	}

	contract, err := contracts.NewContractsCaller(common.HexToAddress(cfg.VaultAddress), client)
	if err != nil {
		return nil, fmt.Errorf("Failed to instantiate contract caller: %w", err)
	}

	return &SdkOptimism{Config: cfg, Contract: contract}, nil
}

func (sdk *SdkOptimism) getPriceFromInput(input optimismContracts.IMixedRouteQuoterV1QuoteExactInputSingleV2Params) (*big.Int, error) {
	client, err := ethclient.Dial(sdk.Config.RpcEndpoint)
	if err != nil {
		return nil, fmt.Errorf("Error loading client: %w", err)
	}

	contractAddress := common.HexToAddress("0xa4ac92a0F54f1a447c55a4082c90742F5E76Df62")

	contractAbi, err := abi.JSON(strings.NewReader(optimismContracts.MixedRouteQuoterV1ABI))
	if err != nil {
		return nil, fmt.Errorf("Failed to parse contract ABI: %w", err)
	}

	data, err := contractAbi.Pack("quoteExactInputSingleV2", input)
	if err != nil {
		return nil, fmt.Errorf("Failed to pack arguments: %w", err)
	}

	// Setting up the call message
	msg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}
	ctx := context.Background()

	// Making the call
	output, err := client.CallContract(ctx, msg, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to call contract: %w", err)
	}

	// Process the output data (this depends on what the function returns)
	results, err := contractAbi.Unpack("quoteExactInputSingleV2", output)
	if err != nil {
		return nil, fmt.Errorf("Failed to unpack function output: %w", err)
	}

	// Type assert the result as *big.Int
	result, ok := results[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("result type assertion to *big.Int failed")
	}

	return result, nil
}

func (sdk *SdkOptimism) GetPrice() (*big.Int, error) {
	// Define the parameters for the function call
	tokenIn := common.HexToAddress("0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85")
	tokenOut := common.HexToAddress("0x2218a117083f5B482B0bB821d27056Ba9c04b1D3")

	// Packing the input arguments
	input := optimismContracts.IMixedRouteQuoterV1QuoteExactInputSingleV2Params{
		TokenIn:  tokenIn,
		TokenOut: tokenOut,
		Stable:   false,
		AmountIn: big.NewInt(1),
	}

	result, err := sdk.getPriceFromInput(input)
	if err != nil {
		return nil, fmt.Errorf("Failed to get price from input params: %w", err)
	}

	return result, nil
}

func (sdk *SdkOptimism) GetPosition(address common.Address) (*big.Int, error) {
	return nil, fmt.Errorf("Not yet implemented")
}

func (sdk *SdkOptimism) GetTotalValue() (*big.Int, error) {
	return nil, fmt.Errorf("Not yet implemented")
}

func (sdk *SdkOptimism) GetSlippage(stable common.Address, amount *big.Int) (float64, error) {
	// Define the parameters for the function call
	tokenIn := common.HexToAddress("0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85")
	tokenOut := common.HexToAddress("0x2218a117083f5B482B0bB821d27056Ba9c04b1D3")

	// Packing the input arguments
	input := optimismContracts.IMixedRouteQuoterV1QuoteExactInputSingleV2Params{
		TokenIn:  tokenIn,
		TokenOut: tokenOut,
		Stable:   false,
		AmountIn: big.NewInt(1),
	}
	resultOne, err := sdk.getPriceFromInput(input)
	if err != nil {
		return 0, fmt.Errorf("Failed to get price from input params: %w", err)
	}

	input = optimismContracts.IMixedRouteQuoterV1QuoteExactInputSingleV2Params{
		TokenIn:  tokenIn,
		TokenOut: tokenOut,
		Stable:   false,
		AmountIn: amount,
	}
	resultAmount, err := sdk.getPriceFromInput(input)
	if err != nil {
		return 0, fmt.Errorf("Failed to get price from input params: %w", err)
	}

	resultOneFloat := new(big.Float).SetInt(resultOne)
	resultAmountFloat := new(big.Float).SetInt(resultAmount)
	amountFloat := new(big.Float).SetInt(amount)
	resultAmountFloat.Quo(resultAmountFloat, amountFloat)
	diff := new(big.Float).Sub(resultAmountFloat, resultOneFloat)
	percentageChange, _ := new(big.Float).Quo(diff, resultOneFloat).Float64()

	return percentageChange, nil
}

func (sdk *SdkOptimism) CreateDepositTransaction(fromAddress common.Address, stable common.Address, amount *big.Int, referral *big.Int) (string, error) {
	client, err := ethclient.Dial(sdk.Config.RpcEndpoint)
	if err != nil {
		return "", fmt.Errorf("Failed to connect to the Optimism client: %v", err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", fmt.Errorf("Failed to get nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("Failed to suggest gas price: %v", err)
	}

	contractAddress := common.HexToAddress(sdk.Config.VaultAddress)

	contractAbi, err := abi.JSON(strings.NewReader(optimismContracts.NovaVaultABI))
	if err != nil {
		return "", fmt.Errorf("Failed to parse contract ABI: %w", err)
	}

	referralUint16 := uint16(referral.Uint64())
	data, err := contractAbi.Pack("deposit", stable, amount, referralUint16)
	if err != nil {
		return "", fmt.Errorf("ABI pack failed: %v", err)
	}

	// Estimating the gas needed for the transaction
	msg := ethereum.CallMsg{From: fromAddress, To: &contractAddress, GasPrice: gasPrice, Value: big.NewInt(0), Data: data}
	gasLimit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		log.Printf("Gas estimation failed, using fallback gas limit: %v", err)
		gasLimit = 2000000 // Fallback gas limit
	}

	tx := types.NewTransaction(nonce, contractAddress, big.NewInt(0), gasLimit, gasPrice, data)

	txJSON, err := json.Marshal(tx)
	if err != nil {
		return "", fmt.Errorf("Failed to marshal transaction: %w", err)
	}

	return string(txJSON), nil
}

func (sdk *SdkOptimism) CreateWithdrawTransaction(fromAddress common.Address, stable common.Address, amount *big.Int, referral *big.Int) (string, error) {
	client, err := ethclient.Dial(sdk.Config.RpcEndpoint)
	if err != nil {
		return "", fmt.Errorf("Failed to connect to the Ethereum client: %v", err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", fmt.Errorf("Failed to get nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("Failed to suggest gas price: %v", err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(optimismContracts.NovaVaultABI))
	if err != nil {
		return "", fmt.Errorf("Failed to parse contract ABI: %w", err)
	}

	contractAddress := common.HexToAddress(sdk.Config.VaultAddress)

	data, err := contractAbi.Pack("withdraw", stable, amount)
	if err != nil {
		return "", fmt.Errorf("ABI pack failed: %v", err)
	}

	// Estimating the gas needed for the transaction
	msg := ethereum.CallMsg{From: fromAddress, To: &contractAddress, GasPrice: gasPrice, Value: big.NewInt(0), Data: data}
	gasLimit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		log.Printf("Gas estimation failed, using fallback gas limit: %v", err)
		gasLimit = 2000000 // Fallback gas limit
	}

	tx := types.NewTransaction(nonce, contractAddress, big.NewInt(0), gasLimit, gasPrice, data)

	txJSON, err := json.Marshal(tx)
	if err != nil {
		return "", fmt.Errorf("Failed to marshal transaction: %w", err)
	}

	return string(txJSON), nil
}

func (sdk *SdkOptimism) Deposit(assets *big.Int, receiver common.Address, referral big.Int) (*types.Transaction, error) {
	return nil, fmt.Errorf("Not yet implemented")
}

func (sdk *SdkOptimism) Withdraw(assets *big.Int, receiver common.Address, referral big.Int) (*types.Transaction, error) {
	return nil, fmt.Errorf("Not yet implemented")
}
