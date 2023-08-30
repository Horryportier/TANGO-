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

func PrintWord(data jisho.Data) string {
    var text Text
    
    text = TextFrom([]Line{
        LineFrom([]Span{
            SpanFrom(data.Slug, JapaneseStyle),

            SpanFrom("(", DimStyle),
            SpanFrom(ReturnFirstOrDef(data.Japanese).Reading, JapaneseStyle),
            SpanFrom(")", DimStyle),

            SpanFrom("=>", AcentStyle),
            SpanFrom(ReturnFirstOrDef(ReturnFirstOrDef(data.Senses).EnglishDefinitions), DefStyle),
    })})

    var metadata []string  = GetMetadata(data)   
    if len(metadata) > 1 {
        text.Append(TextFrom(LineFrom(metadata, AcentStyle)))
    } 

    return text.Render(Styled(ENABLE_STYLE))
}


func PrintWordFull(data jisho.Data) string {
    var text Text
    text = TextFrom([]Line{
        LineFrom(data.Slug, JapaneseStyle),
    })


    var jp Line
    var eDef string
    for i,v := range data.Japanese {
        sen := ReturnFirstOrDefSlice(data.Senses, i)
        ed := ReturnFirstOrDef(sen).EnglishDefinitions

        tjp := LineFrom(LineFrom([]Span{
            SpanFrom(v.Word, JapaneseStyle), 
            SpanFrom(":", DimStyle), 
            SpanFrom(v.Reading, JapaneseStyle),
        }).Render(Delimiter("")))

        jp.Append(tjp)

        eDef = LineFrom(ReturnFirstOrDefSlice(ed, SliceOfLength(ed)), DefStyle).Render(Delimiter("\n"))
    }


    text.Append(TextFrom(jp.Render(Delimiter("\n"))))
    text.Append(TextFrom(eDef))

    var metadata []string  = GetMetadata(data)   
    if len(metadata) > 1 {
        text.Append(TextFrom(LineFrom(metadata, AcentStyle)))
    }

    return text.Render(Styled(ENABLE_STYLE))
}

func GetMetadata(data jisho.Data) []string {
        s := ReturnFirstOrDefSlice(data.Jlpt, SliceOfLength(data.Jlpt))
        s = append(s, "|")
        for _,v := range ReturnFirstOrDefSlice(data.Tags, SliceOfLength(data.Tags)) {
            s = append(s, v)
        }
        s = ClearEmptyStr(s)
      return   s
}
