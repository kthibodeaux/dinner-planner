package main

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	borderForce = -1
	borderSize  = 1
)

func (dp dinnerPlan) View() string {
	if dp.mode == ModeHelp {
		return dp.viewHelp()
	} else {
		unit := dp.size.width / 10
		gap := dp.size.width - (unit * 10)

		recipeColumnWidth := (unit * 6) - (borderSize * 2)
		daysColumnsWidth := (unit * 4) - (borderSize * 2)
		columnHeight := dp.size.height - (borderSize * 2)
		return lipgloss.JoinHorizontal(
			lipgloss.Top,
			dp.recipeColumn(Size{recipeColumnWidth, columnHeight}, gap),
			dp.dayColumns(Size{daysColumnsWidth, columnHeight}),
		)
	}
}

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

func (dp *dinnerPlan) styleSelected() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(*dp.color)).
		Bold(true)
}

func (dp *dinnerPlan) recipeColumn(size Size, gap int) string {
	return dp.paneBorder(0).
		Width(size.width).
		Height(size.height).
		MarginRight(gap).
		Render("# Recipes")
}

func (dp *dinnerPlan) dayColumns(size Size) string {
	dayColumnWidth := size.width/2 - (borderSize * 2)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		dp.daysLeftColumn(Size{dayColumnWidth, size.height}),
		dp.daysRightColumn(Size{dayColumnWidth, size.height}),
	)
}

func (dp *dinnerPlan) daysLeftColumn(size Size) string {
	days := []string{}

	for _, dayNum := range []int{0, 2, 4, 6} {
		days = append(days, dp.dayPane(size, dayNum))
	}

	return lipgloss.JoinVertical(lipgloss.Left, days...)
}

func (dp *dinnerPlan) daysRightColumn(size Size) string {
	days := []string{}

	for _, dayNum := range []int{1, 3, 5} {
		days = append(days, dp.dayPane(size, dayNum))
	}

	return lipgloss.JoinVertical(lipgloss.Left, days...)
}

func (dp *dinnerPlan) dayPane(size Size, index int) string {
	dayPaneHeight := (dp.size.height / 4) - (borderSize * 2)

	return dp.paneBorder(index + 1).
		Width(size.width).
		Height(dayPaneHeight).
		Render(dp.dates[index].Format("# Monday, January 2"))
}
