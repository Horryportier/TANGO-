package main

import (
    "fmt"
	jisho "github.com/Horryportier/go-jisho"
)
func PrintHelp() {
    text := "help"
    fmt.Printf("%s\n", text)
}

func PrintErr(err error, style bool) {
    if style {
        fmt.Println(errorStyle.Render(fmt.Sprint(err)))
        return
    }
    fmt.Println(fmt.Sprint(err))
}

/// pto = print to term
func PrintWord(word jisho.WordData, pto bool) string {
    var text Text
    text = TextFrom([]Line{
        LineFrom([]Span{
            SpanFrom(word.Data[0].Slug, japaneseStyle),

            SpanFrom("(", dimStyle),
            SpanFrom(word.Data[0].Japanese[0].Reading, japaneseStyle),
            SpanFrom(")", dimStyle),

            SpanFrom("=>", arrowStyle),
            SpanFrom(word.FirstTransation(), textStyle),
    }),
    LineFrom([]Span{
        SpanFrom(word.Data[0].Jlpt[0], dimStyle),
        SpanFrom("|", textStyle),
        SpanFrom(word.Data[0].Tags[0], dimStyle),
    }),
    },)
    
    if pto {
        fmt.Printf("%s\n", text.Render(ENABLE_STYLE))
    }
    return text.Render(true)
}
