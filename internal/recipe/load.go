package recipe

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/kthibodeaux/dinner-planner/internal/utils"
)

func Load(directory string) ([]*Recipe, error) {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	var recipes []*Recipe

	for _, entry := range entries {
		if !entry.IsDir() {
			recipe, err := parse(filepath.Join(directory, entry.Name()))
			if err != nil {
				log.Printf("Error loading recipe %s: %v", entry.Name(), err)
			}
			recipes = append(recipes, recipe)
		}
	}

	return recipes, nil
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

	recipe.CategoryID = utils.Slugify(recipe.CookbookCategory)

	return &recipe, nil
}
