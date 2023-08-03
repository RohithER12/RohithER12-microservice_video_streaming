package main

import (
	"log"

	"github.com/RohithER12/api_gateway/pkg/config"
	"github.com/RohithER12/api_gateway/pkg/di"
)

func main() {
	c, configerr := config.LoadConfig()
	if configerr != nil {
		log.Fatal("cannot load config:", configerr)
	}

	server, dierr := di.InitializeAPI(c)
	if dierr != nil {
		log.Fatal("cannot initialize server", dierr)
	}
	server.Start()
}
