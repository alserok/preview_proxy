package config

type Config struct {
	Port string
}

func MustLoad() *Config {
	var cfg Config

	return &cfg
}
