package api

import (
    "fmt"
jisho "github.com/Horryportier/go-jisho"
   tx "github.com/Horryportier/lipgloss-text"
)
func PrintHelp() {
    keys := [][]string{
        {"j"},{"k"}, {"Enter"}, {"Tab"}, {"cntl+c", "esc"},
    }
    desc := []string{"down", "up", "search", "switch inptut/list", "exit"}
    var lines Line 

    for i, v := range keys {
        key := tx.LineFrom(v, DimStyle)
        desc := tx.LineFrom(desc[i], DefStyle)


    } 
    fmt.Printf("%s\n", lines.Render(tx.Styled(ENABLE_STYLE)))
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
    var tx Text
    
    var metadata []string = func () []string {
        s := ReturnFirstOrDefSlice(data.Jlpt, []int{0,1,2,3,4})
        s = append(s, "|")
        s = append(s, ReturnFirstOrDef(data.Tags, 0,1, 2))
        s = ClearEmptyStr(s)
      return   s
    }()
    
    
    tx = TextFrom([]Line{
        LineFrom([]Span{
            SpanFrom(data.Slug, JapaneseStyle),

            SpanFrom("(", DimStyle),
            SpanFrom(ReturnFirstOrDef(data.Japanese).Reading, JapaneseStyle),
            SpanFrom(")", DimStyle),

            SpanFrom("=>", Acentxtyle),
            SpanFrom(ReturnFirstOrDef(ReturnFirstOrDef(data.Senses).EnglishDefinitions), txStyle),
    }),
    LineFrom(metadata, DimStyle),
    },)
    
    if pto {
        fmt.Printf("%s\n", tx.Render(ENABLE_STYLE))
    }
    return tx.Render(ENABLE_STYLE)
}

