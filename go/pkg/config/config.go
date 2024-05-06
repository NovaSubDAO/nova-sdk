package config

import (
    "encoding/json"
    "os"
	"strconv"
)

type Config struct {
    SDaiAddress  string `json:"sDaiAddress"`
    RpcEndpoint  string
    PrivateKey   string
    ChainId      int64
}

func LoadConfig() (*Config, error) {
    config := &Config{}
    configFile, err := os.Open("config.json")
    if err != nil {
        return nil, err
    }
    defer configFile.Close()

    decoder := json.NewDecoder(configFile)
    if err = decoder.Decode(config); err != nil {
        return nil, err
    }

    overrideWithEnv(config)
    return config, nil
}

func overrideWithEnv(c *Config) {
    if value := os.Getenv("RPC_ENDPOINT"); value != "" {
        c.RpcEndpoint = value
    }
    if value := os.Getenv("PRIVATE_KEY"); value != "" {
        c.PrivateKey = value
    }
    if value := os.Getenv("CHAIN_ID"); value != "" {
        if chainID, err := strconv.ParseInt(value, 10, 64); err == nil {
            c.ChainId = chainID
        }
    }
}