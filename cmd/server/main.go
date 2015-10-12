package main

import (
	"fmt"
	"github.com/BencicAndrej/crAPI"
	"net/http"
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type T struct {
	Defaults struct{
		Prefix string
	  }
	Routes []struct{
		Method string `yaml:"method"`
		Path string `yaml:"path"`
		Output string `yaml: "output"`
	}
}

func LoadRoutes(router *router.Router) error {
	fileContents, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return err
	}

	t := T{}
	err = yaml.Unmarshal(fileContents, &t)
	if err != nil {
		return err
	}

	for _, r := range(t.Routes) {
		out := r.Output
		router.Register(r.Method, r.Path, func(rw http.ResponseWriter, req *http.Request) {
			rw.Write([]byte(out))
		})
	}

	return nil
}

func main() {
	router := router.New()

	LoadRoutes(router)

	fmt.Println("Starting up server...")
	log.Fatal(http.ListenAndServe(":3000", router))
}