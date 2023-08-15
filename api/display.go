package api

import (
	"fmt"

	jisho "github.com/Horryportier/go-jisho"
	. "github.com/Horryportier/lipgloss-text"
)
func PrintHelp() {
    keys := [][]string{
        {"j"},{"k"}, {"Enter"}, {"Tab"}, {"cntl+c", "esc"},
    }
    desc := []string{"down", "up", "search", "switch inptut/list", "exit"}
    var lines Line 

    for i, v := range keys {
        key := LineFrom(v, DimStyle)
        desc := LineFrom(desc[i], DefStyle)
    
        lines.Append(key, desc)
    } 
    fmt.Printf("%s\n", lines.Render(Styled(ENABLE_STYLE)))
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
            SpanFrom(ReturnFirstOrDef(ReturnFirstOrDef(data.Senses).EnglishDefinitions), DefStyle),
    }),
    LineFrom(metadata, DimStyle),
    },)
    
    if pto {
        fmt.Printf("%s\n", text.Render(Styled(ENABLE_STYLE)))
    }
    return text.Render(Styled(ENABLE_STYLE))
}

