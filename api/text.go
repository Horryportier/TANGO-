package api

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Text  []Line
type Line  []Span
type Span  struct{ string; lipgloss.Style}

/// takes string and style returns Span
func  SpanFrom(text string, style lipgloss.Style) Span {

      return Span{string: text, Style: style}
}

func (s Span) Render(styled bool) string  {
    if styled {
        return s.Style.Render(s.string)
    }
    return s.string
}

/// takes ([]Span, Span, string) and optional style returns Line
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
    case []string:
        var spans []Span
        for _,s := range t.([]string) {
            spans = append(spans, SpanFrom(s, style))
        }
        return spans
    default: 
        return []Span{SpanFrom(fmt.Sprintf("type (%T) not suported! for Line", t), ErrorStyle)}
    }
}

func (l Line) Render(styled ...bool) string{
    var full []string 
    enable_style := true
    for _,s := range styled {
        if !s {
            enable_style = false
        }
    }
    for _,line := range l {
        full = append(full, line.Render(enable_style))
    }
    return strings.Join(full, " ")
}

/// takes ([]Span, Span, string, Line, []Line) and returns Text
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
    case []string:
        return []Line{LineFrom(t, style)}
    case Line:
        return []Line{t.(Line)}
    case []Line:
        return t.([]Line)
    default: 
        return []Line{LineFrom(fmt.Sprintf("type (%T) not suported! for Text", t), ErrorStyle)}
    }
}

func (t Text) Render(styled ...bool) string{
    enable_style := true
    for _,s := range styled {
        if !s {
            enable_style = false
        }
    }
    var full []string 
    for _,text := range t {
        full = append(full, text.Render(enable_style))
    }
    return strings.Join(full, "\n")
}


