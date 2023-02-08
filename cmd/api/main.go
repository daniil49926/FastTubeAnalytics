package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/daniil49926/FastTubeAnalytics/internal/app/api"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/api.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := api.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := api.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
