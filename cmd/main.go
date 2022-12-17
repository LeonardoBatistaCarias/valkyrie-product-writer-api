package main

import (
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/config"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/utils/logger"
	"log"

	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/server"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	logger := logger.NewAppLogger(cfg.Logger)
	logger.InitLogger()

	s := server.NewServer(logger, cfg)
	logger.Fatal(s.Run())
}
