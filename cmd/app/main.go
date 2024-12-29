package main

import (
	"log"

	"github.com/henriquepw/pobrin-api/api"
	"github.com/henriquepw/pobrin-api/pkg/config"
)

func main() {
	config.Env()

	jobServer, err := api.NewJobServer()
	if err != nil {
		log.Fatal(err.Error())
	}
	if err = jobServer.Start(); err != nil {
		log.Fatal(err.Error())
	}

	apiServer := api.NewApiServer()
	if err = apiServer.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
