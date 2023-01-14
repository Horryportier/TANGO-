package utils

import (
	"log"
	"regexp"

	"github.com/atotto/clipboard"
)

func Try(args ...any) interface{} {
	var filter []any
	for _, arg := range args {
		if _, ok := arg.(error); ok {
			if arg != nil {
				log.Fatal(arg)
			}
		} else {
			filter = append(filter, arg)
		}

	}
	if len(filter) == 1 {
		return filter[0]
	}
	return filter
}

func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func CopyToClipord(content string) {
	clipboard.WriteAll(content)
}

func RemoveTermColor(in string) string {
	const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"
	reg := regexp.MustCompile(ansi)
	res := reg.ReplaceAllString(in, "")
	return res
}
