package planner

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/kthibodeaux/dinner-planner/internal/config"
)

func (dp *dinnerPlan) viewModeHelp() string {
	width := 30
	height := 20

	title := styleSelected.Render("Help")
	subtitle := config.Get().Planner.Keys.MainView + " to close"
	keyBindings := styleSelected.Render("Key Bindings")

	header := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(width - 2).
		Render(title + "\n" + subtitle + "\n\n" + keyBindings + "\n")

	content := header +
		"Help:           " + config.Get().Planner.Keys.Help + "\n" +
		"Main:           " + config.Get().Planner.Keys.MainView + "\n" +
		"Shopping List:  " + config.Get().Planner.Keys.ShoppingList + "\n" +
		"> Toggle:       " + config.Get().Planner.Keys.ShoppingListToggle + "\n" +
		"\n" +
		"Up:             " + config.Get().Planner.Keys.Up + "\n" +
		"Down:           " + config.Get().Planner.Keys.Down + "\n" +
		"Scroll Up:      " + config.Get().Planner.Keys.ScrollUp + "\n" +
		"Scroll Down:    " + config.Get().Planner.Keys.ScrollDown + "\n" +
		"\n" +
		"Focus:          " + config.Get().Planner.Keys.Focus + "\n" +
		"Recipes:        " + config.Get().Planner.Keys.Recipes + "\n" +
		"Day 1:          " + config.Get().Planner.Keys.Day1 + "\n" +
		"Day 2:          " + config.Get().Planner.Keys.Day2 + "\n" +
		"Day 3:          " + config.Get().Planner.Keys.Day3 + "\n" +
		"Day 4:          " + config.Get().Planner.Keys.Day4 + "\n" +
		"Day 5:          " + config.Get().Planner.Keys.Day5 + "\n" +
		"Day 6:          " + config.Get().Planner.Keys.Day6 + "\n" +
		"Day 7:          " + config.Get().Planner.Keys.Day7 + "\n" +
		"\n" +
		"Quit:           " + config.Get().Planner.Keys.Quit

	pane := dp.stylePaneBorder(borderForce).
		Width(width).
		Height(height).
		Padding(1, 2).
		Render(content)

	return lipgloss.Place(
		dp.size.width,
		dp.size.height,
		lipgloss.Center,
		lipgloss.Center,
		pane,
	)
}
