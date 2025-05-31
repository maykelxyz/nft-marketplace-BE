package main

import (
	"log"
	"nft-marketplace-be/src/cmd"
	"nft-marketplace-be/src/config"
)

func main() {
	cfg, err := config.InitConfig(".env")
	if err != nil {
		log.Fatalf("Failed to initialize config: %v", err)
	}
	cmd.InitializeServer(cfg)
}
