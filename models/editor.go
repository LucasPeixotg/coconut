package models

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Editor struct {
	textarea tea.Model
	filename string
}

func newEditor(width, height int, filename string) *Editor {
	textarea := newTextarea(width, height)

	return &Editor{
		filename: filename,
		textarea: textarea,
	}
}

func NewFileEditor(width, height int, filename string) (*Editor, error) {
	// TODO:
	editor := newEditor(width, height, filename)

	return editor, nil
}

func OpenFileEditor(width, height int, filename string) (*Editor, error) {
	// TODO:
	editor := newEditor(width, height, filename)

	return editor, nil
}

// tea model interface
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

// TODO:
func (e *Editor) Focus() tea.Cmd {
	return nil
}
