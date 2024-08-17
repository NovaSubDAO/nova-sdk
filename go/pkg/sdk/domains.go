package sdk

import (
	"math/big"

	"github.com/NovaSubDAO/nova-sdk/go/pkg/config"
	"github.com/NovaSubDAO/nova-sdk/go/pkg/constants"
	"github.com/ethereum/go-ethereum/common"
)

type SdkInterface interface {
	GetPrice(constants.Stablecoin) (*big.Int, error)
	GetPosition(constants.Stablecoin, common.Address) (*big.Int, error)
	GetSlippage(constants.Stablecoin, *big.Int) (float64, float64, float64, error)

	GetConfig() *config.Config
	GetSDaiPrice() (*big.Int, error)
	GetSDaiTotalValue() (*big.Int, error)
	GetSupportedStablecoins() ([]constants.Stablecoin, error)

	// AddDex(common.Address, common.Address) (string, error)
	// SetFunctionApprovalBySignature(common.Address, string, [4]byte) (string, error)

	CreateDepositTransaction(constants.Stablecoin, common.Address, *big.Int, *big.Int) (string, error)
	CreateWithdrawTransaction(constants.Stablecoin, common.Address, *big.Int, *big.Int) (string, error)
}
