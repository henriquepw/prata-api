package main

import (
	"log"

	"github.com/henriquepw/pobrin-api/api"
)

func main() {
	db := api.StartDB()
	defer db.Close()

	jobServer := api.NewJobServer()
	if err := jobServer.Start(); err != nil {
		log.Fatal(err.Error())
	}

	apiServer := api.NewApiServer(db)
	if err := apiServer.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
