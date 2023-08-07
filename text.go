package main

import (
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

func LineFrom(spans []Span) Line {
    return spans
}

func (l Line) Render(styled bool) string{
    var full []string 
    for _,line := range l {
        full = append(full, line.Render(styled))
    }
    return strings.Join(full, " ")
}
func TextFrom(lines []Line) Text {
    return lines
}

func (t Text) Render(styled bool) string{
    var full []string 
    for _,text := range t {
        full = append(full, text.Render(styled))
    }
    return strings.Join(full, "\n")
}


