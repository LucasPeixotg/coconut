package textarea

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Up      key.Binding
	Down    key.Binding
	NewLine key.Binding
}

var keys = keyMap{
	Up:      key.NewBinding(key.WithKeys("up")),
	Down:    key.NewBinding(key.WithKeys("down")),
	NewLine: key.NewBinding(key.WithKeys("enter")),
}
