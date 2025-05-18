package planner

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/kthibodeaux/dinner-planner/internal/config"
)

var (
	borderForce = -1
	borderSize  = 1

	styleSelected = lipgloss.NewStyle().
			Foreground(lipgloss.Color(config.Get().Planner.Color)).
			Bold(true)
)

func (dp *dinnerPlan) stylePaneBorder(index int) lipgloss.Style {
	if index == dp.paneFocusIndex || index == borderForce {
		return lipgloss.NewStyle().
			Border(lipgloss.ThickBorder()).
			BorderForeground(lipgloss.Color(config.Get().Planner.Color))
	} else {
		return lipgloss.NewStyle().
			Border(lipgloss.ThickBorder())
	}
}
