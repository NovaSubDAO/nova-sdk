package config

import (
	"fmt"

	"github.com/NovaSubDAO/nova-sdk/go/pkg/constants"
)

type Config struct {
	VaultAddress  string `json:"vaultAddress"`
	VaultDecimals int64  `json:"vaultDecimals"`
	SDai          string
	VelodromePool string
	RpcEndpoint   string
	ChainId       int64
}

func convertConfigDataToConfig(data constants.ConfigData) Config {
	return Config{
		VaultAddress:  data.VaultAddress,
		VaultDecimals: data.VaultDecimals,
		SDai:          data.SDai,
		VelodromePool: data.VelodromePool,
	}
}

func LoadConfig(rpcEndpoint string, chainId int64) (*Config, error) {
	configData, exists := constants.ConfigDetails[chainId]
	if !exists {
		return nil, fmt.Errorf("unsupported chain id %d", chainId)
	}

	config := convertConfigDataToConfig(configData)
	config.RpcEndpoint = rpcEndpoint
	config.ChainId = chainId

	return &config, nil
}
