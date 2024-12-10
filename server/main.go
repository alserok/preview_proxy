package main

import (
	"github.com/alserok/preview_proxy/server/internal/app"
	"github.com/alserok/preview_proxy/server/internal/config"
)

func main() {
	app.MustStart(config.MustLoad())
}
