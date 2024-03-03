package editor

import "github.com/charmbracelet/lipgloss"

const backgroundColor = "#000000"

const tabHeight = 2
const expWidthRatio = float32(.15)

func getHelpStyle(width, height int) lipgloss.Style {
	return lipgloss.NewStyle().
		Width(width).
		Height(height).
		Align(lipgloss.Center, lipgloss.Center).
		Background(lipgloss.Color(backgroundColor))
}
