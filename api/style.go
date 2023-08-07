package api

import "github.com/charmbracelet/lipgloss"

var (
    DefStyle = lipgloss.NewStyle()
    TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#AAAAAA"))
    ErrorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#AA2222"))
    ArrowStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#4A00AB"))
    JapaneseStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#7AAA00"))
    DimStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#333333")).Faint(true)
)
