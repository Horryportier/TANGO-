package app

import (
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type SearchModel struct {
	Input     textinput.Model
	Paginator paginator.Model
}

func SearchInit() SearchModel {
	// input
	input := textinput.New()
	input.CharLimit = 64
	input.Focus()
	input.Placeholder = "TANGO"
	//paginator
	pagin := paginator.New()

	return SearchModel{Input: input, Paginator: pagin}
}

func SearchUpdate(m model, msg tea.Msg) (model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "esc":
			return m, tea.Quit
		}
	case errMsg:
		m.Error = msg
		m.state = Err
		return m, nil
	}
	m.SearchModel.Input, cmd = m.SearchModel.Input.Update(msg)
	return m, cmd
}

func SearchingUpdate(m model, msg tea.Msg) (model, tea.Cmd) {
	var cmd tea.Cmd
	return m, cmd
}

func SearchView(m model) string {
	return m.SearchModel.Input.View()
}

func SearchingView(m model) string {
	return m.SearchModel.Paginator.View()
}
