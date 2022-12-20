package app

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	// theme
	primaryColor    = lipgloss.Color("#00FFD2")
	secondaryColor  = lipgloss.Color("#FF4499")
	backgroundColor = lipgloss.Color("#333333")
	accentColor1    = lipgloss.Color("#0a0047")
	accentColor2    = lipgloss.Color("#004687")

	PromptStyle = lipgloss.
                        NewStyle().
			Background(backgroundColor).
			Foreground(primaryColor).
			Align(lipgloss.Left)

	Background = lipgloss.
                        NewStyle().
			Foreground(backgroundColor)
	Text = lipgloss.
                NewStyle().
		Foreground(primaryColor)
)
