package config

type Config struct {
	Port string

	API struct {
		YoutubeAddr string
	}
}

func MustLoad() *Config {
	var cfg Config

	return &cfg
}
