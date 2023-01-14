package app

import (
        "github.com/charmbracelet/bubbles/key"
        tea "github.com/charmbracelet/bubbletea"
)


type keyMap struct {
	Help  key.Binding
	Quit  key.Binding
	Enter key.Binding
	Tab   key.Binding
	Clip  key.Binding
	Up    key.Binding
	Down  key.Binding
}

var keys = keyMap{
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "search for word (when in input mode)"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithKeys("esc"),
		key.WithHelp("q/esc", "Exit"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "more help"),
	),
	Tab: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "focus search bar"),
	),
	Clip: key.NewBinding(
		key.WithKeys("cntl+c"),
		key.WithKeys("y"),
		key.WithHelp("cntl+c/y", "copies currrent description"),
	),
	Up: key.NewBinding(
		key.WithKeys("k"),
		key.WithKeys(tea.KeyUp.String()),
		key.WithHelp("k/↑", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j"),
		key.WithKeys(tea.KeyDown.String()),
		key.WithHelp("j/↓", "down"),
	),
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		{k.Help, k.Quit},
		{k.Enter, k.Tab},
		{k.Clip},
	}
}

