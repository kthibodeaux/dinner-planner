package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/kthibodeaux/dinner-planner/internal/config"
	"github.com/kthibodeaux/dinner-planner/internal/recipe"
	"github.com/kthibodeaux/dinner-planner/internal/utils"
)

type dinnerPlan struct {
	color   *string
	keys    *config.KeyConfig
	recipes []*recipe.Recipe
	dates   []time.Time

	focusIndex int
	size       Size
}

type Size struct {
	width  int
	height int
}

func (dp dinnerPlan) Init() tea.Cmd {
	return nil
}

func main() {
	config := config.LoadConfig()

	p := tea.NewProgram(
		dinnerPlan{
			color:   &config.Planner.Color,
			keys:    &config.Planner.Keys,
			recipes: recipe.Load(config.RecipeDirectory),
			dates:   utils.DatesForWeekStartingOn(config.StartDate),
		},
	)
	p.Run()
}
