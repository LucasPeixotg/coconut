package editor

import (
	"github.com/LucasPeixotg/coconut/explorer"
	"github.com/LucasPeixotg/coconut/tabs"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	tab      tabs.Model
	explorer explorer.Model
	help     help.Model

	// focus state
	//	0 = tab content
	//	1 = explorer
	focusState uint8
}

func New() Model {
	model := Model{
		tab:        tabs.New(),
		explorer:   explorer.New(),
		help:       help.New(),
		focusState: 0,
	}

	// show help in expanded view
	model.help.ShowAll = true

	return model
}

func (model Model) Init() tea.Cmd {
	return nil
}

func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	if model.tab.IsEmpty() {
		// show help

	} else {
		// show
	}

	return model, nil
}

func (model Model) View() string {
	return ""
}
