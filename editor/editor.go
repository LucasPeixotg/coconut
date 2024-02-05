package editor

import (
	"os"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var editorStyle = textarea.Style{
	Prompt: lipgloss.NewStyle().
		Foreground(lipgloss.Color("#999999")),
	Text: lipgloss.NewStyle().
		Foreground(lipgloss.Color("#eeeeee")),
	LineNumber: lipgloss.NewStyle().
		Foreground(lipgloss.Color("#40575c")).
		Width(6).
		Align(lipgloss.Right).
		PaddingRight(2),
	CursorLineNumber: lipgloss.NewStyle().
		Foreground(lipgloss.Color("#a2d5e0")).
		Bold(true).
		Width(6).
		Align(lipgloss.Right).
		PaddingRight(2),
}

type Editor struct {
	textarea textarea.Model
	file     *os.File
	filename string
}

func newEditor(width, height int, filename string) *Editor {
	textarea := textarea.New()
	textarea.SetHeight(height)
	textarea.SetWidth(width)
	textarea.MaxHeight = 9999
	textarea.FocusedStyle = editorStyle
	textarea.BlurredStyle = editorStyle
	textarea.Prompt = "│"

	return &Editor{
		filename: filename,
		textarea: textarea,
	}
}

func NewFileEditor(width, height int, filename string) (*Editor, error) {
	editor := newEditor(width, height, filename)
	err := editor.newFile()
	return editor, err
}

func OpenFileEditor(width, height int, filename string) (*Editor, error) {
	editor := newEditor(width, height, filename)
	err := editor.loadFile()
	return editor, err
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

// ** file related functions
func (e *Editor) newFile() error {
	file, err := os.Create(e.filename)

	if err != nil {
		return err
	}

	e.file = file

	return nil
}

func (e *Editor) loadFile() error {
	file, err := os.OpenFile(e.filename, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	e.file = file

	buffer, err := os.ReadFile(e.filename)
	if err != nil {
		return err
	}

	e.textarea.SetValue(string(buffer))
	return nil
}

func (e *Editor) Save() error {
	e.file.Truncate(0)
	e.file.Seek(0, 0)
	_, err := e.file.Write([]byte(e.textarea.Value()))
	return err
}

func (e *Editor) Close() {
	if e.file != nil {
		e.file.Close()
	}
}
