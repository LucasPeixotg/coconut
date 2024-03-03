package explorer

import "github.com/charmbracelet/lipgloss"

func getStyle(width, height int) lipgloss.Style {
	return lipgloss.NewStyle().
		Width(width-1).
		MaxWidth(width).
		Height(height).
		MaxHeight(height).
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderForeground(lipgloss.Color("5")).
		BorderBackground(lipgloss.Color("#000000")).
		Background(lipgloss.Color("#000000"))
}
