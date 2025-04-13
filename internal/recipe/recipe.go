package recipe

type Recipe struct {
	CookbookCategory  string
	IncludeInCookbook bool
	Name              string
	Parts             []Part
	Source            string
}

type Part struct {
	CookTime    TimeUnit
	Ingredients []string
	Name        string
	PrepTime    TimeUnit
	Steps       []string
}

type TimeUnit struct {
	Amount int
	Unit   string
}

func NewRecipe(name, cookbookCategory, source string, includeInCookbook bool, parts []Part) *Recipe {
	return &Recipe{
		Name:              name,
		CookbookCategory:  cookbookCategory,
		Source:            source,
		IncludeInCookbook: includeInCookbook,
		Parts:             parts,
	}
}

func (r *Recipe) AddPart(part Part) {
	r.Parts = append(r.Parts, part)
}
