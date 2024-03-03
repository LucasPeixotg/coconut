package textarea

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	filepath string

	content *node
	cursors []cursor

	width  int
	height int
	style  lipgloss.Style
}

func New(width, height int) Model {
	m := Model{
		width:   width,
		height:  height,
		style:   getStyle(width, height),
		content: newNode(""),
		cursors: []cursor{},
	}
	m.cursors = append(m.cursors, cursor{m.content, 0, 0})

	return m
}

func (model *Model) Open(filepath string) {
	model.filepath = filepath
	// TODO: actually open the file
}

func (model *Model) newLine() {
	for _, c := range model.cursors {
		// TODO: put content in the next line
		appendNode(c.line, "")
		c.line = c.line.next
	}
}

// tea.Model implementation

func (model Model) Init() tea.Cmd {
	return nil
}

func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.NewLine):
			model.newLine()
		}
	}

	return model, nil
}

func (model Model) View() string {
	content := ""

	current := model.content
	for i := 0; i < model.height && current != nil; i++ {
		content += promptStyle.Render(fmt.Sprint(i)) + current.data + "\n"
		current = current.next
	}

	return model.style.Render(content)

	//return model.style.Render(content)
}
