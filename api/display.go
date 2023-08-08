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
func PrintWord(data jisho.Data, pto bool) string {
    var text Text
    
    var metadata []string = func () []string {
        s := ReturnFirstOrDefSlice(data.Jlpt, []int{0,1,2,3,4})
        s = append(s, "|")
        s = append(s, ReturnFirstOrDef(data.Tags, 0,1, 2))
        s = ClearEmptyStr(s)
      return   s
    }()
    
    
    text = TextFrom([]Line{
        LineFrom([]Span{
            SpanFrom(data.Slug, JapaneseStyle),

            SpanFrom("(", DimStyle),
            SpanFrom(ReturnFirstOrDef(data.Japanese).Reading, JapaneseStyle),
            SpanFrom(")", DimStyle),

            SpanFrom("=>", AcentStyle),
            SpanFrom(ReturnFirstOrDef(ReturnFirstOrDef(data.Senses).EnglishDefinitions), TextStyle),
    }),
    LineFrom(metadata, DimStyle),
    },)
    
    if pto {
        fmt.Printf("%s\n", text.Render(ENABLE_STYLE))
    }
    return text.Render(ENABLE_STYLE)
}
