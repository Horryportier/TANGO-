package utils

import (
	"fmt"
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

func PrintHelp() {
        fmt.Println("Tango is an simple English->Japanese cli dictionary.")
        fmt.Println("   this app uses jisho.org open api")
        fmt.Println("   Kanji support doesn't work use romaji or hiragana/katakana")
        fmt.Println("---COMMANDS--")
        fmt.Println("")
        fmt.Println("NONE                   opens tui")
        fmt.Println("-h                     print help")
        fmt.Println("-r                     print raw json response")
        fmt.Println("-r                     print details view")
        fmt.Println("たんご/タンゴ/tango    will search for the word")
        fmt.Println("-c                     will copy result to clipboard")
}
