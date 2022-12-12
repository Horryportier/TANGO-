package main

import (
	"fmt"
	"log"

        jisho "github.com/Horryportier/go-jisho"
)

func main() {
        word, err := jisho.Search("bird")
        if err != nil {
                log.Fatal(err)
        }
        fmt.Printf("data: %v \n", word)

        fmt.Printf("JAPANES: %v\n", word.TransJapan(0))
        fmt.Printf("ENG: %v\n", word.EngDefinition(0))
}
