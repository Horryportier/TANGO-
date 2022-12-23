package app

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	// theme
	primaryColor    = lipgloss.Color("#8abc83")
	secondaryColor  = lipgloss.Color("#909dc0")
	backgroundColor = lipgloss.Color("#2a2a2a")
	accentColor1    = lipgloss.Color("#00c940")
	accentColor2    = lipgloss.Color("#ff2a42")

	appStyle = lipgloss.NewStyle().
			Padding(1, 2).Align(lipgloss.Center)

	accentStyle1 = lipgloss.NewStyle().
			Background(backgroundColor).
			Foreground(accentColor1)

	PrimaryStyle = lipgloss.NewStyle().
			Background(backgroundColor).
			Foreground(primaryColor)

	SecondaryStyle = lipgloss.NewStyle().
			Background(backgroundColor).
			Foreground(secondaryColor)

	PromptStyle = lipgloss.
			NewStyle().
			Background(backgroundColor).
			Foreground(accentColor1).
			Align(lipgloss.Left).
			PaddingLeft(2)

	Text = lipgloss.
		NewStyle().
		Foreground(primaryColor).
		Background(backgroundColor)
)
