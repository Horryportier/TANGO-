package app

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type (
	State  int
	errMsg error
)

const (
	Search State = iota
	List
	Err
	Item
)

var (
	termWidth      int
	GlamourRendere *glamour.TermRenderer
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
	GlamourRendere, _ = glamour.NewTermRenderer(
		glamour.WithWordWrap(40),
		glamour.WithAutoStyle(),
	)
	SetStyle()

	return model{state: Search,
		SearchModel: SearchInit(),
		ListModel:   ListInit(),
		help:        help.New(),
		keys:        keys,
	}
}

func (m model) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, tea.EnterAltScreen)
	return tea.Batch(cmds...)
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
		)

		res = lipgloss.JoinVertical(lipgloss.Left,
			headerStyle.Render(res),
			ListView(m))

		help := RenderFullHelp(m)

		res = lipgloss.JoinVertical(lipgloss.Left, res, help)

		return res
	}()

	return view
}

func Start() error {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		return err
	}
	return nil
}
