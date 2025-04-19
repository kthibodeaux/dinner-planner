package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/kthibodeaux/dinner-planner/internal/web"
)

type config struct {
	directory string
	port      string
}

func main() {
	var config config
	flag.StringVar(&config.directory, "directory", defaultPath(), "Directory containing recipes")
	flag.StringVar(&config.port, "port", ":8080", "Port to run the web server on")
	flag.Parse()

	web.Serve(config.directory, config.port)
}

func defaultPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(home, "recipes")
}
