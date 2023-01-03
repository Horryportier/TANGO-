package app

import (
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

type model struct {
	state       State
	SearchModel SearchModel
	ListModel   ListModel
	Error       error
}

func initialModel() model {
	return model{state: Search,
		SearchModel: SearchInit(),
		ListModel:   ListInit(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
