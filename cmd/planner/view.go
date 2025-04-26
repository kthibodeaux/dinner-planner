package main

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	borderSize = 1

	paneBorder = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder())

	pane = lipgloss.NewStyle()
)

func (dp dinnerPlan) View() string {
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

func (dp *dinnerPlan) recipeColumn(size Size, gap int) string {
	return paneBorder.
		Width(size.width).
		Height(size.height).
		MarginRight(gap).
		Render("")
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

	return paneBorder.
		Width(size.width).
		Height(dayPaneHeight).
		Render(dp.dates[index].Format("Monday, January 2"))
}
