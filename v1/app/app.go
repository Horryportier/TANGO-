package app

import (
	"strings"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	State  int
	errMsg error
)

const (
	Search State = iota
	Searching
	List
	Err
	Item
)

var (
	termWidth int
)

type model struct {
	keys        keyMap
	help        help.Model
	state       State
	SearchModel SearchModel
	ListModel   ListModel
	Error       error
}

func initialModel() model {
	return model{state: Search,
		SearchModel: SearchInit(),
		ListModel:   ListInit(),
		help:        help.New(),
		keys:        keys,
	}
}

func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		termWidth = msg.Width
		return m, nil
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Help):
			m.help.ShowAll = !m.help.ShowAll
			return m, nil
		}
	}

	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch m.state {
	case Search:
		m, cmd = SearchUpdate(m, msg)
		cmds = append(cmds, cmd)
	case Searching:
		m, cmd = SearchingUpdate(m, msg)
		cmds = append(cmds, cmd)
	case List:
		m, cmd = ListUpdate(m, msg)
		cmds = append(cmds, cmd)
	case Err:
		cmd = tea.Quit
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var view string

	view = func() string {
		res := lipgloss.JoinHorizontal(
			lipgloss.Left,
			SearchView(m),
			" | ",
			SearchingView(m),
		)

		res = lipgloss.JoinVertical(lipgloss.Left,
			headerStyle.Render(res),
			ListView(m))

		help := func(m model) string {
                        var str strings.Builder
                        str.WriteString(m.help.View(m.keys))
			return str.String()
		}

		res = lipgloss.JoinVertical(lipgloss.Left, res, help(m))

		return res
	}()
	return view
}

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


func Start() error {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		return err
	}
	return nil
}
