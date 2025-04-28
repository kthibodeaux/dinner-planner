package planner

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/kthibodeaux/dinner-planner/internal/config"
	"github.com/kthibodeaux/dinner-planner/internal/recipe"
)

type Mode int

const (
	ModeAssign Mode = iota
	ModeHelp
	ModeNavigatePane
)

type Size struct {
	width  int
	height int
}

type dinnerPlan struct {
	color     *string
	dayKeyMap map[int]string
	keys      *config.KeyConfig
	recipes   []*recipe.Recipe
	dates     []time.Time

	focusIndex int
	mode       Mode
	size       Size
}

func (dp dinnerPlan) Init() tea.Cmd {
	return nil
}

func (dp dinnerPlan) View() string {
	if dp.mode == ModeHelp {
		return dp.viewModeHelp()
	} else {
		return dp.viewModeAssign()
	}
}

func Run(config *config.Config, recipes []*recipe.Recipe, dates []time.Time) {
	p := tea.NewProgram(
		dinnerPlan{
			color:     &config.Planner.Color,
			dayKeyMap: config.DayKeyMap(),
			keys:      &config.Planner.Keys,
			recipes:   recipes,
			dates:     dates,
			mode:      ModeAssign,
		},
	)
	p.Run()
}
