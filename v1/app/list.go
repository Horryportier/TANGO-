package app

import (

	jisho "github.com/Horryportier/go-jisho"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	utils "src/tango/v1/utils"
)

var (
	Word jisho.Word
)

type ListModel struct {
	list list.Model
}

func ListInit() ListModel {
	numOfEntries := Word.Len()
	indexes := utils.MakeRange(0, numOfEntries-1)
	entries := Word.GetEntries(indexes...)
	items := make([]list.Item, numOfEntries)
	for i := 0; i < numOfEntries; i++ {
		items[i] = ItemGenerator(entries[i])
	}

	delegate := list.NewDefaultDelegate()

	wordList := list.New(items, delegate, 0, 0)
	return ListModel{list: wordList}
}

func ListUpdate(m model, msg tea.Msg) (model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.ListModel.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "esc":
			return m, tea.Quit
		case "enter":
		}
	case errMsg:
		m.Error = msg
		m.state = Err
		return m, nil
	}
	m.ListModel.list, cmd = m.ListModel.list.Update(msg)
	return m, cmd
}

func ListView(m model) string {
	return m.ListModel.list.View()
}
