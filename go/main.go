package main

import (
    "log"

    "github.com/NovaSubDAO/nova-sdk/go/pkg/sdk"
)

func main() {
    // check vault address on ethereum
    chainId := int64(1)
    novaSdk, err := sdk.NewNovaSDK("", chainId)
    if err != nil {
        log.Fatal("Error loading configuration:", err)
    }
    log.Println("chain id:", novaSdk.Config.ChainId)
    log.Println("vault address:", novaSdk.Config.VaultAddress)

    // check vault address on optimism
    chainId = 10
    novaSdk, err = sdk.NewNovaSDK("", chainId)
    if err != nil {
        log.Fatal("Error loading configuration:", err)
    }
    log.Println("chain id:", novaSdk.Config.ChainId)
    log.Println("vault address:", novaSdk.Config.VaultAddress)
}
