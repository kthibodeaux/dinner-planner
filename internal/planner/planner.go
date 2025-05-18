package planner

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/kthibodeaux/dinner-planner/internal/config"
	"github.com/kthibodeaux/dinner-planner/internal/recipe"
	shoppingList "github.com/kthibodeaux/dinner-planner/internal/shopping_list"
)

type Mode int

const (
	ModeAssign Mode = iota
	ModeHelp
	ModeNavigatePane
	ModeShoppingList
)

type Size struct {
	width  int
	height int
}

type dinnerPlan struct {
	recipeLists  []*RecipeList
	shoppingList *shoppingList.ShoppingList

	paneFocusIndex int
	mode           Mode
	size           Size
}

func (dp dinnerPlan) Init() tea.Cmd {
	return nil
}

func (dp dinnerPlan) View() string {
	if dp.size.width < 50 {
		return "window is too narrow"
	}
	if dp.size.height < 24 {
		return "window is too short"
	}

	if dp.mode == ModeHelp {
		return dp.viewModeHelp()
	} else if dp.mode == ModeShoppingList {
		return dp.viewModeShoppingList()
	} else {
		return dp.viewModeAssign()
	}
}

func Run(recipes []*recipe.Recipe, dates []time.Time) {
	dinnerPlan := dinnerPlan{
		recipeLists: make([]*RecipeList, 8),
		mode:        ModeAssign,
	}

	dayKeyMap := map[int]string{
		0: config.Get().Planner.Keys.Day1,
		1: config.Get().Planner.Keys.Day2,
		2: config.Get().Planner.Keys.Day3,
		3: config.Get().Planner.Keys.Day4,
		4: config.Get().Planner.Keys.Day5,
		5: config.Get().Planner.Keys.Day6,
		6: config.Get().Planner.Keys.Day7,
	}

	mainRecipeList := NewRecipeList(config.Get().Planner.Keys.Recipes, "Recipes", recipes)
	dinnerPlan.recipeLists[0] = &mainRecipeList
	for index := range dates {
		hotkey := dayKeyMap[index]
		title := dates[index].Format("Monday, January 2")
		recipeList := NewRecipeList(hotkey, title, make([]*recipe.Recipe, 0))
		dinnerPlan.recipeLists[index+1] = &recipeList
	}

	p := tea.NewProgram(
		dinnerPlan,
		tea.WithAltScreen(),
	)
	p.Run()
}
