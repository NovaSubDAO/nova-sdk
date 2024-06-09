package config

import (
	"fmt"
)

type Config struct {
	VaultAddress  string `json:"vaultAddress"`
	VaultDecimals int64  `json:"vaultDecimals"`
	SDai          string
	RpcEndpoint   string
	ChainId       int64
}

func LoadConfig(rpcEndpoint string, chainId int64) (*Config, error) {
	var config Config
	switch chainId {
	case 1:
		config = Config{
			RpcEndpoint:   rpcEndpoint,
			ChainId:       chainId,
			VaultAddress:  "0x83F20F44975D03b1b09e64809B757c47f942BEeA",
			SDai:          "0x83F20F44975D03b1b09e64809B757c47f942BEeA",
			VaultDecimals: 18,
		}
	case 10:
		config = Config{
			RpcEndpoint:   rpcEndpoint,
			ChainId:       chainId,
			SDai:          "0x2218a117083f5B482B0bB821d27056Ba9c04b1D3",
			VaultAddress:  "0x36A2f7Fb07c102415afe2461a9A43377970E081c",
			VaultDecimals: 18,
		}
	default:
		return nil, fmt.Errorf("unsupported chain id %d", chainId)
	}
	return &config, nil
}
