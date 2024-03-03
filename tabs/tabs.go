package tabs

import (
	"github.com/LucasPeixotg/coconut/textarea"
	"github.com/LucasPeixotg/coconut/utils"
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

	tabStyle   lipgloss.Style
	blockStyle lipgloss.Style
}

func New(width, tabHeight, contentHeight int) Model {
	return Model{
		current:       0,
		width:         width,
		tabHeight:     tabHeight,
		contentHeight: contentHeight,
		tabStyle:      getSingleTabStyle(tabHeight - 1),
		blockStyle:    getBlockStyle(width, tabHeight),
	}
}

func (model *Model) NewTab(content tea.Model, title string) {
	model.list = append(model.list, pairing{
		title,
		content,
	})

	if len(model.list) != 1 {
		model.current++
	}
}

func (model Model) Content() tea.Model {
	return model.list[model.current].content
}

func (model Model) IsEmpty() bool {
	return len(model.list) == 0
}

func (model Model) Count() int {
	return len(model.list)
}

// tea.Model implementation

func (model Model) Init() tea.Cmd {
	return nil
}

func (model Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg.(type) {
	case utils.NewFileMsg:
		text := textarea.New(model.width, model.contentHeight)
		model.NewTab(text, "untitled")
	}

	// update current content
	if !model.IsEmpty() {
		var tmp tea.Cmd
		model.list[model.current].content, tmp = model.list[model.current].content.Update(msg)

		cmds = append(cmds, tmp)
	}

	return model, tea.Batch(cmds...)
}

func (model Model) View() string {
	content := ""

	if !model.IsEmpty() {
		for _, tab := range model.list {
			tabString := model.tabStyle.Render(tab.title)
			content = lipgloss.JoinHorizontal(0, content, tabString)
		}
		content = model.blockStyle.Render(content)
		return lipgloss.JoinVertical(0, content, model.list[model.current].content.View())
		//return lipgloss.JoinVertical(0, content, "teste")
	} else {
		return model.blockStyle.Render(content)
	}

}
