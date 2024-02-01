package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// new program
	p := tea.NewProgram(NewModel(),
		tea.WithAltScreen(), // fullscreen
	)

	// start
	_, err := p.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// main model: indicates the current state
type Model struct{}

func NewModel() *Model {
	return &Model{}
}

// initializes the event loop
func (m Model) Init() tea.Cmd {
	return nil
}

// updates the model based on received messages
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			// closes app on [esc] press
			return m, tea.Quit
		}
	}

	return m, nil
}

// renders UI string, based on the current state
func (m Model) View() string {
	return "test"
}
