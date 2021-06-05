package main

import (
	"flag"
	"log"
	"mishaga/internal/app"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "configs/example.json", "Path to config file")
}

func main() {
	log.Fatal(app.Run(configPath))
}
