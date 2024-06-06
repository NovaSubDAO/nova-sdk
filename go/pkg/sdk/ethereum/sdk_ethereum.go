package ethereum

import (
    "math/big"
    "fmt"

    "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
    "github.com/NovaSubDAO/nova-sdk/go/pkg/config"
)

type SdkEthereum struct {
    Config *config.Config
}

func NewSdkEthereum(cfg *config.Config) *SdkEthereum {
    return &SdkEthereum{Config: cfg}
}

func (sdk *SdkEthereum) GetPrice() (*big.Int, error) {
    // cfg, err := sdk.Config.LoadConfig()
    // if err != nil {
	// 	return nil, fmt.Errorf("Error loading configuration: %w", err)
	// }

    // factor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(cfg.VaultDecimals)), nil)

    // assets, err := sdk.Contract.ConvertToAssets(nil, factor)
    // if err != nil {
    //     return big.NewInt(0), err
    // }

    // return assets, nil

	return nil, fmt.Errorf("Not yet implemented")
}

func (sdk *SdkEthereum) GetPosition(address common.Address) (*big.Int, error) {
    // cfg, err := sdk.Config.LoadConfig()
    // if err != nil {
	// 	return nil, fmt.Errorf("Error loading configuration: %w", err)
	// }

    // balance, err := sdk.Contract.BalanceOf(nil, address)
    // if err != nil {
    //     return big.NewInt(0), err
    // }

    // price, err := sdk.GetPrice()
    // if err != nil {
    //     return big.NewInt(0), err
    // }

    // value := new(big.Int).Mul(balance, price)
    // factor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(cfg.VaultDecimals)), nil)
    // valueNormalized := new(big.Int).Div(value, factor)

    // return valueNormalized, nil

	return nil, fmt.Errorf("Not yet implemented")
}

func (sdk *SdkEthereum) GetTotalValue() (*big.Int, error) {
    // cfg, err := sdk.Config.LoadConfig()
    // if err != nil {
	// 	return nil, fmt.Errorf("Error loading configuration: %w", err)
	// }

    // totalSupply, err := sdk.Contract.TotalSupply(nil)
    // if err != nil {
    //     return big.NewInt(0), err
    // }

    // price, err := sdk.GetPrice()
    // if err != nil {
    //     return big.NewInt(0), err
    // }

    // totalValue := new(big.Int).Mul(totalSupply, price)
    // factor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(cfg.VaultDecimals)), nil)
    // totalValueNormalized := new(big.Int).Div(totalValue, factor)

    // return totalValueNormalized, nil

	return nil, fmt.Errorf("Not yet implemented")
}

func (sdk *SdkEthereum) GetSlippage(amount *big.Int) (*big.Int, error) {
	return nil, fmt.Errorf("Not yet implemented")
}

func (sdk *SdkEthereum) CreateDepositTransaction(big.Int, big.Int) (string, error) {
	return "", fmt.Errorf("Not yet implemented")
}

func (sdk *SdkEthereum) CreateWithdrawTransaction(big.Int, big.Int) (string, error) {
	return "", fmt.Errorf("Not yet implemented")
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
