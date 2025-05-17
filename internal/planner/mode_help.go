package planner

import "github.com/charmbracelet/lipgloss"

func (dp *dinnerPlan) viewModeHelp() string {
	width := 30
	height := 20

	title := dp.styleSelected().Render("Help")
	subtitle := dp.keys.MainView + " to close"
	keyBindings := dp.styleSelected().Render("Key Bindings")

	header := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(width - 2).
		Render(title + "\n" + subtitle + "\n\n" + keyBindings + "\n")

	content := header +
		"Focus:        " + dp.keys.Focus + "\n" +
		"Help:         " + dp.keys.Help + "\n" +
		"Recipes:      " + dp.keys.Recipes + "\n" +
		"Day 1:        " + dp.keys.Day1 + "\n" +
		"Day 2:        " + dp.keys.Day2 + "\n" +
		"Day 3:        " + dp.keys.Day3 + "\n" +
		"Day 4:        " + dp.keys.Day4 + "\n" +
		"Day 5:        " + dp.keys.Day5 + "\n" +
		"Day 6:        " + dp.keys.Day6 + "\n" +
		"Day 7:        " + dp.keys.Day7 + "\n" +
		"Quit:         " + dp.keys.Quit

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
