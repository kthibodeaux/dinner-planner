package planner

import "github.com/charmbracelet/lipgloss"

var (
	borderForce = -1
	borderSize  = 1
)

func (dp *dinnerPlan) stylePaneBorder(index int) lipgloss.Style {
	if index == dp.paneFocusIndex || index == borderForce {
		return lipgloss.NewStyle().
			Border(lipgloss.ThickBorder()).
			BorderForeground(lipgloss.Color(*dp.color))
	} else {
		return lipgloss.NewStyle().
			Border(lipgloss.ThickBorder())
	}
}

func (dp *dinnerPlan) styleSelected() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(*dp.color)).
		Bold(true)
}

func (dp *dinnerPlan) styleListItem() lipgloss.Style {
	return lipgloss.NewStyle().
		PaddingLeft(4)
}

func (dp *dinnerPlan) styleListItemSelected() lipgloss.Style {
	return dp.styleListItem().
		Foreground(lipgloss.Color(*dp.color))
}
