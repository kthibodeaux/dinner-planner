package main

import tea "github.com/charmbracelet/bubbletea"

func (dp dinnerPlan) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case dp.keys.Quit:
			return dp.Quit()
		}
	}
	return dp, nil
}

func (dp *dinnerPlan) Quit() (tea.Model, tea.Cmd) {
	return dp, tea.Quit
}
