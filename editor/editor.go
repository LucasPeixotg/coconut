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

	helpStyle lipgloss.Style

	mainKeys help.KeyMap

	// focus state
	//	0 = tab content
	//	1 = explorer
	focusState uint8
}

func New(mainKeys help.KeyMap, width, height int) *Model {
	explorerWidth := calcExplorerWidth(width)
	contentWidth := width - explorerWidth
	contentHeight := height - tabHeight

	model := &Model{
		focusState: 0,
		mainKeys:   mainKeys,
		help:       help.New(),
		helpStyle:  getHelpStyle(contentWidth, contentHeight),
		explorer:   explorer.New(explorerWidth, height),
		tab:        tabs.New(contentWidth, tabHeight, contentHeight),
	}

	// show help in expanded view
	model.help.ShowAll = true

	return model
}

// tea.Model implementation

func (model Model) Init() tea.Cmd {
	return tea.ClearScreen
}

func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var tmp tea.Cmd

	model.explorer, tmp = model.explorer.Update(msg)
	cmds = append(cmds, tmp)

	model.tab, tmp = model.tab.Update(msg)
	cmds = append(cmds, tmp)

	return model, tea.Batch(cmds...)
}

func (model Model) View() string {
	var content string

	if model.tab.IsEmpty() {
		styledHelp := model.helpStyle.Render(model.help.View(model.mainKeys))
		content = lipgloss.JoinVertical(lipgloss.Center, model.tab.View(), styledHelp)
	}
	content = lipgloss.JoinHorizontal(0, model.explorer.View(), content)

	//debugger := fmt.Sprintf("width: %v | height: %v\n", model.width, model.height)

	return content
}
