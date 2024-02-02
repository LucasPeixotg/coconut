package editor

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var editorStyle = textarea.Style{
	Prompt: lipgloss.NewStyle().
		Foreground(lipgloss.Color("#7d7d7d")),
}

type Editor struct {
	textarea textarea.Model
	content  string
}

func NewEditor(width, height int) *Editor {
	textarea := textarea.New()
	textarea.SetHeight(height)
	textarea.SetWidth(width)
	textarea.MaxHeight = 9999
	textarea.FocusedStyle = editorStyle
	textarea.BlurredStyle = editorStyle
	textarea.Prompt = "│"

	return &Editor{
		content:  "",
		textarea: textarea,
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
