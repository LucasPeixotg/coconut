package editor

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type Editor struct {
	Title    string
	textarea textarea.Model
	content  string
}

func NewEditor() *Editor {
	return &Editor{
		content:  "",
		textarea: textarea.New(),
	}
}

// initializes the event loop
func (e Editor) Init() tea.Cmd {
	return nil
}

func (e Editor) View() string {
	return e.textarea.View()
}

func (e *Editor) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	e.textarea, cmd = e.textarea.Update(msg)

	return e, cmd
}

func (e *Editor) Focus() tea.Cmd {
	return e.textarea.Focus()
}
