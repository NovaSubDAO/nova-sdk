package sdk

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type SdkInterface interface {
	GetPrice() (*big.Int, error)
	GetPosition(common.Address) (*big.Int, error)
	GetTotalValue() (*big.Int, error)
	GetSlippage(common.Address, *big.Int) (*big.Int, error)

	CreateDepositTransaction(common.Address, common.Address, *big.Int, *big.Int) (string, error)
	CreateWithdrawTransaction(common.Address, common.Address, *big.Int, *big.Int) (string, error)

	Deposit(*big.Int, common.Address, big.Int) (*types.Transaction, error)
	Withdraw(*big.Int, common.Address, big.Int) (*types.Transaction, error)
}
