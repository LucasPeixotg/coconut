package main

import (
	"coconut/editor"
	"coconut/tab"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	SaveFile key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.NewFile, k.OpenFile, k.OpenDir}, {k.Help, k.Quit}}
}

type FilepickerType uint8

const (
	None FilepickerType = iota
	NewFile
	OpenFile
	OpenDir
)

// main model: indicates the current state
type model struct {
	keys             keyMap
	help             help.Model
	filepickerStatus FilepickerType
	filepicker       filepicker.Model
	quitting         bool
	width            int
	height           int
	tabs             []tab.Tab
	activeTab        int
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
		SaveFile: key.NewBinding(
			key.WithKeys("ctrl+s"),
			key.WithHelp("ctrl+s", "save file"),
		),
	}

	m := &model{
		help:       help.New(),
		keys:       keys,
		filepicker: filepicker.New(),
	}

	m.filepicker.CurrentDirectory, _ = os.UserHomeDir()
	m.filepicker.ShowHidden = true
	m.filepicker.Styles.Permission = lipgloss.NewStyle().Width(0)

	return m
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

func (m *model) newEditor(path string) {
	tab := tab.Tab{}

	var err error
	tab.Content, err = editor.NewFileEditor(m.width, m.height-3, "teste.md")

	tab.SetTitle(path)

	// this panic is temporary (just for tests)
	if err != nil {
		panic("error while creating file: " + err.Error())
	}

	m.tabs = append(m.tabs, tab)
}

func (m *model) loadEditor(path string) {
	tab := tab.Tab{}

	var err error
	tab.Content, err = editor.OpenFileEditor(m.width, m.height-3, path)

	tab.SetTitle("README.md")

	// this panic is temporary (just for tests)
	if err != nil {
		panic("error while opening file: " + err.Error())
	}

	m.tabs = append(m.tabs, tab)
}

// initializes the event loop
func (m model) Init() tea.Cmd {
	return nil
}

// updates the model based on received messages
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var editorCmd tea.Cmd
	var filepickerCmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.setSizing(msg.Width, msg.Height)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			// toggle full help
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Quit):
			if m.filepickerStatus != None {
				// closes filepicker instead of the entire app
				m.filepickerStatus = None
				return m, tea.ClearScreen
			} else {
				// closes app
				m.quitting = true
				return m, tea.Quit
			}
		case key.Matches(msg, m.keys.NewFile):
			// pick directory
			if m.filepickerStatus != NewFile {
				m.filepickerStatus = NewFile
				m.filepicker.DirAllowed = true
				m.filepicker.FileAllowed = false
				m.filepicker.Height = m.height / 2

				return m, m.filepicker.Init()
			}
		case key.Matches(msg, m.keys.OpenFile):
			if m.filepickerStatus != OpenFile {
				m.filepickerStatus = OpenFile
				m.filepicker.DirAllowed = false
				m.filepicker.FileAllowed = true
				m.filepicker.Height = m.height / 2

				return m, m.filepicker.Init()
			}
		case key.Matches(msg, m.keys.NextTab):
			// change active tab
			m.nextTab()
			cmd = m.tabs[m.activeTab].Content.Focus()
		case key.Matches(msg, m.keys.SaveFile):
			// save current editor
			if len(m.tabs) > 0 {
				err := m.tabs[m.activeTab].Content.Save()

				// temporary panic just for tests
				if err != nil {
					panic("error while saving file: " + err.Error())
				}
			}
		}
	}

	// updates editor
	if len(m.tabs) > 0 {
		_, editorCmd = m.tabs[m.activeTab].Content.Update(msg)
	}

	// updates filepicker
	m.filepicker, filepickerCmd = m.filepicker.Update(msg)

	// check if it was selected
	if m.filepickerStatus != None {
		if selected, path := m.filepicker.DidSelectFile(msg); selected {
			switch m.filepickerStatus {
			case NewFile:
				m.newEditor(path)
				cmd = m.tabs[len(m.tabs)-1].Content.Focus()
			case OpenFile:
				m.loadEditor(path)
				cmd = m.tabs[len(m.tabs)-1].Content.Focus()
			}
			m.filepickerStatus = None
		}
	}

	return m, tea.Batch(cmd, editorCmd, filepickerCmd)
}

// renders UI string, based on the current state
func (m model) View() string {
	if m.quitting {
		return "See you :)"
	}

	var content string

	if len(m.tabs) == 0 {
		// render help if there isn't an open tab
		content += lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, m.help.View(m.keys))
	} else {
		// render all tabs
		tabCount := len(m.tabs)
		for i := 0; i < tabCount; i++ {
			content = lipgloss.JoinHorizontal(lipgloss.Left, content, m.tabs[i].View(i, m.activeTab, tabCount))
		}

		// render editor content
		content = lipgloss.JoinVertical(lipgloss.Top, content, m.tabs[m.activeTab].Content.View())
	}

	// render filepicker
	if m.filepickerStatus != None {
		var title string
		switch m.filepickerStatus {
		case OpenFile:
			title = "Select file"
		default:
			title = "Select directory"
		}

		//title = m.filepicker.CurrentDirectory
		content = lipgloss.Place(
			m.width, m.height,
			lipgloss.Center, lipgloss.Center,
			lipgloss.JoinVertical(
				lipgloss.Left,
				defaultTitleStyle.Render(title), "\n\n",
				m.filepicker.View(),
			))
	}

	return content
}

/* STYLES */
var (
	defaultTitleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#cccccc")).Bold(true)
)
