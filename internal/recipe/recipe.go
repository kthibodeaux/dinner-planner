package recipe

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
