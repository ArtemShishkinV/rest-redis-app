package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"rest-redis-app/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ConfigPath = %s, Addr = %s \n", configPath, config.Store.Host+config.Store.Port)
	fmt.Println(config.BindAddr)

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
