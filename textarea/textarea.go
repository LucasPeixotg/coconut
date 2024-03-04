package textarea

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	filepath string

	head     *node
	viewNode *node
	cursors  []*cursor

	currentViewHeight int
	maxViewHeight     int
	offsetHeight      int

	width  int
	height int
	style  lipgloss.Style
}

func New(width, height int) Model {
	m := Model{
		width:             width,
		height:            height,
		style:             getStyle(width, height),
		head:              newNode(""),
		currentViewHeight: 1,
		offsetHeight:      0,
	}
	m.viewNode = m.head
	m.maxViewHeight = height - m.style.GetPaddingTop() - m.style.GetPaddingBottom() - m.style.GetMarginBottom() - m.style.GetMarginTop()
	m.cursors = append(m.cursors, &cursor{m.head, 0, 0})
	m.head.cursor = m.cursors[0]

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

		c.line.cursor = nil
		c.line = c.line.next
		c.line.cursor = c

		if model.currentViewHeight+1 > model.maxViewHeight {
			model.viewNode = model.viewNode.next
			model.offsetHeight++
			model.currentViewHeight++
		} else {
			model.currentViewHeight++
		}
	}
}

func (model *Model) write() {
	// TODO:
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
		default:
			model.write()
		}
	}

	return model, nil
}

func (model Model) View() string {
	var content string

	current := model.viewNode
	for i := 0; i < model.currentViewHeight && current != nil; i++ {
		line := ""
		line += current.data + " \n"
		if current.cursor != nil {
			if current.cursor.len > 0 {
				// selected
			} else {
				// normal cursor
				line = lipgloss.StyleRunes(line, []int{current.cursor.start}, cursorStyle, textStyle)
			}
		}
		content += promptStyle.Render(fmt.Sprint(model.offsetHeight+i+1)) + line

		current = current.next
	}

	return model.style.Render(content)
}
