package main

import tea "github.com/charmbracelet/bubbletea"

func (dp dinnerPlan) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case dp.keys.FocusRecipes:
			dp.setFocus(0)
		case dp.keys.FocusDay1:
			dp.setFocus(1)
		case dp.keys.FocusDay2:
			dp.setFocus(2)
		case dp.keys.FocusDay3:
			dp.setFocus(3)
		case dp.keys.FocusDay4:
			dp.setFocus(4)
		case dp.keys.FocusDay5:
			dp.setFocus(5)
		case dp.keys.FocusDay6:
			dp.setFocus(6)
		case dp.keys.FocusDay7:
			dp.setFocus(7)
		case dp.keys.Quit:
			return dp.quit()
		}
	case tea.WindowSizeMsg:
		dp.size.width = msg.Width
		dp.size.height = msg.Height
	}

	return dp, nil
}

func (dp *dinnerPlan) quit() (tea.Model, tea.Cmd) {
	return dp, tea.Quit
}

func (dp *dinnerPlan) setFocus(index int) {
	dp.focusIndex = index
}
