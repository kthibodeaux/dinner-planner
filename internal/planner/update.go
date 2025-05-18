package planner

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kthibodeaux/dinner-planner/internal/config"
)

func (dp dinnerPlan) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case config.Get().Planner.Keys.Help:
			dp.mode = ModeHelp
		case config.Get().Planner.Keys.ShoppingList:
			dp.mode = ModeShoppingList
		case config.Get().Planner.Keys.MainView:
			dp.mode = ModeAssign
		case config.Get().Planner.Keys.Focus:
			dp.mode = ModeNavigatePane
		case config.Get().Planner.Keys.Recipes:
			dp.handlePane(0)
		case config.Get().Planner.Keys.Day1:
			dp.handlePane(1)
		case config.Get().Planner.Keys.Day2:
			dp.handlePane(2)
		case config.Get().Planner.Keys.Day3:
			dp.handlePane(3)
		case config.Get().Planner.Keys.Day4:
			dp.handlePane(4)
		case config.Get().Planner.Keys.Day5:
			dp.handlePane(5)
		case config.Get().Planner.Keys.Day6:
			dp.handlePane(6)
		case config.Get().Planner.Keys.Day7:
			dp.handlePane(7)
		case config.Get().Planner.Keys.Down:
			dp.handleDown(false)
		case config.Get().Planner.Keys.Up:
			dp.handleUp(false)
		case config.Get().Planner.Keys.ScrollDown:
			dp.handleDown(true)
		case config.Get().Planner.Keys.ScrollUp:
			dp.handleUp(true)
		case config.Get().Planner.Keys.Quit:
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
		sourceRecipeList.remove()
	} else {
		targetRecipeList.add(sourceRecipeList.selectedRecipe())

		if dp.paneFocusIndex != 0 {
			sourceRecipeList.remove()
		}
	}

	sourceRecipeList.checkSelectedIndex()
}

func (dp *dinnerPlan) handleDown(isScroll bool) {
	dp.recipeLists[dp.paneFocusIndex].handleDown(isScroll)
}

func (dp *dinnerPlan) handleUp(isScroll bool) {
	dp.recipeLists[dp.paneFocusIndex].handleUp(isScroll)
}
