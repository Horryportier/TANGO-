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


        appStyle = lipgloss.NewStyle(). 
                        Padding(1,2)

	PromptStyle = lipgloss.
			NewStyle().
			Background(backgroundColor).
			Foreground(primaryColor).
			Align(lipgloss.Left).
                        PaddingLeft(2)

	Text = lipgloss.
		NewStyle().
		Foreground(primaryColor).
                Background(secondaryColor)
)
