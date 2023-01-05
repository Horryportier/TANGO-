package main

import (
	"log"
	app "github.com/Horryportier/tango/v1/app"
)

func main() {
        if err := app.Start(); err != nil{
                log.Fatal(err)
        }
}
