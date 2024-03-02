package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type State uint8

const (
	editor State = iota
	filepicker
)

type MainModel struct {
	state  State
	width  int
	height int

	currentView tea.Model
}

func (model MainModel) Init() tea.Cmd {
	return nil
}

func (model MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch model.state {
	case initial:
		// state independent messages
		switch msg.(type) {
		case tea.WindowSizeMsg:

		}
	}

	return model, tea.Batch(cmds...)
}

func (model MainModel) View() string {
	return ""
}

func NewMainModel() MainModel {
	m := MainModel{}
	// initial model options
	m.state = editor
	return m
}

func main() {
	if _, err := tea.NewProgram(NewMainModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Printf("Uh Oh! Something unexpected happened: %v\n", err)
		os.Exit(1)
	}
}
