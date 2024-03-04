package textarea

import "github.com/charmbracelet/lipgloss"

const backgroundColor = "#000000"

func getStyle(width, height int) lipgloss.Style {
	return lipgloss.NewStyle().
		Width(width).
		MaxWidth(width).
		Height(height).
		MaxHeight(height).
		Padding(1).
		Background(lipgloss.Color(backgroundColor))
}

var promptStyle = lipgloss.NewStyle().
	BorderBackground(lipgloss.Color(backgroundColor)).
	BorderForeground(lipgloss.Color("#444444")).
	Width(4).
	MaxWidth(5)

var cursorStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#eeeeee")).
	Foreground(lipgloss.Color(backgroundColor))

var textStyle = lipgloss.NewStyle().
	Background(lipgloss.Color(backgroundColor)).
	Foreground(lipgloss.Color("#eeeeee"))
