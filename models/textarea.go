package models

import (
	"fmt"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type keymap struct {
	selection key.Binding
	end       key.Binding
	start     key.Binding
	newline   key.Binding

	// cursor movement
	cursorUp       key.Binding
	cursorDown     key.Binding
	cursorBackward key.Binding
	cursorForward  key.Binding
	//wordForward  key.Binding
	//wordBackward key.Binding

	//
	//viewUp   key.Binding
	//viewDown key.Binding
}

// cursorObject
type cursorObject struct {
	line   int
	start  int
	lenght int
	cModel cursor.Model
}

var cursorStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#256")).
	Blink(true).
	Foreground(lipgloss.Color("#ccc"))

// model that holds current state
type textareaModel struct {
	cursors []*cursorObject
	keys    keymap
	viewy   int
	viewx   int
	prompt  string

	// temporary
	// it will be updated in the future to use a better data structure (probably piece table)
	content []string
}

func newTextarea(width, height int) *textareaModel {
	keys := keymap{
		selection: key.NewBinding(
			key.WithKeys("shift"),
		),
		end: key.NewBinding(
			key.WithKeys("end"),
		),
		start: key.NewBinding(
			key.WithKeys("home"),
		),
		newline: key.NewBinding(
			key.WithKeys("enter"),
		),
		cursorForward: key.NewBinding(
			key.WithKeys("right"),
		),
		cursorBackward: key.NewBinding(
			key.WithKeys("left"),
		),
		cursorUp: key.NewBinding(
			key.WithKeys("up"),
		),
		cursorDown: key.NewBinding(
			key.WithKeys("down"),
		),
	}

	m := &textareaModel{
		keys:   keys,
		viewy:  0,
		viewx:  0,
		prompt: "│ %-3d  ",
	}

	m.cursors = append(m.cursors, &cursorObject{0, 0, 0, cursor.New()})
	m.content = append(m.content, "")

	return m
}

// tea Model interface
func (model *textareaModel) Init() tea.Cmd {
	return nil
}

func (model *textareaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, model.keys.newline):
			model.newLine()
		case key.Matches(msg, model.keys.cursorUp):
			// TODO:

		case key.Matches(msg, model.keys.cursorDown):
			// TODO:

		default:
			model.write(msg.Runes)
		}
	}

	for i := 0; i < len(model.cursors); i++ {
		var temp tea.Cmd
		model.cursors[i].cModel, temp = model.cursors[i].cModel.Update(msg)
		cmds = append(cmds, temp)
	}

	return model, tea.Batch(cmds...)
}

func (model *textareaModel) View() string {
	content := ""
	for i, val := range model.content {
		content += fmt.Sprintf(model.prompt, i)
		content += val + "\n"
	}

	return content
}

/*
file related functions
*/
func (model *textareaModel) LoadFile(filepath string) {}

// cursors and editing functions
func (model *textareaModel) write(runes []rune) {
	for _, r := range runes {
		value := string(r)
		for _, c := range model.cursors {
			model.content[c.line] = model.content[c.line][:c.start] + value + model.content[c.line][c.start:]
			c.start++
		}
	}
}

func (model *textareaModel) newLine() {
	// TODO: implement this on every cursor
	lastCursor := model.cursors[len(model.cursors)-1]

	var new_content []string
	for i := 0; i < lastCursor.line; i++ {
		new_content = append(new_content, model.content[i])
	}

	new_content = append(new_content, model.content[lastCursor.line][:lastCursor.start])
	new_content = append(new_content, model.content[lastCursor.line][lastCursor.start:])

	for i := lastCursor.line + 1; i < len(model.content); i++ {
		new_content = append(new_content, model.content[i])
	}
	model.content = new_content

	lastCursor.line++
	lastCursor.start = 0
}
