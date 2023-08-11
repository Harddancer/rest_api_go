package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Harddancer/rest_api_go/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {

	flag.Parse()

	config := apiserver.Newconfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)

	}

	fmt.Println(configPath)
	fmt.Println(config)

}
