package config

import (
	"fmt"
	"github.com/BencicAndrej/crAPI"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

type RouteConfig struct {
	Method             string `yaml:"method"`
	Path               string `yaml:"path"`
	String             string `yaml:"string"`
	File               string `yaml:"file"`
	MultipartVariables bool   `yaml:"multipartVariables"`
	Command            string `yaml: "command"`
}

type Config struct {
	Defaults struct {
		Port string `yaml:"port"`
	} `yaml:"defaults"`
	Routes []RouteConfig `yaml:"routes"`
}

func Load(configPath string) (config Config, err error) {
	fileContents, err := ioutil.ReadFile(configPath)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(fileContents, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func (config Config) ParseRoutes() (routes []*router.Route, err error) {
	for _, route := range config.Routes {
		handler, err := resolveHandler(route)
		if err != nil {
			return nil, err
		}

		routes = append(routes, &router.Route{
			Method:             route.Method,
			Path:               route.Path,
			MultiPartVariables: route.MultipartVariables,
			Handler:            handler,
		})
	}

	return routes, nil
}

func resolveHandler(rConf RouteConfig) (handler http.HandlerFunc, err error) {
	if len(rConf.Command) > 0 {
		return newCommandHandler(rConf.Command), nil
	}
	if len(rConf.File) > 0 {
		return newFileHandler(rConf.File)
	}
	if len(rConf.String) > 0 {
		return newStringHandler(rConf.String), nil
	}

	err = fmt.Errorf("No handlers found for route %s %s.", rConf.Method, rConf.Path)

	return nil, err
}

func newCommandHandler(commandString string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		//@TODO: See how supervisor handles command input. #srama#bez

		commandParts := strings.Split(commandString, " ")

		command := exec.Command(commandParts[0], commandParts[1:]...)

		output, err := command.Output()
		if err != nil {
			log.Println(err)
		}

		fmt.Fprint(rw, string(output))
	}
}

func newFileHandler(path string) (handler http.HandlerFunc, err error) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return newStringHandler(string(body)), err
}

func newStringHandler(body string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, body)
	}
}
