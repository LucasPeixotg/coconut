package tabs

import "github.com/charmbracelet/lipgloss"

var tabStyle = lipgloss.NewStyle().
	Border(lipgloss.BlockBorder(), false, true, false, false)

func getStyle(width, height int) lipgloss.Style {
	return lipgloss.NewStyle().
		Width(width).
		Height(height-1).
		MaxWidth(width).
		MaxHeight(height).
		Background(lipgloss.Color("#000000")).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(lipgloss.Color("5"))
}
