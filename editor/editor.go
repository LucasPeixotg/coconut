package editor

import (
	"github.com/LucasPeixotg/coconut/explorer"
	"github.com/LucasPeixotg/coconut/tabs"
	"github.com/LucasPeixotg/coconut/utils"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type focusState uint8

const (
	tabFocus focusState = iota
	explorerFocus
)

type Model struct {
	tab      tabs.Model
	explorer explorer.Model
	help     help.Model

	width  int
	height int

	helpStyle lipgloss.Style

	focusState focusState
}

func New(width, height int) Model {
	explorerWidth := calcExplorerWidth(width)
	contentWidth := width - explorerWidth
	contentHeight := height - tabHeight

	model := Model{
		focusState: tabFocus,

		help:      help.New(),
		helpStyle: getHelpStyle(contentWidth, contentHeight),
		explorer:  explorer.New(explorerWidth, height),
		tab:       tabs.New(contentWidth, tabHeight, contentHeight),

		width:  width,
		height: height,
	}

	// show help in expanded view
	model.help.ShowAll = true

	return model
}

// tea.Model implementation

func (model Model) Init() tea.Cmd {
	return tea.ClearScreen
}

func (model Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd
	var tmp tea.Cmd

	// state dependent
	switch model.focusState {
	case explorerFocus:
		model.explorer, tmp = model.explorer.Update(msg)
		cmds = append(cmds, tmp)
	case tabFocus:
		model.tab, tmp = model.tab.Update(msg)
		cmds = append(cmds, tmp)

	}

	return model, tea.Batch(cmds...)
}

func (model Model) View() string {
	content := ""

	if model.tab.IsEmpty() {
		styledHelp := model.helpStyle.Render(model.help.View(utils.FileKeys))
		content += lipgloss.JoinVertical(lipgloss.Center, model.tab.View(), styledHelp)
	} else {
		content += model.tab.View()
	}
	content = lipgloss.JoinHorizontal(0, model.explorer.View(), content)

	//debugger := fmt.Sprintf("width: %v | height: %v | tabs: %v\n", model.width, model.height, model.tab.Count())

	return content
}
