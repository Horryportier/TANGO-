package main

import (
	"fmt"
	"log"
	"os"

	jisho "github.com/Horryportier/go-jisho"
	app "github.com/Horryportier/tango/v1/app"
	"github.com/Horryportier/tango/v1/utils"
)

func main() {
	var args = os.Args[1:]

	if len(args) == 0 {
		if err := app.Start(); err != nil {
			log.Fatal(err)
		}
	}

	for i, arg := range args {
		switch arg {
		case "-h":
			utils.PrintHelp()
        case "-r":
            d, err := jisho.Search(args[i+1])
            if err != nil {
                fmt.Printf("Error: %v", err)
                os.Exit(1)
            }
            fmt.Print(string(d))

        case "-d":
            var word jisho.WordData
            err := word.Get(args[i+1])
            if err != nil {
                fmt.Printf("Error: %v", err)
                os.Exit(1)
            }
            f, err := word.First()
            if err != nil {
                fmt.Printf("Error: %v", err)
                os.Exit(1)
            }
            fmt.Print(app.DetialsView(f, 100, true))

		default:
			var word jisho.WordData
            word.Get(arg)

			fmt.Printf("%s: %s %s", arg, word.Data[0].Slug, word.Data[0].Senses[0].EnglishDefinitions[0])
		}
	}

}
