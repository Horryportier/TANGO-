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

	for _, arg := range args {
		switch arg {
		case "-h":
			utils.PrintHelp()
		default:
			var word jisho.Word
			res, err := jisho.Search(arg)
			if err != nil {
				log.Fatal(err)
			}
			word, err = word.Parse(res)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%s: %s %s", arg, word.Data[0].Slug, word.Data[0].Senses[0].EnglishDefinitions[0])
		}
	}

}
