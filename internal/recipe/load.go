package recipe

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/kthibodeaux/dinner-planner/internal/config"
	"github.com/kthibodeaux/dinner-planner/internal/utils"
)

func Load() ([]*Recipe, error) {
	entries, err := os.ReadDir(config.Get().RecipeDirectory)
	if err != nil {
		return nil, err
	}

	var recipes []*Recipe

	for _, entry := range entries {
		if !entry.IsDir() {
			recipe, err := parse(filepath.Join(config.Get().RecipeDirectory, entry.Name()))
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
		return nil, err
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	recipe := Recipe{
		ID: strings.TrimSuffix(filepath.Base(filePath), ".toml"),
	}

	err = toml.Unmarshal(b, &recipe)
	if err != nil {
		return nil, err
	}

	recipe.CategoryID = utils.Slugify(recipe.CookbookCategory)

	return &recipe, nil
}
