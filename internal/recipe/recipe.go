package recipe

type Recipe struct {
	CookbookCategory string `toml:"category"`
	ID               string
	Name             string `toml:"name"`
	Parts            []Part `toml:"groups"`
	Source           string `toml:"source"`
	DependencyIDs    []string
}

type Part struct {
	CookTime    int      `toml:"cook_time"`
	Ingredients []string `toml:"ingredients"`
	Name        string   `toml:"name"`
	PrepTime    int      `toml:"prep_time"`
	Steps       []string `toml:"steps"`
}
