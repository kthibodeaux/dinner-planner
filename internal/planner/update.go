package planner

import (
	"slices"

	tea "github.com/charmbracelet/bubbletea"
)

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
		case dp.keys.Down:
			dp.handleDown()
		case dp.keys.Up:
			dp.handleUp()
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
		dp.paneFocusIndex = index
		dp.mode = ModeAssign
	} else {
		dp.assign(index)
	}
}

func (dp *dinnerPlan) assign(targetPaneIndex int) {
	if dp.paneFocusIndex == targetPaneIndex {
		return
	}

	sourceRecipeList := dp.recipeLists[dp.paneFocusIndex]
	targetRecipeList := dp.recipeLists[targetPaneIndex]

	if len(sourceRecipeList.Recipes) == 0 {
		return
	}

	if targetPaneIndex == 0 {
		sourceRecipeList.Recipes = slices.Delete(sourceRecipeList.Recipes, sourceRecipeList.SelectedIndex, sourceRecipeList.SelectedIndex+1)
		return
	}

	sourceRecipe := sourceRecipeList.Recipes[sourceRecipeList.SelectedIndex]
	targetRecipeList.Recipes = append(targetRecipeList.Recipes, sourceRecipe)

	if dp.paneFocusIndex != 0 {
		sourceRecipeList.Recipes = slices.Delete(sourceRecipeList.Recipes, sourceRecipeList.SelectedIndex, sourceRecipeList.SelectedIndex+1)
	}

	if sourceRecipeList.SelectedIndex > len(sourceRecipeList.Recipes)-1 {
		sourceRecipeList.SelectedIndex = len(sourceRecipeList.Recipes) - 1
	}
}

func (dp *dinnerPlan) handleDown() {
	dp.recipeLists[dp.paneFocusIndex].handleDown()
}

func (dp *dinnerPlan) handleUp() {
	dp.recipeLists[dp.paneFocusIndex].handleUp()
}
