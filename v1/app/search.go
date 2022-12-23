package app

import (
	"src/tango/v1/utils"

	jisho "github.com/Horryportier/go-jisho"
	"github.com/charmbracelet/bubbles/list"
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
	input.PromptStyle = PromptStyle.Copy()
	input.TextStyle = Text.Copy()
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
		case "enter":
                        val := m.SearchModel.Input.Value()
                        var err error
                        m.ListModel.List, err = SearchForPhrase(val, m.ListModel.List)
                        if err != nil {
                                return m, tea.Quit
                        }
                        m.state = List
                        return m, nil
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

func SearchForPhrase(word string, listItems list.Model) (list.Model, error) {
        Word, err := jisho.Search(word)

	numOfEntries := Word.Len()
	entries := Word.GetEntries(utils.MakeRange(0, numOfEntries-1)...)

	items := make([]list.Item, numOfEntries)
	for i := 0; i < numOfEntries; i++ {
		items[i] = ItemGenerator(entries[i])
	}

        newList := list.New(items, list.NewDefaultDelegate(), 1, 0)
        return newList,  err
}
