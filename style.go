package main

import "github.com/charmbracelet/lipgloss"

var (
    defStyle = lipgloss.NewStyle()
    textStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#AAAAAA"))
    errorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#AA2222"))
    arrowStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#4A00AB"))
    japaneseStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#7AAA00"))
    dimStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#333333")).Faint(true)
)
