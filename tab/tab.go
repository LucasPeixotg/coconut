package tab

import (
	"coconut/models"

	"github.com/charmbracelet/lipgloss"
)

type Tab struct {
	title   string
	Content *models.Editor
}

type borderSide struct {
	top    string
	middle string
	bottom string
}

func (t *Tab) SetTitle(title string) {
	t.title = " " + title + " "
}

func borderfy(style lipgloss.Style, left, right borderSide, top, bottom string) lipgloss.Style {
	return style.Border(lipgloss.Border{
		Top:         top,
		Bottom:      bottom,
		Left:        left.middle,
		Right:       right.middle,
		TopLeft:     left.top,
		TopRight:    right.top,
		BottomLeft:  left.bottom,
		BottomRight: right.bottom,
	})
}

func (t Tab) View(index, activeIndex, tabQuant int) string {
	var style lipgloss.Style
	var leftBorder borderSide
	var rightBorder borderSide

	if index == activeIndex {
		if index == 0 {
			leftBorder = activeBorderStartLeft
		} else {
			leftBorder = activeBorderLeft
		}
		if index == tabQuant-1 {
			rightBorder = activeBorderEndRight
		} else {
			rightBorder = activeBorderRight
		}

		style = borderfy(activeTabStyle, leftBorder, rightBorder, "─", " ")
	} else {
		if index == 0 {
			leftBorder = inactiveBorderStartLeft
		} else {
			leftBorder = hiddenBorder
		}

		if index == tabQuant-1 {
			rightBorder = inactiveBorderEndRight
		} else {
			rightBorder = inactiveBorderRight
		}

		if index+1 == activeIndex {
			rightBorder = hiddenBorder
		} else if index-1 == activeIndex {
			leftBorder = hiddenBorder
		}

		style = borderfy(inactiveTabStyle, leftBorder, rightBorder, "─", "─")
	}

	return style.Render(t.title)
}

var inactiveTabStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#7d7d7d")).
	Padding(0, 2, 0, 2).
	BorderForeground(lipgloss.Color("167"))

var activeTabStyle = lipgloss.NewStyle().
	Inherit(inactiveTabStyle).
	Foreground(lipgloss.Color("#dbdbdb"))

var (
	activeBorderLeft      = borderSide{top: "┬", middle: "│", bottom: "╯"}
	activeBorderStartLeft = borderSide{top: "┌", middle: "│", bottom: "│"}
	activeBorderRight     = borderSide{top: "┬", middle: "│", bottom: "╰"}
	activeBorderEndRight  = borderSide{top: "╮", middle: "│", bottom: "╵"}

	inactiveBorderStartLeft = borderSide{top: "┌", middle: "│", bottom: "├"}
	inactiveBorderRight     = borderSide{top: "┬", middle: "│", bottom: "┴"}
	inactiveBorderEndRight  = borderSide{top: "╮", middle: "│", bottom: "╯"}

	hiddenBorder = borderSide{top: "─", middle: " ", bottom: "─"}
)
