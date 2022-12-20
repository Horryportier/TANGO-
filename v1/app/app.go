package app


import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type(
        State int
        errMsg error
)

const (
	Search State = iota
	Searching
	List
        Err
)

type model struct {
	state     State
	SearchModel SearchModel
	ListModel ListModel
        Error error
}

func initialModel() model {
	return model{state: Search,
        SearchModel: SearchInit(),
        ListModel: ListInit(),
        }
}

func (m model) Init() tea.Cmd {
	return textinput.Blink 
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
	switch m.state {
	case Search:
		return SearchView(m)
	case Searching:
		return SearchingView(m)
	case List:
		return ListView(m)
        case Err:
                return fmt.Sprintf("Error: %v",m.Error)
	}
        return "OPPPS :("
}

func Start() error {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		return err
	}
	return nil
}
