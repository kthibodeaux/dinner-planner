package recipe

import "fmt"

type Recipe struct {
	CookbookCategory string       `toml:"category"`
	Dependencies     []Dependency `toml:"dependencies"`
	ID               string
	Name             string `toml:"name"`
	Notes            string `toml:"notes"`
	Parts            []Part `toml:"parts"`
	Source           string `toml:"source"`
}

type Part struct {
	CookTime    int            `toml:"cook_time"`
	Ingredients IngredientList `toml:"ingredients"`
	Name        string         `toml:"name"`
	Notes       string         `toml:"notes"`
	PrepTime    int            `toml:"prep_time"`
	Steps       []string       `toml:"steps"`
}

type IngredientList []Ingredient

type Ingredient struct {
	Name     string `toml:"name"`
	Quantity string `toml:"quantity"`
	Unit     string `toml:"unit"`
}

type Dependency struct {
	Notes    string `toml:"notes"`
	RecipeID string `toml:"recipe_id"`
	Replaces string `toml:"replaces"`
	Required bool   `toml:"required"`
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
