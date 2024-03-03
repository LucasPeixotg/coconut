package main

import (
	"fmt"
	"os"

	"github.com/LucasPeixotg/coconut/editor"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type State uint8

const (
	editorState State = iota
	filepickerState
	initializingState
	quittingState
)

type MainModel struct {
	state State

	width  int
	height int

	editorView *editor.Model
}

func New() MainModel {
	m := MainModel{}

	// initial model options
	m.state = initializingState
	return m
}

func (model *MainModel) SetSize(width, height int) {
	model.width = width
	model.height = height

	model.editorView = editor.New(keys, width, height)
}

// tea.Model implementation

func (model MainModel) Init() tea.Cmd {
	return tea.SetWindowTitle("Coconut")
}

func (model MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	// state independent keys
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		model.SetSize(msg.Width, msg.Height)

		if model.state == initializingState {
			model.state = editorState
		}
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			model.state = quittingState
			cmds = append(cmds, tea.Quit)
		}
	}

	return model, tea.Batch(cmds...)
}

func (model MainModel) View() string {
	switch model.state {
	case quittingState:
		return "bye!"
	case initializingState:
		return "initializing..."
	case editorState:
		return model.editorView.View()
	}

	return "unhandled state"
}

// run

func main() {
	if _, err := tea.NewProgram(New(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Printf("Uh Oh! Something unexpected happened: %v\n", err)
		os.Exit(1)
	}
}
