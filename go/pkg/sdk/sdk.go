package sdk

import (
    "fmt"

    "github.com/NovaSubDAO/nova-sdk/go/pkg/config"
    "github.com/NovaSubDAO/nova-sdk/go/pkg/sdk/ethereum"
    "github.com/NovaSubDAO/nova-sdk/go/pkg/sdk/optimism"
)

type NovaSDK struct {
    Config   *config.Config
    SdkDomain SdkInterface
}

func NewNovaSDK(rpcEndpoint string, chainId int64) (*NovaSDK, error) {
    cfg, err := config.LoadConfig(rpcEndpoint, chainId)
    if err != nil {
        return nil, err
    }

    var sdkDomain SdkInterface
    switch chainId {
    case 1:
        sdkDomain, err = ethereum.NewSdkEthereum(cfg)
    case 10:
        sdkDomain, err = optimism.NewSdkOptimism(cfg)
    default:
        return nil, fmt.Errorf("unsupported chain id %d", chainId)
    }

    if err != nil {
        return nil, fmt.Errorf("Failed to instantiate SDK: %w", err)
    }

    return &NovaSDK{
        Config:   cfg,
        SdkDomain: sdkDomain,
    }, nil
}
