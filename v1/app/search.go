package app

import (
	"github.com/Horryportier/tango/v1/utils"

	jisho "github.com/Horryportier/go-jisho"
	"github.com/charmbracelet/bubbles/key"
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
		switch {
                case key.Matches(msg, keys.Quit):
			return m, tea.Quit
                case key.Matches(msg, keys.Enter):
			val := m.SearchModel.Input.Value()
			var err error

			items := SearchForPhrase(val)
			cmd = m.ListModel.List.SetItems(items)

			m.ListModel.List.SetHeight(len(items)*2)
			if err != nil {
				return m, tea.Quit
			}
			m.state = List
			return m, cmd
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

func SearchForPhrase(word string) []list.Item {
	Word, _ := jisho.Search(word)

	numOfEntries := Word.Len()
	entries := Word.GetEntries(utils.MakeRange(0, numOfEntries-1)...)

	items := make([]list.Item, numOfEntries)
	for i := 0; i < numOfEntries; i++ {
		items[i] = ItemGenerator(entries[i])
	}
	return items
}
