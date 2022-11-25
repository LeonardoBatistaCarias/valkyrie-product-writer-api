package main

import (
	"log"

	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/config"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/server"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(cfg)
	log.Fatal(s.Run())
}
