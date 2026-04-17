package ui

import "github.com/charmbracelet/lipgloss"

var (
	cmdBarStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#1a1a2e")).
			Foreground(lipgloss.Color("#5DCAA5")).
			Padding(0, 1)

	cmdTextStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffffff"))

	hintStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888780"))

	errStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#E24B4A"))

	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#5DCAA5"))
)

func RenderCmdBar(cmdMode bool, cmdInput string, lastCmd string, lastCmdErr bool, width int) string {
	bar := cmdBarStyle.Width(width)

	if cmdMode {
		return bar.Render(":" + cmdTextStyle.Render(cmdInput) + "█")
	}

	if lastCmd != "" {
		if lastCmdErr {
			return bar.Render(errStyle.Render("error: " + lastCmd))
		}
		return bar.Render(successStyle.Render(lastCmd))
	}

	return bar.Render(hintStyle.Render("press : for commands  |  q to quit"))
}
