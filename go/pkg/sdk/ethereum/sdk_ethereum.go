package ethereum

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/NovaSubDAO/nova-sdk/go/pkg/config"
	"github.com/NovaSubDAO/nova-sdk/go/pkg/contracts"
	ethereumContracts "github.com/NovaSubDAO/nova-sdk/go/pkg/sdk/ethereum/abis"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type SdkEthereum struct {
	Config   *config.Config
	Contract *contracts.ContractsCaller
}

func NewSdkEthereum(cfg *config.Config) (*SdkEthereum, error) {
	client, err := ethclient.Dial(cfg.RpcEndpoint)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to Ethereum client: %w", err)
	}

	contract, err := contracts.NewContractsCaller(common.HexToAddress(cfg.VaultAddress), client)
	if err != nil {
		return nil, fmt.Errorf("Failed to instantiate contract caller: %w", err)
	}

	return &SdkEthereum{Config: cfg, Contract: contract}, nil
}

func (sdk *SdkEthereum) GetPrice() (*big.Int, error) {
	factor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(sdk.Config.VaultDecimals)), nil)

	assets, err := sdk.Contract.ConvertToAssets(nil, factor)
	if err != nil {
		return big.NewInt(0), err
	}

	return assets, nil
}

func (sdk *SdkEthereum) GetPosition(address common.Address) (*big.Int, error) {
	balance, err := sdk.Contract.BalanceOf(nil, address)
	if err != nil {
		return big.NewInt(0), err
	}

	price, err := sdk.GetPrice()
	if err != nil {
		return big.NewInt(0), err
	}

	value := new(big.Int).Mul(balance, price)
	factor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(sdk.Config.VaultDecimals)), nil)
	valueNormalized := new(big.Int).Div(value, factor)

	return valueNormalized, nil
}

func (sdk *SdkEthereum) GetTotalValue() (*big.Int, error) {
	totalSupply, err := sdk.Contract.TotalSupply(nil)
	if err != nil {
		return big.NewInt(0), err
	}

	price, err := sdk.GetPrice()
	if err != nil {
		return big.NewInt(0), err
	}

	totalValue := new(big.Int).Mul(totalSupply, price)
	factor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(sdk.Config.VaultDecimals)), nil)
	totalValueNormalized := new(big.Int).Div(totalValue, factor)

	return totalValueNormalized, nil
}

func (sdk *SdkEthereum) GetSlippage(amount *big.Int) (*big.Int, error) {
	return nil, fmt.Errorf("Not yet implemented")
}

func (sdk *SdkEthereum) CreateDepositTransaction(fromAddress common.Address, stable common.Address, amount *big.Int, referral *big.Int) (string, error) {
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

	contractAddress := common.HexToAddress(sdk.Config.VaultAddress)

	contractAbi, err := abi.JSON(strings.NewReader(ethereumContracts.SavingsDaiABI))
	if err != nil {
		return "", fmt.Errorf("Failed to parse contract ABI: %w", err)
	}

	data, err := contractAbi.Pack("deposit", amount, fromAddress)
	if err != nil {
		return "", fmt.Errorf("failed to create deposit message: %w", err)
	}

	// Estimating the gas needed for the transaction
	call := ethereum.CallMsg{From: fromAddress, To: &contractAddress, GasPrice: gasPrice, Value: big.NewInt(0), Data: data}
	gasLimit, err := client.EstimateGas(context.Background(), call)
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

func (sdk *SdkEthereum) CreateWithdrawTransaction(fromAddress common.Address, stable common.Address, amount *big.Int, referral *big.Int) (string, error) {
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

	contractAddress := common.HexToAddress(sdk.Config.VaultAddress)

	contractAbi, err := abi.JSON(strings.NewReader(ethereumContracts.SavingsDaiABI))
	if err != nil {
		return "", fmt.Errorf("failed to parse contract ABI: %w", err)
	}

	data, err := contractAbi.Pack("withdraw", amount, fromAddress, fromAddress)
	if err != nil {
		return "", fmt.Errorf("failed to pack withdraw message: %w", err)
	}

	// Estimating the gas needed for the transaction
	call := ethereum.CallMsg{From: fromAddress, To: &contractAddress, GasPrice: gasPrice, Value: big.NewInt(0), Data: data}
	gasLimit, err := client.EstimateGas(context.Background(), call)
	if err != nil {
		log.Printf("Gas estimation failed, using fallback gas limit: %v", err)
		gasLimit = 2000000 // Fallback gas limit
	}

	tx := types.NewTransaction(nonce, contractAddress, big.NewInt(0), gasLimit, gasPrice, data)

	txJSON, err := json.Marshal(tx)
	if err != nil {
		return "", fmt.Errorf("failed to marshal transaction: %w", err)
	}

	return string(txJSON), nil
}

func (sdk *SdkEthereum) Deposit(assets *big.Int, receiver common.Address, referral big.Int) (*types.Transaction, error) {
	// cfg, err := sdk.Config.LoadConfig()
	// if err != nil {
	// 	return nil, fmt.Errorf("Error loading configuration: %w", err)
	// }

	// client, err := ethclient.Dial(cfg.RpcEndpoint)
	// if err != nil {
	// 	return nil, fmt.Errorf("Failed to connect to Ethereum client: %w", err)
	// }

	// contract, err := contracts.NewContractsCaller(common.HexToAddress(cfg.VaultAddress), client)
	// if err != nil {
	// 	return nil, fmt.Errorf("Failed to instantiate contract caller: %w", err)
	// }

	return nil, fmt.Errorf("Not yet implemented")
}

func (sdk *SdkEthereum) Withdraw(assets *big.Int, receiver common.Address, referral big.Int) (*types.Transaction, error) {
	return nil, fmt.Errorf("Not yet implemented")
}
