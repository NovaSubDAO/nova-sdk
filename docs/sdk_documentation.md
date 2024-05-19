The SDK is made of 4 functions:

- **deposit**: a function to deposit your stablecoins to Nova contract,
- **withdraw**: a function to withdraw your stablecoins from Nova contract,
- **getSDaiPrice**: a function to know the current price of sDAI in DAI,
- **getSplipagge**: a function to know the current slippage of the underlying swap.

## **deposit** function

=== "Python"

````python
from nova_sdk import deposit

    deposit(amount, stablecoin)
    ```

=== "Go"
```go
package main

    import github.com/NovaSubDAO/nova-sdk/go/pkg/sdk

    sdk.deposit(amount, stablecoin)
    ```
````
