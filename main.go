package main

import (
	"github.com/IgorCastilhos/go-ecommerce-app/config"
	"github.com/IgorCastilhos/go-ecommerce-app/internal/api"
	"log"
)

func main() {
	cfg, err := config.SetupEnv()

	if err != nil {
		log.Fatalf("arquivo de configuração não foi carregado: %v\n", err)
	}

	api.StartServer(cfg)
}
