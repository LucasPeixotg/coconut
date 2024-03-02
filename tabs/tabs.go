package tabs

import (
	tea "github.com/charmbracelet/bubbletea"
)

type pairing struct {
	title   string
	content tea.Model
}

type Model struct {
	list    []pairing
	current int
}

func New() Model {
	return Model{
		current: 0,
	}
}

// tea.Model implementation

func (model Model) Init() tea.Cmd {
	return nil
}

func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return model, nil
}

func (model Model) View() string {
	return ""
}

// own functions

func (model Model) Content() tea.Model {
	return model.list[model.current].content
}

func (model Model) IsEmpty() bool {
	return len(model.list) == 0
}
