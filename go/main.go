package main

import (
	"log"

	"github.com/NovaSubDAO/nova-sdk/go/pkg/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}
	log.Println("sDAI address:", cfg.SDaiAddress)
}
