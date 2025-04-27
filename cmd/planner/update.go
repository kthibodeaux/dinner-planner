package main

import tea "github.com/charmbracelet/bubbletea"

func (dp dinnerPlan) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case dp.keys.Help:
			dp.mode = ModeHelp
		case dp.keys.MainView:
			dp.mode = ModeAssign
		case dp.keys.Focus:
			dp.mode = ModeNavigatePane
		case dp.keys.Recipes:
			dp.handlePane(0)
		case dp.keys.Day1:
			dp.handlePane(1)
		case dp.keys.Day2:
			dp.handlePane(2)
		case dp.keys.Day3:
			dp.handlePane(3)
		case dp.keys.Day4:
			dp.handlePane(4)
		case dp.keys.Day5:
			dp.handlePane(5)
		case dp.keys.Day6:
			dp.handlePane(6)
		case dp.keys.Day7:
			dp.handlePane(7)
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

func (dp *dinnerPlan) handlePane(index int) {
	if dp.mode == ModeNavigatePane {
		dp.focusIndex = index
		dp.mode = ModeAssign
	}
}
