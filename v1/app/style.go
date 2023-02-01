package app

import (
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

var (
	// theme
	primaryColor   lipgloss.Color
	secondaryColor lipgloss.Color
	inactiveColor  lipgloss.Color
	accentColor    lipgloss.Color

	appStyle lipgloss.Style

	PromptStyle lipgloss.Style
	Text        lipgloss.Style

	headerStyle lipgloss.Style

	// item styles
	itemStyle         lipgloss.Style
	titleStyle        lipgloss.Style
	descStyle         lipgloss.Style
	selectedItemStyle lipgloss.Style

	// description styles
	PrimaryStyle   lipgloss.Style
	SecondaryStyle lipgloss.Style
	accentStyle    lipgloss.Style
	inactiveStyle  lipgloss.Style
)

func SetStyle() {
	if termenv.HasDarkBackground() {
		primaryColor = lipgloss.Color(*glamour.DarkStyleConfig.Document.Color)
		secondaryColor = lipgloss.Color(*glamour.DarkStyleConfig.Heading.Color)
		inactiveColor = lipgloss.Color(*glamour.DarkStyleConfig.H1.Color)
		accentColor = lipgloss.Color(*glamour.DarkStyleConfig.H1.BackgroundColor)
	} else {
		primaryColor = lipgloss.Color(*glamour.LightStyleConfig.Document.Color)
		secondaryColor = lipgloss.Color(*glamour.LightStyleConfig.Heading.Color)
		inactiveColor = lipgloss.Color(*glamour.LightStyleConfig.H1.Color)
		accentColor = lipgloss.Color(*glamour.LightStyleConfig.H1.BackgroundColor)
	}

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
	itemStyle = lipgloss.NewStyle().PaddingLeft(4)
	titleStyle = lipgloss.NewStyle().Foreground(primaryColor)
	descStyle = lipgloss.NewStyle().Foreground(secondaryColor)
	selectedItemStyle = lipgloss.NewStyle().Foreground(accentColor)

	// description styles
	PrimaryStyle = lipgloss.NewStyle().Foreground(primaryColor)
	SecondaryStyle = lipgloss.NewStyle().Foreground(secondaryColor)
	accentStyle = lipgloss.NewStyle().Foreground(accentColor)
	inactiveStyle = lipgloss.NewStyle().Foreground(inactiveColor)

}
