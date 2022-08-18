package store

import (
	"fmt"
	"os"
)

type Config struct {
	Host string
	Port string
}

func NewConfig() *Config {
	config := &Config{}

	//flag.StringVar(&config.Host, "host", "localhost", "host for db")
	//flag.StringVar(&config.Port, "port", ":6379", "port for db")
	//flag.Parse()
	config.Host = os.Getenv("host")
	config.Port = os.Getenv("port")

	fmt.Printf("Host: %s, port: %s \n", config.Host, config.Port)

	return config
}
