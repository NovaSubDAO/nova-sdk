package sdk

import (
	"math/big"

	"github.com/NovaSubDAO/nova-sdk/go/pkg/constants"
	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/core/types"
)

type SdkInterface interface {
	GetPrice(constants.Stablecoin) (*big.Int, error)
	GetPosition(constants.Stablecoin, common.Address) (*big.Int, error)
	GetSlippage(constants.Stablecoin, *big.Int) (float64, float64, float64, error)

	GetSDaiPrice() (*big.Int, error)
	GetSDaiTotalValue() (*big.Int, error)

	CreateDepositTransaction(constants.Stablecoin, common.Address, *big.Int, *big.Int) (string, error)
	CreateWithdrawTransaction(constants.Stablecoin, common.Address, *big.Int, *big.Int) (string, error)

	// Deposit(constants.Stablecoin, *big.Int, common.Address, big.Int) (*types.Transaction, error)
	// Withdraw(constants.Stablecoin, *big.Int, common.Address, big.Int) (*types.Transaction, error)
}
