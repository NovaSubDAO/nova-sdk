package optimism

import (
    "context"
	"encoding/json"
    "fmt"
    "log"
    "math/big"
    "os"
    "strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
    "github.com/NovaSubDAO/nova-sdk/go/pkg/config"
    "github.com/NovaSubDAO/nova-sdk/go/pkg/contracts"
)

type SdkOptimism struct {
    Config *config.Config
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

func (sdk *SdkOptimism) GetPrice() (*big.Int, error) {
	return nil, fmt.Errorf("Not yet implemented")
}

func (sdk *SdkOptimism) GetPosition(address common.Address) (*big.Int, error) {
	return nil, fmt.Errorf("Not yet implemented")
}

func (sdk *SdkOptimism) GetTotalValue() (*big.Int, error) {
	return nil, fmt.Errorf("Not yet implemented")
}

func (sdk *SdkOptimism) GetSlippage(amount *big.Int) (*big.Int, error) {
    // cfg, err := sdk.Config.LoadConfig()
    // if err != nil {
	// 	return nil, fmt.Errorf("Error loading configuration: %w", err)
    // }

    // client, err := ethclient.Dial(cfg.RpcEndpoint)
    // if err != nil {
	// 	return nil, fmt.Errorf("Error loading client: %w", err)
    // }

    // contractAddress := common.HexToAddress("0xa4ac92a0F54f1a447c55a4082c90742F5E76Df62")

    // abiPath := "abis/MixedRouteQuoterV1.json"
    // file, err := os.ReadFile(abiPath)
    // if err != nil {
	// 	return nil, fmt.Errorf("Failed to read ABI file: %w", err)
    // }

    // parsedABI, err := abi.JSON(strings.NewReader(string(file)))
    // if err != nil {
	// 	return nil, fmt.Errorf("Failed to parse ABI: %w", err)
    // }

    // // Define the parameters for the function call
    // tokenIn := common.HexToAddress("0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85")
    // tokenOut := common.HexToAddress("0x2218a117083f5B482B0bB821d27056Ba9c04b1D3")
    // amountIn := big.NewInt(100) // example value; adjust as needed
    // fee := big.NewInt(3000)     // Fee tier, adjust according to your scenario

    // // Packing the input arguments
    // data, err := parsedABI.Pack("quoteExactInputSingleV2", tokenIn, tokenOut, fee, amountIn, false)
    // if err != nil {
	// 	return nil, fmt.Errorf("Failed to pack arguments: %w", err)
    // }

    // // Setting up the call message
    // msg := ethereum.CallMsg{
    //     To:   &contractAddress,
    //     Data: data,
    // }
    // ctx := context.Background()

    // // Making the call
    // output, err := client.CallContract(ctx, msg, nil)
    // if err != nil {
	// 	return nil, fmt.Errorf("Failed to call contract: %w", err)
    // }

    // // Process the output data (this depends on what the function returns)
    // results, err := parsedABI.Unpack("quoteExactInputSingleV2", output)
    // if err != nil {
	// 	return nil, fmt.Errorf("Failed to unpack function output: %w", err)
    // }
    // fmt.Println("Output from contract call:", results)

    // // Type assert the result as *big.Int
    // result, ok := results[0].(*big.Int)
    // if !ok {
    //     return nil, fmt.Errorf("result type assertion to *big.Int failed")
    // }

    // return result, nil

	return nil, fmt.Errorf("Not yet implemented")
}

func (sdk *SdkOptimism) CreateDepositTransaction(fromAddress common.Address, amount *big.Int, referral *big.Int) (string, error) {
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

    abiPath := "pkg/sdk/optimism/abis/NovaVault.json"
    file, err := os.ReadFile(abiPath)
    if err != nil {
		return "", fmt.Errorf("Failed to read ABI file: %w", err)
    }

    parsedABI, err := abi.JSON(strings.NewReader(string(file)))
    if err != nil {
		return "", fmt.Errorf("Failed to parse ABI: %w", err)
    }

    contractAddress := common.HexToAddress(sdk.Config.VaultAddress)

	data, err := parsedABI.Pack("deposit", amount, fromAddress, referral)
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

func (sdk *SdkOptimism) CreateWithdrawTransaction(fromAddress common.Address, amount *big.Int, referral *big.Int) (string, error) {
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

    abiPath := "pkg/sdk/optimism/abis/SavingsDai.json"
    file, err := os.ReadFile(abiPath)
    if err != nil {
		return "", fmt.Errorf("Failed to read ABI file: %w", err)
    }

    parsedABI, err := abi.JSON(strings.NewReader(string(file)))
    if err != nil {
		return "", fmt.Errorf("Failed to parse ABI: %w", err)
    }

    contractAddress := common.HexToAddress(sdk.Config.VaultAddress)

	data, err := parsedABI.Pack("withdraw", amount, fromAddress, fromAddress)
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
