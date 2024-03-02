package main

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	NewFile    key.Binding
	OpenFile   key.Binding
	NewFolder  key.Binding
	OpenFolder key.Binding
	Quit       key.Binding
}

// used just to implement help.keyMap interface
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{}
}

// FullHelp returns keybindings for the help view.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.NewFile, k.OpenFile},
		{k.NewFolder, k.NewFolder, k.Quit},
	}
}

var keys = keyMap{
	NewFile: key.NewBinding(
		key.WithKeys("ctrl+n"),
		key.WithHelp("ctrl + n", "new file"),
	),
	OpenFile: key.NewBinding(
		key.WithKeys("ctrl+o"),
		key.WithHelp("ctrl + o", "open file"),
	),
	NewFolder: key.NewBinding(
		key.WithKeys("ctrl+d"),
		key.WithHelp("ctrl  + d", "open directory"),
	),
	OpenFolder: key.NewBinding(
		key.WithKeys("ctrl+f"),
		key.WithHelp("ctrl + f", "new directory"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+q"),
		key.WithHelp("ctrl + q", "quit"),
	),
}
