package app

import (
	"fmt"
	"strings"

	jisho "github.com/Horryportier/go-jisho"
        "github.com/charmbracelet/glamour"
)

func DetialsView(data jisho.Data, listWidth int, noStyle bool) string {
	var str strings.Builder

	// japanes headear
	str.WriteString(fmt.Sprintf("%s %v %s %v",
		"#",
		data.Slug,
		"source: ",
		strings.Join(data.Tags, " , ")))
	str.WriteRune('\n')
	str.WriteString("___")
	str.WriteRune('\n')

	// jlpt
	str.WriteString(fmt.Sprint(strings.Join(data.Jlpt, " | ")))
	str.WriteRune('\n')
	str.WriteRune('\n')

	// japanes readings
	jp := func(data jisho.Data) string {
		var s strings.Builder
		s.WriteString(
			fmt.Sprintf("%s %s",
				"##",
				"JP/Reading"))
		s.WriteRune('\n')
		for i, val := range data.Japanese {
			s.WriteString(fmt.Sprintf("%v. %s [%s]",
				i+1,
				val.Word, val.Reading))
			s.WriteRune('\n')
		}
		return s.String()
	}
	str.WriteString(jp(data))
	str.WriteRune('\n')
	str.WriteString("___")
	str.WriteRune('\n')

	// eng def
	eng := func(data jisho.Data) string {
		var s strings.Builder
		s.WriteString(
			fmt.Sprintf("%s %s",
				"##",
				"ENG definition"))
		s.WriteRune('\n')
		for i, val := range data.Senses {
			s.WriteString(fmt.Sprintf("%v. %s [%s]()",
				i+1,
				strings.Join(val.EnglishDefinitions, ","),
				strings.Join(val.PartsOfSpeech, ",")))
			s.WriteRune('\n')
		}
		return s.String()
	}
	str.WriteString(eng(data))
	str.WriteRune('\n')
	str.WriteString("___")

	if noStyle {
		return appStyle.Render(str.String())
	}
        
        out, _ := glamour.Render(str.String(), "/home/horryportier/Documents/TANGO-/red_green.json")


	return appStyle.Render(out)
}
