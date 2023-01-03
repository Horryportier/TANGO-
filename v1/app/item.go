package app

import (
	"fmt"
	"io"
	"strings"

	jisho "github.com/Horryportier/go-jisho"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type item struct {
	title       string
	description string
	data        jisho.Data
}

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {

	fn := func(data jisho.Data) string { return "" }

	i, ok := listItem.(item)
	if !ok {
		return
	}

	data := i.data

	if index == m.Index() {
		fn = func(data jisho.Data) string {
			var str strings.Builder

			str.WriteString(
				fmt.Sprintf("%s %s -> %s", selectedItemStyle.Render("=>"),
					titleStyle.Render(data.Slug),
					descStyle.Render(data.Senses[0].EnglishDefinitions[0])))

			return str.String()
		}
	} else {
		fn = func(data jisho.Data) string {
			var str strings.Builder

			str.WriteString(
				fmt.Sprintf("%s %s -> %s", selectedItemStyle.Render("# "),
					titleStyle.Render(data.Slug),
					descStyle.Render(data.Senses[0].EnglishDefinitions[0])))

			return str.String()
		}
	}

	fmt.Fprint(w, fn(data))
}

func (i item) FilterValue() string {
	return ""
}

func ItemGenerator(data jisho.Data) item {
	return item{data: data}
}
