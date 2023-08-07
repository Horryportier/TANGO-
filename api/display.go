package api

import (
    "fmt"
	jisho "github.com/Horryportier/go-jisho"
)
func PrintHelp() {
    text := "help"
    fmt.Printf("%s\n", text)
}

func PrintErr(err error) {
    if ENABLE_STYLE {
        fmt.Println(ErrorStyle.Render(fmt.Sprint(err)))
        return
    }
    fmt.Println(fmt.Sprint(err))
}

/// pto = print to term
func PrintWord(word jisho.WordData, pto bool) string {
    var text Text
    text = TextFrom([]Line{
        LineFrom([]Span{
            SpanFrom(word.Data[0].Slug, JapaneseStyle),

            SpanFrom("(", DimStyle),
            SpanFrom(word.Data[0].Japanese[0].Reading, JapaneseStyle),
            SpanFrom(")", DimStyle),

            SpanFrom("=>", ArrowStyle),
            SpanFrom(word.FirstTransation(), TextStyle),
    }),
    LineFrom([]Span{
        SpanFrom(word.Data[0].Jlpt[0], DimStyle),
        SpanFrom("|", TextStyle),
        SpanFrom(word.Data[0].Tags[0], DimStyle),
    }),
    },)
    
    if pto {
        fmt.Printf("%s\n", text.Render(ENABLE_STYLE))
    }
    return text.Render(true)
}
