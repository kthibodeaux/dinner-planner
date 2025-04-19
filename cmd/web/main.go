package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/kthibodeaux/dinner-planner/internal/recipe"
)

type config struct {
	directory string
}

func main() {
	var config config
	flag.StringVar(&config.directory, "directory", defaultPath(), "Directory containing recipes")
	flag.Parse()

	recipes := recipe.Load(config.directory)
	recipes[3].ExportToWebsiteHTML(recipes[3].ID + ".html")
}

func defaultPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(home, "recipes")
}
