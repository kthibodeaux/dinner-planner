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

func NewRecipe(name, cookbookCategory, source string, parts []Part) *Recipe {
	return &Recipe{
		Name:             name,
		CookbookCategory: cookbookCategory,
		Source:           source,
		Parts:            parts,
	}
}

func (r *Recipe) AddPart(part Part) {
	r.Parts = append(r.Parts, part)
}
