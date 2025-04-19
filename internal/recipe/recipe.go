package recipe

import (
	"fmt"

	"github.com/kthibodeaux/dinner-planner/internal/utils"
)

type Recipe struct {
	CookbookCategory string `toml:"category"`
	ID               string
	Name             string   `toml:"name"`
	Notes            []string `toml:"notes"`
	Parts            []Part   `toml:"parts"`
	Source           string   `toml:"source"`
}

type Part struct {
	CookTime    int          `toml:"cook_time"`
	Ingredients []Ingredient `toml:"ingredients"`
	Name        string       `toml:"name"`
	Notes       []string     `toml:"notes"`
	PrepTime    int          `toml:"prep_time"`
	Steps       []string     `toml:"steps"`
}

type IngredientList []Ingredient

type Ingredient struct {
	Name     string `toml:"name"`
	Quantity string `toml:"quantity"`
	Unit     string `toml:"unit"`
	Note     string `toml:"note"`
	RecipeID string `toml:"recipe_id"`
}

func (recipe *Recipe) GetCategorySlug() string {
	return utils.Slugify(recipe.CookbookCategory)
}

func (part *Part) GetName() string {
	if part.Name == "" {
		return "Ingredients"
	}

	return part.Name
}

func (part *Part) GetPrepTime() string {
	if part.PrepTime == 0 {
		return ""
	}

	return fmt.Sprintf("%d minutes", part.PrepTime)
}

func (part *Part) GetCookTime() string {
	if part.CookTime == 0 {
		return ""
	}

	return fmt.Sprintf("%d minutes", part.CookTime)
}
