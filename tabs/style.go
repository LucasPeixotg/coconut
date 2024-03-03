package tabs

import "github.com/charmbracelet/lipgloss"

func getBlockStyle(width, height int) lipgloss.Style {
	return lipgloss.NewStyle().
		Width(width).
		Height(height-1).
		MaxWidth(width).
		MaxHeight(height).
		Background(lipgloss.Color("#000000")).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(lipgloss.Color("5"))
}

func getSingleTabStyle(height int) lipgloss.Style {
	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, true, false, false).
		AlignVertical(lipgloss.Center).
		Height(height).
		MaxHeight(height).
		Padding(0, 2).
		BorderForeground(lipgloss.Color("5"))
}
