package planner

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (dp *dinnerPlan) viewModeShoppingList() string {
	title := styleSelected.Render("Shopping List")
	header := lipgloss.NewStyle().Render(title + "\n")

	content := make([]string, 0)
	content = append(content, header)

	return dp.stylePaneBorder(borderForce).
		Width(dp.size.width).
		Height(dp.size.height - 2).
		Render(strings.Join(content, "\n"))
}
