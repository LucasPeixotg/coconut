package explorer

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	width  int
	height int
	style  lipgloss.Style
}

func New(width, height int) Model {
	return Model{
		width:  width,
		height: height,
		style:  getStyle(width, height),
	}
}

// implements tea.Model interface

func (model Model) Init() tea.Cmd {
	return nil
}

func (model Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return model, nil
}

func (model Model) View() string {
	content := "Explorer\n\n"

	return model.style.Render(content)
}
