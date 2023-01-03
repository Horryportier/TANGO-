package app

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	// theme
	primaryColor    = lipgloss.Color("#00c940")
	secondaryColor  = lipgloss.Color("#909dc0")
	backgroundColor = lipgloss.Color("#2a2a2a")
	accentColor     = lipgloss.Color("#ff2a42")

	appStyle = lipgloss.NewStyle().
			PaddingLeft(4).
                        Align(lipgloss.Left)

	PromptStyle = lipgloss.
			NewStyle().
			Foreground(accentColor).
			Align(lipgloss.Left).
			PaddingLeft(2)
	Text = lipgloss.
		NewStyle().
		Foreground(primaryColor)

        headerStyle = lipgloss.NewStyle().PaddingBottom(4)

	// item styles
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	titleStyle        = lipgloss.NewStyle().Foreground(primaryColor)
	descStyle         = lipgloss.NewStyle().Foreground(secondaryColor)
	selectedItemStyle = lipgloss.NewStyle().Foreground(accentColor)

	// description styles
	PrimaryStyle    = lipgloss.NewStyle().Foreground(primaryColor)
	SecondaryStyle  = lipgloss.NewStyle().Foreground(secondaryColor)
	accentStyle    = lipgloss.NewStyle().Foreground(accentColor)
	backgroundStyle = lipgloss.NewStyle().Background(backgroundColor)
)
