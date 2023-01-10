package app

import (
	"fmt"
	"regexp"
	"strings"

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

			utils.copyToClipord(content)
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

// seperate styling from genereing leyout.
func DetialsView(data jisho.Data, listWidth int, noStyle bool) string {
	var str strings.Builder

	// japanes headear
	str.WriteString(fmt.Sprintf("%s %v %s %v",
		accentStyle.Render("#"),
		PrimaryStyle.Render(data.Slug),
		accentStyle.Render("source: "),
		SecondaryStyle.Render(
			strings.Join(data.Tags, " , "))))
	str.WriteRune('\n')
	str.WriteString("___")
	str.WriteRune('\n')

	// jlpt
	str.WriteString(
		SecondaryStyle.Render(
			fmt.Sprint(strings.Join(data.Jlpt, " | "))))
	str.WriteRune('\n')
	str.WriteRune('\n')

	// japanes readings
	jp := func(data jisho.Data) string {
		var s strings.Builder
		s.WriteString(
			fmt.Sprintf("%s %s", accentStyle.Render("##"),
				PrimaryStyle.Render("JP/Reading")))
		s.WriteRune('\n')
		for i, val := range data.Japanese {
			s.WriteString(fmt.Sprintf("%v. %s [%s]",
				i+1,
				val.Word, val.Reading))
			s.WriteRune('\n')
		}
		return SecondaryStyle.Render(s.String())
	}
	str.WriteString(jp(data))
	str.WriteRune('\n')
	str.WriteString("___")
	str.WriteRune('\n')

	// eng def
	eng := func(data jisho.Data) string {
		var s strings.Builder
		s.WriteString(
			fmt.Sprintf("%s %s", accentStyle.Render("##"),
				PrimaryStyle.Render("ENG definition")))
		s.WriteRune('\n')
		for i, val := range data.Senses {
			s.WriteString(fmt.Sprintf("%v. %s %s",
				i+1,
				strings.Join(val.EnglishDefinitions, ","),
				accentStyle.Render(
					strings.Join(val.PartsOfSpeech, ","))))
			s.WriteRune('\n')
		}
		return SecondaryStyle.Render(s.String())
	}
	str.WriteString(eng(data))
	str.WriteRune('\n')
	str.WriteString("___")

	rd := func(str string) string {
		var s strings.Builder
		var offset int = appStyle.GetPaddingLeft()
		width := (termWidth - listWidth - offset) * 2

		strs := strings.Split(str, "\n")

		for _, val := range strs {
			if len(val) > width {
				a := val[:width]
				b := val[width:]
				s.WriteString(fmt.Sprintf(
					"%s\n%s\n", a,
					accentStyle.Render(b)))
			} else {
				s.WriteString(val)
				s.WriteRune('\n')
			}
		}

		return s.String()
	}

	if noStyle {
		const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"
		reg := regexp.MustCompile(ansi)
		res := reg.ReplaceAllString(str.String(), "")
		return res
	}

	return appStyle.Render(rd(str.String()))
}
