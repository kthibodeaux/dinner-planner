package planner

import (
	"slices"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/kthibodeaux/dinner-planner/internal/config"
	"github.com/kthibodeaux/dinner-planner/internal/recipe"
)

type RecipeList struct {
	Hotkey        string
	Offset        int
	Recipes       []*recipe.Recipe
	SelectedIndex int
	Title         string
	ViewCount     int
}

func NewRecipeList(hotkey string, title string, recipes []*recipe.Recipe) RecipeList {
	return RecipeList{
		Title:   title,
		Hotkey:  hotkey,
		Recipes: recipes,
	}
}

func (rl *RecipeList) Render(isActive bool, size Size) string {
	rl.ViewCount = min(size.height-2, len(rl.Recipes))

	end := min(rl.Offset+rl.ViewCount, len(rl.Recipes))
	recipes := rl.Recipes[rl.Offset:end]
	recipeNames := make([]string, 0)

	for i, r := range recipes {
		name := r.Name
		if len(name) > size.width-4 {
			name = name[:size.width-4] + "..."
		}
		actualIndex := rl.Offset + i
		if actualIndex == rl.SelectedIndex && isActive {
			name = styleSelected.Render(name)
		}
		recipeNames = append(recipeNames, name)
	}

	return rl.renderHeader() + strings.Join(recipeNames, "\n")
}

func (rl *RecipeList) renderHeader() string {
	keyInfo := ""

	if rl.Hotkey != "" {
		keyInfo = styleSelected.Render("[" + rl.Hotkey + "] ")
	}

	return lipgloss.NewStyle().Render(keyInfo + rl.Title + "\n\n")
}

func (rl *RecipeList) handleDown(isScroll bool) {
	if rl.SelectedIndex >= len(rl.Recipes)-1 {
		return
	}

	if isScroll {
		if rl.SelectedIndex+rl.scrollAmount() >= len(rl.Recipes) {
			rl.SelectedIndex = len(rl.Recipes) - 1
		} else {
			rl.SelectedIndex += rl.scrollAmount()
		}
	} else {
		rl.SelectedIndex++
	}
	rl.setVisible()
}

func (rl *RecipeList) handleUp(isScroll bool) {
	if rl.SelectedIndex <= 0 {
		return
	}

	if isScroll {
		if rl.SelectedIndex-rl.scrollAmount() < 0 {
			rl.SelectedIndex = 0
		} else {
			rl.SelectedIndex -= rl.scrollAmount()
		}
	} else {
		rl.SelectedIndex--
	}
	rl.setVisible()
}

func (rl *RecipeList) scrollAmount() int {
	if rl.ViewCount < config.Get().Planner.ScrollAmount {
		return rl.ViewCount - 2
	} else {
		return config.Get().Planner.ScrollAmount
	}
}

func (rl *RecipeList) setVisible() {
	if rl.SelectedIndex < rl.Offset+config.Get().Planner.ScrollOffset {
		rl.Offset = max(0, rl.SelectedIndex-config.Get().Planner.ScrollOffset)
	}

	if rl.SelectedIndex >= rl.Offset+rl.ViewCount-config.Get().Planner.ScrollOffset {
		rl.Offset = max(min(rl.SelectedIndex-(rl.ViewCount-config.Get().Planner.ScrollOffset-1), len(rl.Recipes)-rl.ViewCount), 0)
	}
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
