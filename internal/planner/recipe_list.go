package planner

import (
	"slices"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/kthibodeaux/dinner-planner/internal/recipe"
)

type RecipeList struct {
	Hotkey        string
	Offset        int
	Recipes       []*recipe.Recipe
	SelectedIndex int
	Title         string
}

func NewRecipeList(hotkey string, title string, recipes []*recipe.Recipe) RecipeList {
	return RecipeList{
		Title:   title,
		Hotkey:  hotkey,
		Recipes: recipes,
	}
}

func (rl *RecipeList) Render(isActive bool, size Size) string {
	viewCount := min(size.height-2, len(rl.Recipes))

	recipes := rl.Recipes[rl.Offset:viewCount]
	recipeNames := make([]string, 0)
	for r := range recipes {
		name := recipes[r].Name
		if len(recipes[r].Name) > size.width-4 {
			name = recipes[r].Name[:size.width-4] + "..."
		}
		if r == rl.SelectedIndex && isActive {
			name = styleSelected.Render(name)
		}
		recipeNames = append(recipeNames, name)
	}

	return rl.RenderHeader() + strings.Join(recipeNames, "\n")
}

func (rl *RecipeList) RenderHeader() string {
	keyInfo := ""

	if rl.Hotkey != "" {
		keyInfo = styleSelected.Render("[" + rl.Hotkey + "] ")
	}

	return lipgloss.NewStyle().Render(keyInfo + rl.Title + "\n\n")
}

func (rl *RecipeList) handleDown() {
	if rl.SelectedIndex >= len(rl.Recipes)-1 {
		return
	}

	rl.SelectedIndex++
}

func (rl *RecipeList) handleUp() {
	if rl.SelectedIndex == 0 {
		return
	}

	rl.SelectedIndex--
}

func (rl *RecipeList) remove() {
	rl.Recipes = slices.Delete(rl.Recipes, rl.SelectedIndex, rl.SelectedIndex+1)
}

func (rl *RecipeList) add(recipe *recipe.Recipe) {
	rl.Recipes = append(rl.Recipes, recipe)
}

func (rl *RecipeList) checkSelectedIndex() {
	if rl.SelectedIndex > len(rl.Recipes)-1 {
		rl.SelectedIndex = max(0, len(rl.Recipes)-1)
	}
}

func (rl *RecipeList) selectedRecipe() *recipe.Recipe {
	return rl.Recipes[rl.SelectedIndex]
}
