package tui

import "github.com/charmbracelet/lipgloss"

var (
	borderForce = -1
	borderSize  = 1
)

func (dp *dinnerPlan) paneBorder(index int) lipgloss.Style {
	if index == dp.focusIndex || index == borderForce {
		return lipgloss.NewStyle().
			Border(lipgloss.ThickBorder()).
			BorderForeground(lipgloss.Color(*dp.color))
	} else {
		return lipgloss.NewStyle().
			Border(lipgloss.ThickBorder())
	}
}

func (dp *dinnerPlan) styleSelected(bold bool) lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(*dp.color)).
		Bold(bold)
}

func (dp *dinnerPlan) paneHeader(key string, title string) string {
	keyInfo := ""

	if key != "" {
		keyInfo = dp.styleSelected(false).Render("[" + key + "] ")
	}

	return lipgloss.NewStyle().Render(keyInfo + title)
}
