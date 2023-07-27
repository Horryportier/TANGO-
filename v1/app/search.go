package app

import (
	"github.com/Horryportier/tango/v1/utils"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type SearchModel struct {
	Input   textinput.Model
	Spinner spinner.Model
}

func SearchInit() SearchModel {
	// input
	input := textinput.New()
	input.CharLimit = 64
	input.Focus()
	input.Placeholder = "TANGO"
	input.PromptStyle = PromptStyle.Copy()
	input.TextStyle = Text.Copy()
	//spinner
	s := spinner.New(
		spinner.WithSpinner(spinner.Dot),
		spinner.WithStyle(accentStyle),
	)

	return SearchModel{Input: input, Spinner: s}
}

func SearchUpdate(m model, msg tea.Msg) (model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SearchModel.Input, cmd = m.SearchModel.Input.Update(msg)
		return m, cmd

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, keys.Enter):
			val := m.SearchModel.Input.Value()

			ls := make(chan []list.Item)
			go SearchForPhrase(val, ls)
			items, ok := <-ls
			if ok {
				cmd = m.ListModel.List.SetItems(items)
				m.ListModel.List.SetHeight(len(items) * 2)
				m.state = List
				return m, cmd
			}
		}
	case errMsg:
		m.Error = msg
		m.state = Err
		return m, nil
	}
	m.SearchModel.Input, cmd = m.SearchModel.Input.Update(msg)
	cmds = append(cmds, cmd)

	if m.SearchModel.Input.Focused() {
		m.SearchModel.Input.SetCursorMode(textinput.CursorBlink)
	}

	return m, tea.Batch(cmds...)
}

func SearchView(m model) string {
	return m.SearchModel.Input.View()
}

func SearchForPhrase(word string, ls chan []list.Item) error {
    err := Word.Get(word)
	if err != nil {
		return err
	}

	numOfEntries := Word.Len()
	entries, err := Word.GetEntries(utils.MakeRange(0, numOfEntries-1)...)
	if err != nil {
		return err
	}

	items := make([]list.Item, numOfEntries)
	for i := 0; i < numOfEntries; i++ {
		items[i] = ItemGenerator(entries[i])
	}
	ls <- items
	close(ls)
    return nil
}
