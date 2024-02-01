package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	// new program
	p := tea.NewProgram(newModel(),
		tea.WithAltScreen(), // fullscreen
	)

	// start
	_, err := p.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// keybindings
type keyMap struct {
	Quit  key.Binding
	Help  key.Binding
	OpenF key.Binding
	OpenD key.Binding
	New   key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.New, k.OpenF, k.OpenD}, {k.Help, k.Quit}}
}

// main model: indicates the current state
type model struct {
	keys     keyMap
	help     help.Model
	quitting bool
	width    int
	height   int
}

func newModel() *model {
	keys := keyMap{
		Quit: key.NewBinding(
			key.WithKeys("esc", "ctrl+c"),
			key.WithHelp("esc", "quit"),
		),
		Help: key.NewBinding(
			key.WithKeys("?", "h"),
			key.WithHelp("?", "toggle help"),
		),
		New: key.NewBinding(
			key.WithKeys("ctrl+n"),
			key.WithHelp("crtl+n", "new file"),
		),
		OpenF: key.NewBinding(
			key.WithKeys("ctrl+o"),
			key.WithHelp("ctrl+o", "open file"),
		),
		OpenD: key.NewBinding(
			key.WithKeys("ctrl+d"),
			key.WithHelp("ctrl+d", "open dir"),
		),
	}

	return &model{
		help: help.New(),
		keys: keys,
	}
}

// set correct window size on all relevant components
func (m *model) setSizing(width, height int) {
	m.help.Width = width

	m.width = width
	m.height = height
}

// initializes the event loop
func (m model) Init() tea.Cmd {
	return nil
}

// updates the model based on received messages
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.setSizing(msg.Width, msg.Height)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			// toggle full help
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Quit):
			// closes app
			m.quitting = true
			return m, tea.Quit
		}
	}

	return m, nil
}

// renders UI string, based on the current state
func (m model) View() string {
	if m.quitting {
		return "See you :)"
	}

	// centralize help
	s := lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, m.help.View(m.keys))

	return s
}
