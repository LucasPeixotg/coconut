package models

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type keymap struct {
	selection key.Binding
	end       key.Binding
	start     key.Binding
}

type cursor struct {
	start  int
	lenght int
}

type Model struct {
	width   int
	height  int
	cursors []cursor
}

func (model *Model) Init() tea.Cmd {

}

func (model *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

}

func (model *Model) View() string {

}

func (model *Model) LoadFile(filepath string) {}
