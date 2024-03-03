package editor

func calcExplorerWidth(width int) int {
	return int(float32(width) * expWidthRatio)
}
