package sdk

import (
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/NovaSubDAO/nova-sdk/go/pkg/config"
	"github.com/NovaSubDAO/nova-sdk/go/pkg/contracts"
)

type NovaSDK struct {
    Contract *contracts.ContractsCaller
    Client   *ethclient.Client
}

func NewNovaSDK(client *ethclient.Client, contractAddress string) (*NovaSDK, error) {
    contract, err := contracts.NewContractsCaller(common.HexToAddress(contractAddress), client)
    if err != nil {
        return nil, err
    }

    return &NovaSDK{
        Contract: contract,
        Client:   client,
    }, nil
}

func (sdk *NovaSDK) GetSDaiPrice() (float64, error) {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatal("Error loading configuration:", err)
    }

    one := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(cfg.SDaiDecimals)), nil)

    assets, err :=sdk.Contract.ConvertToAssets(nil, one)
    if err != nil {
        return 0, err
    }

    assetsFloat := new(big.Float).SetInt(assets)
    oneFloat := new(big.Float).SetInt(one)

    price := new(big.Float).Quo(assetsFloat, oneFloat)
    priceFloat, _ := price.Float64()

    return priceFloat, nil
}

func (sdk *NovaSDK) GetPosition(address common.Address) (float64, error) {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatal("Error loading configuration:", err)
    }

    balance, err := sdk.Contract.BalanceOf(nil, address)
    if err != nil {
        return 0, err
    }

    price, err := sdk.GetSDaiPrice()
    if err != nil {
        return 0, err
    }

    balanceFloat := new(big.Float).SetInt(balance)
    totalValue := new(big.Float).Mul(balanceFloat, big.NewFloat(price))

    one := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(cfg.SDaiDecimals)), nil)
    oneFloat := new(big.Float).SetInt(one)

    position := new(big.Float).Quo(totalValue, oneFloat)
    positionFloat, _ := position.Float64()

    return positionFloat, nil
}
