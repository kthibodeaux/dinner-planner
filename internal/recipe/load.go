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

func (ingredientList *IngredientList) UnmarshalTOML(data any) error {
	list, ok := data.([]any)
	if !ok {
		return fmt.Errorf("expected list of ingredients, got %T", data)
	}

	for _, item := range list {
		entry, ok := item.([]any)
		if !ok {
			return fmt.Errorf("expected ingredient entry to be a list, got %T", item)
		}

		ingredient := Ingredient{}
		switch len(entry) {
		case 3:
			ingredient.Unit = fmt.Sprint(entry[2])
			fallthrough
		case 2:
			ingredient.Quantity = fmt.Sprint(entry[1])
			fallthrough
		case 1:
			ingredient.Name = fmt.Sprint(entry[0])
		default:
			return fmt.Errorf("invalid ingredient format: %v", entry)
		}

		*ingredientList = append(*ingredientList, ingredient)
	}
	return nil
}
