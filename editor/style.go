package editor

import "github.com/charmbracelet/lipgloss"

const tabHeight = 3
const expWidthRatio = float32(.2)

func getHelpStyle(width, height int) lipgloss.Style {
	return lipgloss.NewStyle().
		Align(lipgloss.Center, lipgloss.Center).
		Width(width).
		Height(height).
		Background(lipgloss.Color("#000000"))
}
