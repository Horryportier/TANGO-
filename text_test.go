package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
    api "github.com/Horryportier/tango/api"
)

func TestSpan(t *testing.T) {
        var span Span
        span = SpanFrom("test", api.defStyle)
        if span.string !=  "test" {
            t.Errorf("span string should be [test] not: %s", span.string)
        }
}

func TestLine(t *testing.T) {
    var line Line 
    var span Span
    spans := []Span{span.From("1", arrowStyle), span.From("2", textStyle)}
    line.From(spans)
    if cmp.Equal(line, spans) {
            t.Errorf("line %v is not equal to %v", line, spans)
    }

}

