package config

import (
	"flag"
)

type AppConfig struct {
	Port             string
	BaseShortenerURL string
}

var Config = &AppConfig{}

func ParseFlags() {
	flag.StringVar(&Config.Port, "a", ":8080", "address and port to run server")
	flag.StringVar(&Config.BaseShortenerURL, "b", "http://localhost:8000/qsd54gFg", "base shortener URL")
	flag.Parse()
}
