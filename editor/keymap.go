package editor

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	NewFile    key.Binding
	OpenFile   key.Binding
	NewFolder  key.Binding
	OpenFolder key.Binding
}

// used just to implement key.Map interface
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{}
}

// FullHelp returns keybindings for the help view.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.NewFile, k.OpenFile},
		{k.NewFolder, k.NewFolder},
	}
}

var keys = keyMap{
	NewFile: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	OpenFile: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	NewFolder: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "move left"),
	),
	OpenFolder: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "move right"),
	),
}
