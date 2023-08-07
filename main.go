package main

import (
	"os"

	jisho "github.com/Horryportier/go-jisho"
    tui "github.com/Horryportier/tango/tui"
    "github.com/Horryportier/tango/api"
)
var (
    args = os.Args[1:]
)

func main() {
    if len(args) == 0 {
        if err := tui.Run(); err != nil{
            api.PrintErr(err)
            os.Exit(1)
        }
        os.Exit(0)
    }
    for _, arg := range args {
        switch arg {
        case "-h":
            api.PrintHelp()

        default: 
            var word jisho.WordData
            word.Get(arg)
            api.PrintWord(word, true)
        }
    }
}
