package editor

import (
	"github.com/LucasPeixotg/coconut/explorer"
	"github.com/LucasPeixotg/coconut/tabs"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	tab      tabs.Model
	explorer explorer.Model
	help     help.Model

	mainKeys help.KeyMap

	width  int
	height int

	// focus state
	//	0 = tab content
	//	1 = explorer
	focusState uint8
}

func New(mainKeys help.KeyMap) Model {
	model := Model{
		tab:        tabs.New(),
		explorer:   explorer.New(),
		help:       help.New(),
		focusState: 0,
		mainKeys:   mainKeys,
	}

	// show help in expanded view
	model.help.ShowAll = true

	return model
}

func (model *Model) SetSize(width, height int) {
	model.width = width
	model.height = height
}

// tea.Model implementation

func (model Model) Init() tea.Cmd {
	return nil
}

func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//var cmds []tea.Cmd

	return model, nil
}

func (model Model) View() string {
	var content string

	if model.tab.IsEmpty() {
		content = lipgloss.JoinVertical(lipgloss.Left, model.tab.View(), model.help.View(model.mainKeys))
	}
	content = lipgloss.JoinHorizontal(0, model.explorer.View(), content)

	return content
}
