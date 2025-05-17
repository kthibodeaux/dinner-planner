package planner

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/kthibodeaux/dinner-planner/internal/recipe"
)

type RecipeList struct {
	Hotkey        string
	Offset        int
	Recipes       []*recipe.Recipe
	SelectedIndex int
	SelectedStyle lipgloss.Style
	Title         string
}

func NewRecipeList(selectedStyle lipgloss.Style, hotkey string, title string, recipes []*recipe.Recipe) RecipeList {
	return RecipeList{
		SelectedStyle: selectedStyle,
		Title:         title,
		Hotkey:        hotkey,
		Recipes:       recipes,
	}
}

func (rl *RecipeList) Render(size Size) string {
	viewCount := min(size.height-2, len(rl.Recipes))

	recipes := rl.Recipes[rl.Offset:viewCount]
	recipeNames := make([]string, 0)
	for r := range recipes {
		name := recipes[r].Name
		if len(recipes[r].Name) > size.width-4 {
			name = recipes[r].Name[:size.width-4] + "..."
		}
		if r == rl.SelectedIndex {
			name = rl.SelectedStyle.Render(name)
		}
		recipeNames = append(recipeNames, name)
	}

	return rl.RenderHeader() + strings.Join(recipeNames, "\n")
}

func (rl *RecipeList) RenderHeader() string {
	keyInfo := ""

	if rl.Hotkey != "" {
		keyInfo = rl.SelectedStyle.Render("[" + rl.Hotkey + "] ")
	}

	return lipgloss.NewStyle().Render(keyInfo + rl.Title + "\n\n")
}
