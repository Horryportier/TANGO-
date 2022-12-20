package app

import (
	utils "src/tango/v1/utils"

	tea "github.com/charmbracelet/bubbletea"
)

type State int

const (
	Search State = iota
	Searching
	List
)

type model struct {
	state     State
	searchBar SearchModel
	results   ListModel
}

func initialModel() model {
	return model{}
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
	}
	return m, tea.Quit
}

func (m model) View() string {
	switch m.state {
	case Search:
		return SearchView(m)
	case Searching:
		return SearchingView(m)
	case List:
		return ListView(m)
	}
	return "POG"
}

func Start() error {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		return err
	}
	return nil
}
