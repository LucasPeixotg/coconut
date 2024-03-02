package explorer

import tea "github.com/charmbracelet/bubbletea"

type Model struct{}

func New() Model {
	return Model{}
}

// implements tea.Model interface

func (model Model) Init() tea.Cmd {
	return nil
}

func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return model, nil
}

func (model Model) View() string {
	return ""
}
