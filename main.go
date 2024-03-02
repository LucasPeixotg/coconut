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
)

type MainModel struct {
	state    State
	quitting bool

	currentView tea.Model
}

func New() MainModel {
	m := MainModel{}
	// initial model options
	m.state = editorState
	m.currentView = editor.New(keys)
	return m
}

// tea.Model implementation

func (model MainModel) Init() tea.Cmd {
	return nil
}

func (model MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	// state independent keys
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			model.quitting = true
			cmds = append(cmds, tea.Quit)
		}

	}

	switch model.state {
	case editorState:
		// editor type assertion
		e, _ := model.currentView.(editor.Model)

		switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			e.SetSize(msg.Width, msg.Height)
		}
	}

	return model, tea.Batch(cmds...)
}

func (model MainModel) View() string {
	if model.quitting {
		return ""
	}

	return model.currentView.View()
}

// run

func main() {
	if _, err := tea.NewProgram(New(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Printf("Uh Oh! Something unexpected happened: %v\n", err)
		os.Exit(1)
	}
}
