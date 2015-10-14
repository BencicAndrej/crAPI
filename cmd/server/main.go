package main

import (
	"fmt"
	"github.com/BencicAndrej/crAPI"
	"github.com/BencicAndrej/crAPI/config"
	"log"
	"net/http"
)

func main() {
	router := router.New()

	configuration, err := config.Load("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	routes, err := configuration.ParseRoutes()
	if err != nil {
		log.Fatal(err)
	}

	router.RegisterBatch(routes)

	fmt.Println("Starting up server...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
