package app

import (
	"fmt"

	jisho "github.com/Horryportier/go-jisho"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	utils "src/tango/v1/utils"
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
	fmt.Print(entries)
	items := make([]list.Item, numOfEntries)
	for i := 0; i < numOfEntries; i++ {
		items[i] = ItemGenerator(entries[i])
	}

	delegate := list.NewDefaultDelegate()
	wordList := list.New(items, delegate, 1, 50)

	wordList.SetShowTitle(true)
	wordList.Title = "Results"
	//style
	wordList.Styles.Title = accentStyle1.Copy()
	wordList.Styles.Title.Blink(true)
	wordList.Styles.FilterPrompt = PrimaryStyle.Copy()
	wordList.Styles.FilterCursor = PrimaryStyle.Copy()
	wordList.Styles.DefaultFilterCharacterMatch = accentStyle1.Copy()

	wordList.FilterInput.BackgroundStyle = PrimaryStyle.Copy()

	return ListModel{List: wordList}
}

func ListUpdate(m model, msg tea.Msg) (model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.ListModel.List.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "esc":
			return m, tea.Quit
		case "enter", "tab":
                        m.state = Search
                        m.SearchModel.Input.Reset()
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
	return appStyle.Render(m.ListModel.List.View())
}
