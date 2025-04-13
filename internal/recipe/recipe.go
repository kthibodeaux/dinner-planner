package recipe

type Recipe struct {
	CookbookCategory string
	Name             string
	Parts            []Part
	Source           string
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
