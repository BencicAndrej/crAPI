package main

import (
	"github.com/BencicAndrej/crAPI"
	"github.com/BencicAndrej/crAPI/config"
	"log"
	"net/http"
	"flag"
)

var configPath string

func init() {
	const (
		defaultPath = "config.yml"
		usage = "Path to the configuration file."
	)
	flag.StringVar(&configPath, "config", defaultPath, usage)
	flag.StringVar(&configPath, "c", defaultPath, usage)
}

func main() {
	log.Print("Loading flags...")
	flag.Parse()

	router := router.New()

	log.Printf("Loading cofiguration file: %v...", configPath)
	configuration, err := config.Load(configPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Parsing routes from configuration...")
	routes, err := configuration.ParseRoutes()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Registering routes with the router...")
	router.RegisterBatch(routes)

	port := configuration.Defaults.Port
	if len(port) == 0 {
		port = "8080"
	}

	log.Printf("Starting up server on port: %v...", port)
	log.Fatal(http.ListenAndServe(":" + port, router))
}
