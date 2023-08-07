package api

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Text  []Line
type Line  []Span
type Span  struct{ string; lipgloss.Style}

func  SpanFrom(text string, style lipgloss.Style) Span {
      return Span{string: text, Style: style}
}

func (s Span) Render(styled bool) string  {
    if styled {
        return s.Style.Render(s.string)
    }
    return s.string
}

func LineFrom(t interface{}, options ...interface{}) Line {
    var style lipgloss.Style = DefStyle
    for _,o := range options {
        switch o.(type) {
        case lipgloss.Style:
            style = o.(lipgloss.Style)
        }
    }

    switch t.(type) {
    case []Span:
        return t.([]Span)
    case Span:
        return []Span{t.(Span)}
    case string:
        return []Span{SpanFrom(t.(string), style)}
    default: 
        return []Span{SpanFrom(fmt.Sprintf("type (%T) not suported! for Line", t), ErrorStyle)}
    }
}

func (l Line) Render(styled bool) string{
    var full []string 

    for _,line := range l {
        full = append(full, line.Render(styled))
    }
    return strings.Join(full, " ")
}
func TextFrom(t interface{}, options ...interface{}) Text {
    var style lipgloss.Style = DefStyle
    for _,o := range options {
        switch o.(type) {
        case lipgloss.Style:
            style = o.(lipgloss.Style)
        }
    }
    switch t.(type) {
    case []Span:
        return []Line{LineFrom(t)}
    case Span:
        return []Line{LineFrom(t)}
    case string:
        return []Line{LineFrom(t, style)}
    case Line:
        return []Line{t.(Line)}
    case []Line:
        return t.([]Line)
    default: 
        return []Line{LineFrom(fmt.Sprintf("type (%T) not suported! for Text", t), ErrorStyle)}
    }
}

func (t Text) Render(styled bool) string{
    var full []string 
    for _,text := range t {
        full = append(full, text.Render(styled))
    }
    return strings.Join(full, "\n")
}


