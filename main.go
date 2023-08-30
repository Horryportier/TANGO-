package main

import (
	"fmt"
	"os"

	jisho "github.com/Horryportier/go-jisho"
	"github.com/Horryportier/tango/api"
	tui "github.com/Horryportier/tango/tui"
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
        case "-r":
            api.ENABLE_STYLE = false
        case "-h":
            api.PrintHelp(api.ENABLE_STYLE)
        default: 
            var word jisho.WordData
            word = api.DefWord()
            if err := word.Get(arg); err != nil {
                api.PrintErr(err)
            }
            fmt.Println(api.PrintWord(api.ReturnFirstOrDef(word.Data)))
        }
    }
}

