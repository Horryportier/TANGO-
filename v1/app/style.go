package app

import (
	"github.com/charmbracelet/lipgloss"
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

func SetStyle(isCustom bool) {
	if isCustom {
		primaryColor = lipgloss.Color("#00c940")   //41
		secondaryColor = lipgloss.Color("#909dc0") //109
		inactiveColor = lipgloss.Color("#8a8a8a")  //102
		accentColor = lipgloss.Color("#ff2a42")    //197
	} else {
		primaryColor = "#d0d0d0"
		secondaryColor = "#00afff"
		inactiveColor = "#ffff87"
		accentColor = "#5f5fff"
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
