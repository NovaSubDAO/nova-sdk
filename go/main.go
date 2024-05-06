package main

import (
    "fmt"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum/ethclient"
	"github.com/NovaSubDAO/nova-python-sdk/go/pkg/config"
	"github.com/NovaSubDAO/nova-python-sdk/go/pkg/sdk"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatal("Error loading configuration:", err)
    }

    // Initialize the Ethereum client
    ethClient, err := ethclient.Dial(cfg.RpcEndpoint)
    if err != nil {
        log.Fatalf("Failed to connect to Ethereum client: %v", err)
    }
    defer ethClient.Close()

    // Initialize the SDK
    novaSDK, err := sdk.NewNovaSDK(ethClient, cfg.SDaiAddress)
    if err != nil {
        log.Fatalf("Failed to initialize the Nova SDK: %v", err)
	}

    // Use the SDK to convert assets
    bigNum := big.NewInt(1e18)
    output, err := novaSDK.ConvertToAssets(bigNum)
    if err != nil {
        log.Fatalf("Failed to convert assets: %v", err)
    }

    // Output the result
    fmt.Printf("Output from ConvertToAssets: %s\n", output)
}
