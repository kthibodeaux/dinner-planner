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
		if len(entry) == 3 {
			ingredient.Quantity = fmt.Sprint(entry[0])
			ingredient.Unit = fmt.Sprint(entry[1])
			ingredient.Name = fmt.Sprint(entry[2])
		} else if len(entry) == 2 {
			ingredient.Quantity = fmt.Sprint(entry[0])
			ingredient.Name = fmt.Sprint(entry[1])
		} else if len(entry) == 1 {
			ingredient.Name = fmt.Sprint(entry[0])
		}

		*ingredientList = append(*ingredientList, ingredient)
	}
	return nil
}
