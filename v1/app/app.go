package app

import (
	"fmt"
	"strings"

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
	Searching
	List
	Err
	Item
)

var (
	termWidth int
	R         *glamour.TermRenderer
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
	R, _ = glamour.NewTermRenderer(
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
			"  ",
			SearchingView(m),
		)

		res = lipgloss.JoinVertical(lipgloss.Left,
			headerStyle.Render(res),
			ListView(m))

                // don't judge
		help := func(m model) string {
			var str strings.Builder
			if m.help.ShowAll {
				styles := m.help.Styles
				k := keys.FullHelp()
				help := func(k [][]key.Binding) string {
					var r strings.Builder
					for _, val := range k {
						for _, w := range val {
							r.WriteString(
								styles.Ellipsis.Render(
									fmt.Sprintf("%s %s • ",
										w.Help().Key,
										w.Help().Desc,
									),
								),
							)
						}
						r.WriteRune('\n')
					}
					return r.String()
				}

				return help(k)
			}
			str.WriteString(m.help.View(m.keys))
			return str.String()

		}

		res = lipgloss.JoinVertical(lipgloss.Left, res, help(m))

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
