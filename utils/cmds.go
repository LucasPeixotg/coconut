package utils

import tea "github.com/charmbracelet/bubbletea"

type NewFileMsg struct{}

func NewFileCmd() tea.Msg {
	return NewFileMsg{}
}
