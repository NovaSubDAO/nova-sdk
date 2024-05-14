package tests

import (
    "os"
    "path/filepath"
    "testing"
)

func TestMain(m *testing.M) {
    cwd, err := os.Getwd()
    if err != nil {
        panic("Failed to get current working directory")
    }

    configPath := filepath.Join(cwd, "../config.json")
	os.Setenv("CONFIG_FILE", configPath)

	code := m.Run()
    os.Exit(code)
}
