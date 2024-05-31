package main

import (
    "log"

    "github.com/NovaSubDAO/nova-sdk/go/pkg/config"
    "github.com/NovaSubDAO/nova-sdk/go/pkg/sdk"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatal("Error loading configuration:", err)
    }
    log.Println("sDAI address:", cfg.VaultAddress)

    // Ensure sdk is properly initialized
    client, cleanup := sdk.NewNovaSDK(cfg.VaultAddress)
	defer cleanup()

    result, err := client.GetPrice()
    if err != nil {
        log.Fatal("Error getting price:", err)
    }
    log.Println("Price:", result)
}
