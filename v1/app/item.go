package app

import (
	"strings"

	jisho "github.com/Horryportier/go-jisho"
	"github.com/charmbracelet/lipgloss"
)

type item struct {
	title       string
	description string
	data        jisho.Data
}

func (i item) Title() string {
	var msg strings.Builder

	msg.WriteString(addStyle("=> ", i.data.Slug, accentStyle1, SecondaryStyle))

	i.title = msg.String()
	return i.title
}

func (i item) Description() string {
	var msg strings.Builder

	msg.WriteString(addStyle("Translation: ",
		i.data.Senses[0].EnglishDefinitions[0],
		PrimaryStyle,
		SecondaryStyle))

	msg.WriteRune('\n')

	msg.WriteString(addStyle("Part of speach: ",
		i.data.Senses[0].PartsOfSpeech[0],
		PrimaryStyle,
		SecondaryStyle))

	i.description = msg.String()
	return i.description
}

func (i item) AdditionalData() string {
        return ""
}

func (i item) FilterValue() string { return "LIGMA" }

func ItemGenerator(data jisho.Data) item {
	return item{data: data}
}

func addStyle(title, desc string, titleStyle, descStyle lipgloss.Style) string {
	var msg strings.Builder
	msg.WriteString(titleStyle.Render(title))
	msg.WriteString(descStyle.Render(desc))
	return msg.String()
}
