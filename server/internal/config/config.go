package config

import "os"

type Config struct {
	Port string
	Env  string

	Cache struct {
		Addr string
	}
}

func MustLoad() *Config {
	var cfg Config

	cfg.Port = os.Getenv("PORT")
	cfg.Env = os.Getenv("ENV")
	cfg.Cache.Addr = os.Getenv("REDIS_ADDR")

	return &cfg
}
