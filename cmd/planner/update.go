package main

import tea "github.com/charmbracelet/bubbletea"

func (dp dinnerPlan) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case dp.keys.Quit:
			return dp.Quit()
		case "0", "1", "2", "3", "4", "5", "6", "7":
			dp.focusIndex = int(msg.Runes[0] - '0')
		}
	case tea.WindowSizeMsg:
		dp.size.width = msg.Width
		dp.size.height = msg.Height
	}

	return dp, nil
}

func (dp *dinnerPlan) Quit() (tea.Model, tea.Cmd) {
	return dp, tea.Quit
}
