package main

import (
	"fmt"
	"os"

	jisho "github.com/Horryportier/go-jisho"
)
var (
    args = os.Args[1:]
)

func main() {
    if len(args) == 0 {
                PrintErr(fmt.Errorf("%s", "no args where passed"), ENABLE_STYLE)               
                os.Exit(1)
            }
    for _, arg := range args {
        switch arg {
        case "-h":
            PrintHelp()

        default: 
            var word jisho.WordData
            word.Get(arg)
            PrintWord(word, true)
        }
    }
}
