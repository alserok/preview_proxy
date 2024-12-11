package config

type Config struct {
	Port string
	Env  string

	Cache struct {
		Addr string
	}

	API struct {
		YoutubeAddr string
	}
}

func MustLoad() *Config {
	var cfg Config

	return &cfg
}
