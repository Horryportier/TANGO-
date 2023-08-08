package api

import "github.com/charmbracelet/lipgloss"

var (
    DefStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(THEME[0]))
    TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(THEME[1]))
    ErrorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(THEME[2]))
    AcentStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(THEME[3]))
    JapaneseStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(THEME[4]))
    DimStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(THEME[5]))
    FaintStyle = lipgloss.NewStyle().Faint(true)
)


