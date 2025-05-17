package web

import (
	"github.com/kthibodeaux/dinner-planner/internal/recipe"
	"github.com/kthibodeaux/dinner-planner/internal/utils"
)

func getRecipes() (map[string]*recipe.Recipe, error) {
	allRecipes, err := recipe.Load()
	if err != nil {
		return nil, err
	}

	recipes := make(map[string]*recipe.Recipe)
	for _, recipe := range allRecipes {
		if _, exists := recipes[recipe.ID]; !exists {
			recipes[recipe.ID] = recipe
		}
	}

	return recipes, nil
}

func getUniqueCategories() map[string]string {
	seen := make(map[string]string)

	for _, recipe := range recipes {
		if recipe.CookbookCategory == "" {
			continue
		}
		slug := utils.Slugify(recipe.CookbookCategory)
		if _, exists := seen[slug]; !exists {
			seen[slug] = recipe.CookbookCategory
		}
	}

	return seen
}
