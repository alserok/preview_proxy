package main

import (
	"github.com/alserok/preview_proxy/server/internal/app"
	"github.com/alserok/preview_proxy/server/internal/config"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app.MustStart(config.MustLoad())
}
