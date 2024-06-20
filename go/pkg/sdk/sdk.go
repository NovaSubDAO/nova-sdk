package sdk

import (
	"fmt"

	"github.com/NovaSubDAO/nova-sdk/go/pkg/config"
	"github.com/NovaSubDAO/nova-sdk/go/pkg/sdk/ethereum"
	"github.com/NovaSubDAO/nova-sdk/go/pkg/sdk/optimism"
)

func NewNovaSDK(rpcEndpoint string, chainId int64) (SdkInterface, error) {
	cfg, err := config.LoadConfig(rpcEndpoint, chainId)
	if err != nil {
		return nil, err
	}

	var novaSdk SdkInterface
	switch chainId {
	case 1:
		novaSdk, err = ethereum.NewSdkEthereum(cfg)
	case 10:
		novaSdk, err = optimism.NewSdkOptimism(cfg)
	default:
		return nil, fmt.Errorf("unsupported chain id %d", chainId)
	}

	if err != nil {
		return nil, fmt.Errorf("Failed to instantiate SDK: %w", err)
	}

	return novaSdk, nil
}
