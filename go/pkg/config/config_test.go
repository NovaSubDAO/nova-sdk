package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Define a sample config to be written to the file
	sampleConfig := Config{
		SDaiAddress:  "0x123",
		SDaiDecimals: 18,
		RpcEndpoint:  "http://localhost:8545",
		PrivateKey:   "your-private-key",
		ChainId:      1,
	}

    // Override with env variables
    overrideWithEnv(&sampleConfig)

	// Marshal the sample config into JSON
	configData, err := json.Marshal(sampleConfig)
	if err != nil {
		t.Fatalf("Unable to marshal sample config: %v", err)
	}

	// Create a temporary file
	tmpFile, err := ioutil.TempFile("", "test_config_tmp.json")
	if err != nil {
		t.Fatalf("Unable to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up after the test

	// Write the JSON data to the temp file
	if _, err := tmpFile.Write(configData); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	if err := tmpFile.Close(); err != nil {
		t.Fatalf("Failed to close file: %v", err)
	}

	// Set environment variable to point to the temp file
	os.Setenv("CONFIG_FILE", tmpFile.Name())
	defer os.Unsetenv("CONFIG_FILE") // Clean up environment variable after test

	// Call the function under test
	config, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig returned an error: %v", err)
	}

	// Check the loaded config against the expected values
	if config.SDaiAddress != sampleConfig.SDaiAddress ||
		config.SDaiDecimals != sampleConfig.SDaiDecimals ||
		config.RpcEndpoint != sampleConfig.RpcEndpoint ||
		config.PrivateKey != sampleConfig.PrivateKey ||
		config.ChainId != sampleConfig.ChainId {
		t.Errorf("Config did not match expected values: got %+v, want %+v", config, sampleConfig)
	}
}
