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
	switch m.state {
	case Search:
		return SearchUpdate(m, msg)
	case Searching:
		return SearchingUpdate(m, msg)
	case List:
		return ListUpdate(m, msg)
	case Err:
		return m, tea.Quit
	}
	return m, tea.Quit
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
		res = lipgloss.JoinVertical(lipgloss.Left, res, ListView(m))

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
