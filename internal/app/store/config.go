package store

import "flag"

type Config struct {
	Host string
	Port string
}

var (
	host string
	port string
)

func NewConfig() *Config {
	flag.StringVar(&host, "host", "localhost", "host for Redis DB")
	flag.StringVar(&port, "port", ":6379", "port for Redis DB")

	return &Config{
		Host: host,
		Port: port,
	}
}
