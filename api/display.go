package api

import (
	"fmt"

	jisho "github.com/Horryportier/go-jisho"
	. "github.com/Horryportier/lipgloss-text"
)
func PrintHelp(styled bool) {
    topText := TextFrom([]Line{
        LineFrom("---TANGO---", AcentStyle),
        LineFrom("https://github.com/Horryportier/TANGO-", JapaneseStyle),
        LineFrom("tango is an simple Japanese/Endlish dictionry", DefStyle),
        LineFrom("that uses jisho.org api.", DefStyle),
        LineFrom("-----------", AcentStyle),
        LineFrom("--Args--", JapaneseStyle),
        })
    bottomText := TextFrom([]Line{
        LineFrom(""),
        LineFrom("None    will lunch interactive tui"),
        LineFrom(""),
        LineFrom("[word]  will return translatino witohut tui"),
        LineFrom("        you can pass kanji/hiragana/katakana/romaji"),
        LineFrom(""),
        LineFrom("-h      for help"),
        LineFrom(""),
        LineFrom("-r      will disable styling"),
    })

    var text Text = TextFrom("", DefStyle)
    text.Append(topText, bottomText)
    println(text.Render(Styled(styled)))
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

