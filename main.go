package main

import (
	"context"
	"log"

	"github.com/nilsson-petter/go-rest-test/api"
	"github.com/nilsson-petter/go-rest-test/config"
)

func main() {
	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(cfg)
	}

	server := api.NewServer(cfg.HTTPServer)

	server.Start(ctx)

}
