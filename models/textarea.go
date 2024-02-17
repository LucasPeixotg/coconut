package models

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type keymap struct {
	selection key.Binding
	end       key.Binding
	start     key.Binding
	newline   key.Binding

	up    key.Binding
	right key.Binding
	down  key.Binding
	left  key.Binding

	//viewup    key.Binding
	//viewdown  key.Binding
	//viewright key.Binding
	//viewleft  key.Binding
}

type cursor struct {
	line   int
	start  int
	lenght int
}

type textareaModel struct {
	cursors []cursor
	keys    keymap
	viewy   int
	viewx   int
	prompt  string

	// temporary
	// it will be updated in the future to use a better data structure
	content []string

	// for debug purposes
	//last_runes []rune
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
		up: key.NewBinding(
			key.WithKeys("up"),
		),
		right: key.NewBinding(
			key.WithKeys("right"),
		),
		down: key.NewBinding(
			key.WithKeys("down"),
		),
		left: key.NewBinding(
			key.WithKeys("left"),
		),
	}

	m := &textareaModel{
		keys:   keys,
		viewy:  0,
		viewx:  0,
		prompt: "│ %-3d  ",
	}

	m.cursors = append(m.cursors, cursor{0, 0, 0})
	m.content = append(m.content, "")

	return m
}

// tea Model interface
func (model *textareaModel) Init() tea.Cmd {
	return nil
}

func (model *textareaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, model.keys.newline):
			model.newLine()
		case key.Matches(msg, model.keys.up):
			// TODO:

		case key.Matches(msg, model.keys.down):
			// TODO:

		default:
			model.write(msg.Runes)
		}
	}
	return model, nil
}

func (model *textareaModel) View() string {
	content := ""
	for i, val := range model.content {
		content += fmt.Sprintf(model.prompt, i)
		content += val + "\n"
	}
	return content
}

// file related functions
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
	last_cursor := model.cursors[len(model.cursors)]

	var new_content []string
	for i := 0; i < last_cursor.line; i++ {
		new_content = append(new_content, model.content[i])
	}

	new_content = append(new_content, model.content[last_cursor.line][:last_cursor.start])
	new_content = append(new_content, model.content[last_cursor.line][last_cursor.start:])

	for i := last_cursor.line + 1; i < len(model.content); i++ {
		new_content = append(new_content, model.content[i])
	}
	model.content = new_content

	last_cursor.line++
	last_cursor.start = 0
}
