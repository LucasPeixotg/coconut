package main

import (
	"coconut/editor"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// tab bar style
var (
	activeTabBorder     = lipgloss.Border{"─", " ", "│", "│", "┬", "┬", "┘", "└", "", "", "", "", ""}
	inactiveTabBorder   = lipgloss.Border{"─", "─", "", "│", "─", "┬", "─", "┴", "", "", "", "", "─"}
	neighbourTabBorderR = lipgloss.Border{"─", "─", "", "│", "─", "┬", "─", "┴", "", "", "", "", "─"}
	neighbourTabBorderL = lipgloss.Border{
		Top:          "─",
		Bottom:       "─",
		Left:         "",
		Right:        "",
		TopLeft:      "─",
		TopRight:     "─",
		BottomLeft:   "─",
		BottomRight:  "─",
		MiddleLeft:   "",
		MiddleRight:  "",
		Middle:       "",
		MiddleTop:    "",
		MiddleBottom: "─"}

	tabStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7d7d7d")).
			Padding(0, 2, 0, 2).
			BorderForeground(lipgloss.Color("#7d7d7d"))
	activeTabStyle = lipgloss.NewStyle().
			Inherit(tabStyle).
		//BorderForeground(lipgloss.Color("#dbdbdb")).
		Foreground(lipgloss.Color("#dbdbdb")).
		Border(activeTabBorder)
	inactiveTabStyle = lipgloss.NewStyle().
				Inherit(tabStyle).
				Border(inactiveTabBorder)
	neighbourTabStyleR = lipgloss.NewStyle().
				Inherit(tabStyle).
				Border(neighbourTabBorderR)
	neighbourTabStyleL = lipgloss.NewStyle().
				Inherit(tabStyle).
				Border(neighbourTabBorderL)
)

// **
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
	Quit     key.Binding
	Help     key.Binding
	OpenFile key.Binding
	OpenDir  key.Binding
	NewFile  key.Binding
	NextTab  key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.NewFile, k.OpenFile, k.OpenDir}, {k.Help, k.Quit}}
}

// main model: indicates the current state
type model struct {
	keys      keyMap
	help      help.Model
	quitting  bool
	width     int
	height    int
	tabs      []*editor.Editor
	activeTab int
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
		NewFile: key.NewBinding(
			key.WithKeys("ctrl+n"),
			key.WithHelp("crtl+n", "new file"),
		),
		OpenFile: key.NewBinding(
			key.WithKeys("ctrl+o"),
			key.WithHelp("ctrl+o", "open file"),
		),
		OpenDir: key.NewBinding(
			key.WithKeys("ctrl+d"),
			key.WithHelp("ctrl+d", "open dir"),
		),
		NextTab: key.NewBinding(
			key.WithKeys("shift+tab"),
			key.WithHelp("shift+tab", "next tab"),
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

// changes view to next tab
func (m *model) nextTab() {
	if m.activeTab == len(m.tabs)-1 {
		m.activeTab = 0
	} else {
		m.activeTab++
	}
}

func (m *model) newEditor() {
	newEditor := editor.NewEditor()
	newEditor.Title = "  unnamed  "
	m.tabs = append(m.tabs, newEditor)
}

// initializes the event loop
func (m model) Init() tea.Cmd {
	return nil
}

// updates the model based on received messages
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var editor_cmd tea.Cmd
	if len(m.tabs) > 0 {
		_, editor_cmd = m.tabs[m.activeTab].Update(msg)
	}

	var cmd tea.Cmd
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
		case key.Matches(msg, m.keys.NewFile):
			m.newEditor()
			cmd = m.tabs[len(m.tabs)-1].Focus()
		case key.Matches(msg, m.keys.NextTab):
			// change active tab
			m.nextTab()
			cmd = m.tabs[m.activeTab].Focus()
		}
	}

	return m, tea.Batch(editor_cmd, cmd)
}

// renders UI string, based on the current state
func (m model) View() string {
	if m.quitting {
		return "See you :)"
	}

	var content string

	// centralize help
	if len(m.tabs) == 0 {
		content += lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, m.help.View(m.keys))
	} else {
		for i := 0; i < len(m.tabs); i++ {
			var tabstring string
			if i == m.activeTab {
				tabstring = activeTabStyle.Render(m.tabs[i].Title)
			} else if i+1 == m.activeTab {
				tabstring = neighbourTabStyleL.Render(m.tabs[i].Title)
			} else if i-1 == m.activeTab {
				tabstring = neighbourTabStyleR.Render(m.tabs[i].Title)
			} else {
				tabstring = inactiveTabStyle.Render(m.tabs[i].Title)
			}
			content = lipgloss.JoinHorizontal(lipgloss.Left, content, tabstring)
		}

		content = lipgloss.JoinVertical(lipgloss.Top, content, m.tabs[m.activeTab].View())
	}

	return content
}
