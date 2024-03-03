package tabs

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type pairing struct {
	title   string
	content tea.Model
}

type Model struct {
	list    []pairing
	current int

	width         int
	tabHeight     int
	contentHeight int
	style         lipgloss.Style
}

func New(width, tabHeight, contentHeight int) Model {
	return Model{
		current:       0,
		width:         width,
		tabHeight:     tabHeight,
		contentHeight: contentHeight,
		style:         getStyle(width, tabHeight),
	}
}

// tea.Model implementation

func (model Model) Init() tea.Cmd {
	return nil
}

func (model Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return model, nil
}

func (model Model) View() string {
	content := ""

	if !model.IsEmpty() {
		for _, tab := range model.list {
			tabString := lipgloss.PlaceVertical(model.tabHeight, lipgloss.Center, tab.title)
			tabString = tabStyle.Render(tabString)
			content = lipgloss.JoinHorizontal(0, content, tabString)
		}
	}

	return model.style.Render(content)
}

// own functions

func (model Model) Content() tea.Model {
	return model.list[model.current].content
}

func (model Model) IsEmpty() bool {
	return len(model.list) == 0
}
