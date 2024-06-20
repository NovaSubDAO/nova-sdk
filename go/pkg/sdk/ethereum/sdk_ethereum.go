package ethereum

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/NovaSubDAO/nova-sdk/go/pkg/config"
	"github.com/NovaSubDAO/nova-sdk/go/pkg/constants"
	ethereumContracts "github.com/NovaSubDAO/nova-sdk/go/pkg/sdk/ethereum/abis"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type SdkEthereum struct {
	Config   *config.Config
	Contract *ethereumContracts.SavingsDaiCaller
}

func NewSdkEthereum(cfg *config.Config) (*SdkEthereum, error) {
	client, err := ethclient.Dial(cfg.RpcEndpoint)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to Ethereum client: %w", err)
	}

	contract, err := ethereumContracts.NewSavingsDaiCaller(common.HexToAddress(cfg.SDai), client)
	if err != nil {
		return nil, fmt.Errorf("Failed to instantiate contract caller: %w", err)
	}

	return &SdkEthereum{Config: cfg, Contract: contract}, nil
}

func (sdk *SdkEthereum) isStablecoinSupported(stable constants.Stablecoin) bool {
	if stablecoins, ok := constants.StablecoinDetails[sdk.Config.ChainId]; ok {
		_, exists := stablecoins[stable]
		return exists
	}
	return false
}

func (sdk *SdkEthereum) GetConfig() *config.Config {
	return sdk.Config
}

func (sdk *SdkEthereum) GetPrice(stable constants.Stablecoin) (*big.Int, error) {
	ok := sdk.isStablecoinSupported(stable)
	if !ok {
		return big.NewInt(0), fmt.Errorf("stablecoin %s is not supported", stable)
	}

	price, err := sdk.GetSDaiPrice()
	if err != nil {
		return big.NewInt(0), err
	}
	return price, nil
}

func (sdk *SdkEthereum) GetPosition(stable constants.Stablecoin, address common.Address) (*big.Int, error) {
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

func (sdk *SdkEthereum) GetSDaiPrice() (*big.Int, error) {
	factor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(sdk.Config.VaultDecimals)), nil)
	assets, err := sdk.Contract.ConvertToAssets(nil, factor)
	if err != nil {
		return big.NewInt(0), err
	}
	return assets, nil
}

func (sdk *SdkEthereum) GetSDaiTotalValue() (*big.Int, error) {
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

func (sdk *SdkEthereum) GetSupportedStablecoins() ([]constants.Stablecoin, error) {
	if stablecoins, ok := constants.StablecoinDetails[sdk.Config.ChainId]; ok {
		var supportedStablecoins []constants.Stablecoin
		for sc := range stablecoins {
			supportedStablecoins = append(supportedStablecoins, sc)
		}
		return supportedStablecoins, nil
	}
	return nil, fmt.Errorf("no stablecoins found for chain ID %d", sdk.Config.ChainId)
}

func (sdk *SdkEthereum) GetSlippage(stable constants.Stablecoin, amount *big.Int) (float64, float64, float64, error) {
	ok := sdk.isStablecoinSupported(stable)
	if !ok {
		return 0, 0, 0, fmt.Errorf("stablecoin %s is not supported", stable)
	}

	price, err := sdk.GetSDaiPrice()
	if err != nil {
		return 0, 0, 0, err
	}

	sDaiDecimals := sdk.Config.VaultDecimals
	base := big.NewInt(10)

	sDaiTenPow := new(big.Int).Exp(base, big.NewInt(int64(sDaiDecimals)), nil)
	sDaiTenPowFloat := new(big.Float).SetInt(sDaiTenPow)

	priceFloat := new(big.Float).SetInt(price)
	priceFloat.Quo(priceFloat, sDaiTenPowFloat)

	priceFloat64, _ := priceFloat.Float64()

	return float64(0), priceFloat64, priceFloat64, nil
}

func (sdk *SdkEthereum) CreateDepositTransaction(stable constants.Stablecoin, fromAddress common.Address, amount *big.Int, referral *big.Int) (string, error) {
	ok := sdk.isStablecoinSupported(stable)
	if !ok {
		return "", fmt.Errorf("stablecoin %s is not supported", stable)
	}

	stableAddress := common.HexToAddress(constants.StablecoinDetails[sdk.Config.ChainId][stable].Address)
	vaultAddress := common.HexToAddress(sdk.Config.VaultAddress)

	client, err := ethclient.Dial(sdk.Config.RpcEndpoint)
	if err != nil {
		return "", fmt.Errorf("Failed to connect to the Ethereum client: %v", err)
	}

	stableContract, err := ethereumContracts.NewDaiCaller(stableAddress, client)
	if err != nil {
		return "", fmt.Errorf("Failed to load Dai contract: %w", err)
	}

	result, err := stableContract.Allowance(nil, fromAddress, vaultAddress)
	if err != nil {
		return "", fmt.Errorf("Failed to call Allowance function: %w", err)
	}

	if amount.Cmp(result) > 0 {
		return "", fmt.Errorf("Allowance is too low. First call approve function on DAI contract.")
	}

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", fmt.Errorf("Failed to get nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("Failed to suggest gas price: %v", err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(ethereumContracts.SavingsDaiABI))
	if err != nil {
		return "", fmt.Errorf("Failed to parse contract ABI: %w", err)
	}

	data, err := contractAbi.Pack("deposit", amount, fromAddress)
	if err != nil {
		return "", fmt.Errorf("failed to create deposit message: %w", err)
	}

	// Estimating the gas needed for the transaction
	call := ethereum.CallMsg{From: fromAddress, To: &vaultAddress, GasPrice: gasPrice, Value: big.NewInt(0), Data: data}
	gasLimit, err := client.EstimateGas(context.Background(), call)
	if err != nil {
		log.Printf("Gas estimation failed, using fallback gas limit: %v", err)
		gasLimit = 2000000 // Fallback gas limit
	}

	tx := types.NewTransaction(nonce, vaultAddress, big.NewInt(0), gasLimit, gasPrice, data)

	txJSON, err := json.Marshal(tx)
	if err != nil {
		return "", fmt.Errorf("Failed to marshal transaction: %w", err)
	}

	return string(txJSON), nil
}

func (sdk *SdkEthereum) CreateWithdrawTransaction(stable constants.Stablecoin, fromAddress common.Address, amount *big.Int, referral *big.Int) (string, error) {
	ok := sdk.isStablecoinSupported(stable)
	if !ok {
		return "", fmt.Errorf("stablecoin %s is not supported", stable)
	}

	client, err := ethclient.Dial(sdk.Config.RpcEndpoint)
	if err != nil {
		return "", fmt.Errorf("Failed to connect to the Ethereum client: %v", err)
	}

	vaultAddress := common.HexToAddress(sdk.Config.VaultAddress)

	sDaiContract, err := ethereumContracts.NewSavingsDaiCaller(vaultAddress, client)
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

	contractAbi, err := abi.JSON(strings.NewReader(ethereumContracts.SavingsDaiABI))
	if err != nil {
		return "", fmt.Errorf("failed to parse contract ABI: %w", err)
	}

	data, err := contractAbi.Pack("withdraw", amount, fromAddress, fromAddress)
	if err != nil {
		return "", fmt.Errorf("failed to pack withdraw message: %w", err)
	}

	// Estimating the gas needed for the transaction
	call := ethereum.CallMsg{From: fromAddress, To: &vaultAddress, GasPrice: gasPrice, Value: big.NewInt(0), Data: data}
	gasLimit, err := client.EstimateGas(context.Background(), call)
	if err != nil {
		log.Printf("Gas estimation failed, using fallback gas limit: %v", err)
		gasLimit = 2000000 // Fallback gas limit
	}

	tx := types.NewTransaction(nonce, vaultAddress, big.NewInt(0), gasLimit, gasPrice, data)

	txJSON, err := json.Marshal(tx)
	if err != nil {
		return "", fmt.Errorf("failed to marshal transaction: %w", err)
	}

	return string(txJSON), nil
}
