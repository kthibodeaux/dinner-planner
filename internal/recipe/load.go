package recipe

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

func Load(directory string) []*Recipe {
	entries, err := os.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	var recipes []*Recipe

	for _, entry := range entries {
		if !entry.IsDir() {
			fmt.Println("Loading recipe:", entry.Name())
			recipe, err := parse(filepath.Join(directory, entry.Name()))
			if err != nil {
				log.Printf("Error loading recipe %s: %v", entry.Name(), err)
			}
			recipes = append(recipes, recipe)
		}
	}

	return recipes
}

func parse(filePath string) (*Recipe, error) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	recipe := Recipe{
		ID: strings.TrimSuffix(filepath.Base(filePath), ".toml"),
	}

	err = toml.Unmarshal(b, &recipe)
	if err != nil {
		panic(err)
	}

	return &recipe, nil
}
