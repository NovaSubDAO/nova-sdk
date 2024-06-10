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


## Function Details

### GetPrice

Retrieves the current price of a specified stablecoin in the liquidity pool stable/sDAI.

#### Parameters

- `stablecoin (constants.Stablecoin)`: The type of stablecoin for which to retrieve the price.

#### Returns

- `price (*big.Int)`: The current price of the specified stablecoin.
- `error`: Error if the price retrieval fails.

#### Example

```go
package main

import (
    "fmt"
    "math/big"
    "github.com/NovaSubDAO/nova-sdk/go/pkg/sdk"
)

func main() {
    price, err := novaSdk.SdkDomain.GetPrice("USDC")
    if err != nil {
        fmt.Println("Error retrieving price:", err)
    } else {
        fmt.Println("Price:", price)
    }
}
```

### GetPosition

Retrieves the value of the position held by a specified user address in the given stablecoin.

#### Parameters

- `stablecoin` (`constants.Stablecoin`): The type of stablecoin for which to retrieve the position.
- `address` (`common.Address`): The user's wallet address.

#### Returns

- `position` (`*big.Int`): The current value of the user's position.
- `error`: Error if the position retrieval fails.

#### Example

```go
package main

import (
    "fmt"
    "github.com/ethereum/go-ethereum/common"
    "github.com/NovaSubDAO/nova-sdk/go/pkg/sdk"
)

func main() {
    address := common.HexToAddress("0x123...")
    position, err := novaSdk.SdkDomain.GetPosition("USDC", address)
    if err != nil {
        fmt.Println("Error retrieving position:", err)
    } else {
        fmt.Println("Position value:", position)
    }
}
```

### GetSlippage

Calculates the current slippage percentage for a swap transaction based on the specified amount to swap.

#### Parameters

- `stablecoin` (`constants.Stablecoin`): The type of stablecoin involved in the swap.
- `amount` (`*big.Int`): The amount of stablecoin to be swapped.

#### Returns

- `slippage` (`float64`): The calculated slippage percentage.
- `error`: Error if the slippage calculation fails.

#### Example

```go
package main

import (
    "fmt"
    "math/big"
    "github.com/NovaSubDAO/nova-sdk/go/pkg/sdk"
)

func main() {
    rpcEndpoint := "https://your-rpc-endpoint"
    chainId := int64(123)
	novaSdk, err := sdk.NewNovaSDK(rpcEndpoint, chainId)

    amountToSwap := big.NewInt(1000) // example amount to swap
    slippage, err := novaSdk.SdkDomain.GetSlippage("USDC", amountToSwap)
    if err != nil {
        fmt.Println("Error calculating slippage:", err)
    } else {
        fmt.Printf("Slippage percentage: %.2f%%\n", slippage)
    }
}
```

### GetSDaiPrice

Retrieves the current price of sDAI on Ethereum.

#### Returns

- `price` (*big.Int): The current price of sDAI.
- `error`: Error if the price retrieval fails.

#### Example

```go
package main

import (
    "fmt"
    "github.com/NovaSubDAO/nova-sdk/go/pkg/sdk"
)

func main() {
    rpcEndpoint := "https://your-rpc-endpoint"
    chainId := int64(123)
	novaSdk, err := sdk.NewNovaSDK(rpcEndpoint, chainId)

    price, err := novaSdk.SdkDomain.GetSDaiPrice()
    if err != nil {
        fmt.Println("Error retrieving sDAI price:", err)
    } else {
        fmt.Println("sDAI Price:", price)
    }
}
```

### GetSDaiTotalValue

Retrieves the total value of sDAI in the domain, using the Ethereum price of sDAI.

#### Returns

- `totalValue` (*big.Int): The total value of sDAI.
- `error`: Error if the value retrieval fails.

#### Example

```go
package main

import (
    "fmt"
    "github.com/NovaSubDAO/nova-sdk/go/pkg/sdk"
)

func main() {
    rpcEndpoint := "https://your-rpc-endpoint"
    chainId := int64(123)
	novaSdk, err := sdk.NewNovaSDK(rpcEndpoint, chainId)

    totalValue, err := novaSdk.SdkDomain.GetSDaiTotalValue()
    if err != nil {
        fmt.Println("Error retrieving total value of sDAI:", err)
    } else {
        fmt.Println("Total value of sDAI:", totalValue)
    }
}
```

### CreateDepositTransaction

Creates a deposit transaction for a specified stablecoin to be signed by the user. This function requires the stablecoin type, the user wallet address, the amount to deposit, and an optional referral code.

#### Parameters

- `stablecoin` (`constants.Stablecoin`): The type of stablecoin to deposit.
- `address` (`common.Address`): The user's wallet address.
- `amount` (`*big.Int`): The amount of stablecoin to deposit.
- `referral` (`*big.Int`, optional): A referral code for tracking or rewards, if applicable.

#### Returns

- `transactionObject` (string): Deposit transaction object, to be signed by the user.
- `error`: Error if the transaction creation fails.

#### Example

```go
package main

import (
    "fmt"
    "math/big"
    "github.com/ethereum/go-ethereum/common"
    "github.com/NovaSubDAO/nova-sdk/go/pkg/sdk"
	"github.com/NovaSubDAO/nova-sdk/go/pkg/constants"
)

func main() {
    rpcEndpoint := "https://your-rpc-endpoint"
    chainId := int64(123)
	novaSdk, err := sdk.NewNovaSDK(rpcEndpoint, chainId)

    stablecoin := constants.USDC
    userAddress := common.HexToAddress("0x123...")
    depositAmount := big.NewInt(1000)
    referralCode := big.NewInt(101)

    transactionObject, err := novaSdk.SdkDomain.CreateDepositTransaction(stablecoin, userAddress, depositAmount, referralCode)
    if err != nil {
        fmt.Println("Error creating deposit transaction:", err)
    } else {
        fmt.Println("Deposit transaction to be signed:", transactionObject)
    }
}
```

### **CreateWithdrawTransaction**

Creates a withdraw transaction for a specified stablecoin to be signed by the user. This function requires the stablecoin type, the user wallet address, the amount to withdraw, and an optional referral code.

#### Parameters

- `stablecoin` (`constants.Stablecoin`): The type of stablecoin to withdraw.
- `address` (`common.Address`): The user's wallet address.
- `amount` (`*big.Int`): The amount of stablecoin to withdraw.
- `referral` (`*big.Int`, optional): A referral code for tracking or rewards, if applicable.

#### Returns

- `transactionObject` (string): Withdraw transaction object, to be signed by the user.
- `error`: Error if the transaction creation fails.

#### Example

```go
package main

import (
    "fmt"
    "math/big"
    "github.com/ethereum/go-ethereum/common"
    "github.com/NovaSubDAO/nova-sdk/go/pkg/sdk"
	"github.com/NovaSubDAO/nova-sdk/go/pkg/constants"
)

func main() {
    rpcEndpoint := "https://your-rpc-endpoint"
    chainId := int64(123)
	novaSdk, err := sdk.NewNovaSDK(rpcEndpoint, chainId)

    stablecoin := constants.USDC
    userAddress := common.HexToAddress("0x123...")
    withdrawAmount := big.NewInt(500)
    referralCode := big.NewInt(102)

    transactionObject, err := novaSdk.SdkDomain.CreateWithdrawTransaction(stablecoin, userAddress, withdrawAmount, referralCode)
    if err != nil {
        fmt.Println("Error creating withdraw transaction:", err)
    } else {
        fmt.Println("Withdraw transaction to be signed:", transactionObject)
    }
}
```