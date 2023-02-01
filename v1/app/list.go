package app

import (
	jisho "github.com/Horryportier/go-jisho"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	utils "github.com/Horryportier/tango/v1/utils"
)

var (
	Word jisho.Word
)

type ListModel struct {
	List list.Model
}

func ListInit() ListModel {
	numOfEntries := Word.Len()
	entries := Word.GetEntries(utils.MakeRange(0, numOfEntries-1)...)

	items := make([]list.Item, numOfEntries)
	for i := 0; i < numOfEntries; i++ {

		items[i] = ItemGenerator(entries[i])
	}

	wordList := list.New(items, itemDelegate{}, 0, len(items))

	wordList.SetShowTitle(true)
	wordList.Title = "Results"
	//style
	wordList.Styles.Title = PrimaryStyle.Copy()
	wordList.Styles.Title.Blink(true)
	wordList.Styles.FilterPrompt = PrimaryStyle.Copy()
	wordList.Styles.FilterCursor = PrimaryStyle.Copy()
	wordList.Styles.DefaultFilterCharacterMatch = PrimaryStyle.Copy()
	wordList.SetShowHelp(false)

	wordList.FilterInput.BackgroundStyle = PrimaryStyle.Copy()

	return ListModel{List: wordList}
}

func ListUpdate(m model, msg tea.Msg) (model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.ListModel.List.SetSize(msg.Width-h, msg.Height-v)

		newListModel, cmd := m.ListModel.List.Update(msg)
		m.ListModel.List = newListModel
                return m, cmd
                

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, keys.Tab):
			m.state = Search
			m.SearchModel.Input.Reset()
			return m, nil

		case key.Matches(msg, keys.Enter):
			return m, nil
		case key.Matches(msg, keys.Clip):
			i := func(m model) jisho.Data {
				listItem := m.ListModel.List.SelectedItem()

				i, ok := listItem.(item)

				if !ok {
					return jisho.Data{}
				}

				return i.data
			}
			content := DetialsView(i(m), m.ListModel.List.Width(), true)

			utils.CopyToClipord(content)
			return m, nil
		}
	case errMsg:
		m.Error = msg
		m.state = Err
		return m, nil
	}

	newListModel, cmd := m.ListModel.List.Update(msg)
	m.ListModel.List = newListModel
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func ListView(m model) string {
	var desc string

	listItem := m.ListModel.List.SelectedItem()

	i, ok := listItem.(item)

	if !ok {
		return ""
	}

	desc = DetialsView(i.data, m.ListModel.List.Width(), false)
	return appStyle.Render(
		lipgloss.JoinHorizontal(lipgloss.Left, m.ListModel.List.View(), desc))
}
