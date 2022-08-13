package store

type Config struct {
	Host string `toml:"host_db"`
	Port string `toml:"port_db"`
}

func NewConfig() *Config {
	return &Config{
		Host: "localhost",
		Port: ":6379",
	}
}
