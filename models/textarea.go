package models

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type keymap struct {
	selection key.Binding
	end       key.Binding
	start     key.Binding
	newline   key.Binding
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

type Model struct {
	cursors []cursor
	keys    keymap
	viewy   int
	viewx   int

	// temporary
	// it will be updated in the future to use a better data structure
	content []string
}

func New(width, height int) *Model {
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
	}

	m := &Model{
		keys:  keys,
		viewy: 0,
		viewx: 0,
	}

	m.cursors = append(m.cursors, cursor{0, 0, 0})

	return m
}

func (model *Model) Init() tea.Cmd {
	return nil
}

func (model *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, model.keys.newline):
			model.newLine()
		default:
			model.write()
		}

	}
}

func (model *Model) View() string {
	return ""
}

func (model *Model) LoadFile(filepath string) {}

// cursors and editing functions
func (model *Model) write(text string) {
	for _, c := range model.cursors {
		model.content[c.line][c.start] += 
	}
}
