package sdk

import (
    "math/big"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
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

func (sdk *NovaSDK) ConvertToAssets(amount *big.Int) (*big.Int, error) {
    return sdk.Contract.ConvertToAssets(nil, amount)
}
