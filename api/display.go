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
    
    data := ReturnFirstOrDef(word.Data)
    

    text = TextFrom([]Line{
        LineFrom([]Span{
            SpanFrom(data.Slug, JapaneseStyle),

            SpanFrom("(", DimStyle),
            SpanFrom(ReturnFirstOrDef(data.Japanese).Reading, JapaneseStyle),
            SpanFrom(")", DimStyle),

            SpanFrom("=>", AcentStyle),
            SpanFrom(word.FirstTransation(), TextStyle),
    }),
    LineFrom([]Span{
        SpanFrom(ReturnFirstOrDef(data.Jlpt), DimStyle),
        SpanFrom("|", TextStyle),
        SpanFrom(ReturnFirstOrDef(data.Tags), DimStyle),
    }),
    },)
    
    if pto {
        fmt.Printf("%s\n", text.Render(ENABLE_STYLE))
    }
    return text.Render(true)
}
