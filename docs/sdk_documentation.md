# SDK Documentation

The NovaVault SDK for Go enables developers to interact with the NovaVault contract for managing stablecoins.
The SDK simplifies tasks like depositing, withdrawing, and querying financial data within a liquidity pool.

## Overview of Functions

This SDK implements the follow interface:

```go
type SdkInterface interface {
	GetPrice(constants.Stablecoin) (*big.Int, error)
	GetPosition(constants.Stablecoin, common.Address) (*big.Int, error)
	GetSlippage(constants.Stablecoin, *big.Int) (float64, error)

	GetSDaiPrice() (*big.Int, error)
	GetSDaiTotalValue() (*big.Int, error)

	CreateDepositTransaction(constants.Stablecoin, common.Address, *big.Int, *big.Int) (string, error)
	CreateWithdrawTransaction(constants.Stablecoin, common.Address, *big.Int, *big.Int) (string, error)
}
```
- **GetPrice**: Retrieve the current price of sDAI in the liquidity pool stable/sDAI.
- **GetPosition**: Retrieve the value of the position held by the user.
- **GetSlippage**: Calculate the current slippage percentage for a swap transaction, based on the specified amount to swap.
- **GetSDaiPrice**: Retrieve the current price of sDAI on Ethereum.
- **GetSDaiTotalValue**: Retrieve the total value of sDAI in the domain, using Ethereum price of sDAI.
- **CreateDepositTransaction**: Creates a deposit transaction with referral code, to be signed by the user.
- **CreateWithdrawTransaction**: Creates a withdraw transaction with referral code, to be signed by the user.


## Function Details

### **Deposit**

Deposits a specified amount of stablecoins into the NovaVault.
This function requires the stablecoin type, amount, and an optional referral identifier.

#### Parameters

- `amount` (int): The amount of stablecoin to deposit.
- `stablecoin` (string): The type of stablecoin to deposit.
- `referral` (int, optional): A referral code for tracking or rewards.

#### Example

```go
package main

import (
    "github.com/NovaSubDAO/nova-sdk/go/pkg/sdk"
    "fmt"
)

func main() {
    err := sdk.Deposit(100, "USDC", 123)
    if err != nil {
        fmt.Println("Error depositing stablecoin:", err)
    }
}
```

### **Withdraw**

Withdraws a specified amount of stablecoins from the NovaVault. It requires the stablecoin type and amount.

#### Parameters

- `amount` (int): The amount of stablecoin to withdraw.
- `stablecoin` (string): The type of stablecoin to withdraw.

#### Example

```go
package main

import (
    "github.com/NovaSubDAO/nova-sdk/go/pkg/sdk"
    "fmt"
)

func main() {
    err := sdk.Withdraw(50, "USDC")
    if err != nil {
        fmt.Println("Error withdrawing stablecoin:", err)
    }
}
```

### **GetPrice**

Fetches the current price of sDAI from the liquidity pool.

#### Example

```go
package main

import (
    "github.com/NovaSubDAO/nova-sdk/go/pkg/sdk"
    "fmt"
)

func main() {
    price, err := sdk.GetPrice()
    if err != nil {
        fmt.Println("Error getting price:", err)
        return
    }
    fmt.Println("Current price of sDAI:", price)
}
```

### **GetSlippage**

Calculates the slippage for a given transaction amount in the liquidity pool.

#### Parameters

- `amount` (int): The transaction amount for which to calculate slippage.

#### Example

```go
package main

import (
    "github.com/NovaSubDAO/nova-sdk/go/pkg/sdk"
    "fmt"
)

func main() {
    slippage, err := sdk.GetSlippage(100)
    if err != nil {
        fmt.Println("Error calculating slippage:", err)
        return
    }
    fmt.Println("Estimated slippage:", slippage)
}
```
