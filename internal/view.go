package internal

import "github.com/charmbracelet/lipgloss"

const (
	warningText = "WARNING"
	okText      = "OK"
	failedText  = "FAILED"
)

func coloredTestName(testname string) string {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#000000")).
		Background(lipgloss.Color("#FFFFFF"))
	return style.Render(testname)
}

func coloredOk() string {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#000000")).
		Background(lipgloss.Color("#FFFFFF"))
	return style.Render(okText)
}

func coloredWarning() string {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#000000")).
		Background(lipgloss.Color("#FFFFFF"))
	return style.Render(warningText)
}

func coloredFailed() string {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#000000")).
		Background(lipgloss.Color("#FFFFFF"))
	return style.Render(failedText)
}
