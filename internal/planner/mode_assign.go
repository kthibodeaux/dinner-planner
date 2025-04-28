package planner

import "github.com/charmbracelet/lipgloss"

func (dp *dinnerPlan) viewModeAssign() string {
	unit := dp.size.width / 10
	gap := dp.size.width - (unit * 10)
	recipeColumnWidth := (unit * 6) - (borderSize * 2) + gap
	daysColumnsWidth := (unit * 4)
	columnHeight := dp.size.height - (borderSize * 2)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		dp.recipeColumn(Size{width: recipeColumnWidth, height: columnHeight}),
		dp.dayColumns(Size{width: daysColumnsWidth, height: columnHeight}),
	)
}

func (dp *dinnerPlan) recipeColumn(size Size) string {
	return dp.paneBorder(0).
		Width(size.width).
		Height(size.height).
		Render(dp.paneHeader(dp.keys.Recipes, "Recipes"))
}

func (dp *dinnerPlan) dayColumns(size Size) string {
	dayColumnWidth := size.width/2 - (borderSize * 2)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		dp.daysLeftColumn(Size{width: dayColumnWidth, height: size.height}),
		dp.daysRightColumn(Size{width: dayColumnWidth, height: size.height}),
	)
}

func (dp *dinnerPlan) daysLeftColumn(size Size) string {
	days := []string{}

	for _, dayNum := range []int{0, 1, 2, 3} {
		days = append(days, dp.dayPane(size, dayNum))
	}

	return lipgloss.JoinVertical(lipgloss.Left, days...)
}

func (dp *dinnerPlan) daysRightColumn(size Size) string {
	days := []string{}

	for _, dayNum := range []int{4, 5, 6} {
		days = append(days, dp.dayPane(size, dayNum))
	}

	return lipgloss.JoinVertical(lipgloss.Left, days...)
}

func (dp *dinnerPlan) dayPane(size Size, index int) string {
	dayPaneHeight := (dp.size.height / 4) - (borderSize * 2)

	return dp.paneBorder(index + 1).
		Width(size.width).
		Height(dayPaneHeight).
		Render(dp.paneHeader(dp.dayKeyMap[index], dp.dates[index].Format("Monday, January 2")))
}
