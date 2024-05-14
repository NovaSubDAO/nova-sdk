package tests

import (
    "testing"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/NovaSubDAO/nova-sdk/go/pkg/config"
    "github.com/NovaSubDAO/nova-sdk/go/pkg/sdk"
)

func setup(t *testing.T) (*sdk.NovaSDK, func()) {
    // Load configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        t.Fatalf("Error loading configuration: %v", err)
    }

    // Initialize the Ethereum client
    ethClient, err := ethclient.Dial(cfg.RpcEndpoint)
    if err != nil {
        t.Fatalf("Failed to connect to Ethereum client: %v", err)
    }

    // Initialize the SDK
    novaSDK, err := sdk.NewNovaSDK(ethClient, cfg.SDaiAddress)
    if err != nil {
        t.Fatalf("Failed to initialize the Nova SDK: %v", err)
    }

    // Return the initialized SDK and a cleanup function to close the client
    return novaSDK, func() {
        ethClient.Close()
    }
}

func TestGetSDaiPrice(t *testing.T) {
    novaSDK, cleanup := setup(t)
    defer cleanup()

    // Use the SDK to get sDAI price
    output, err := novaSDK.GetSDaiPrice()
    if err != nil {
        t.Fatalf("Failed to get sDAI price: %v", err)
    }

    // Output the result (you can add more assertions as needed)
    t.Logf("Output from GetSDaiPrice: %f", output)
}

func TestGetPosition(t *testing.T) {
    novaSDK, cleanup := setup(t)
    defer cleanup()

    // Use the SDK to get the position of a user
    addressStr := "0x4aa42145aa6ebf72e164c9bbc74fbd3788045016"
    address := common.HexToAddress(addressStr)
    output, err := novaSDK.GetPosition(address)
    if err != nil {
        t.Fatalf("Failed to get position: %v", err)
    }

    // Output the result (you can add more assertions as needed)
    t.Logf("Output from GetPosition: %f", output)
}
