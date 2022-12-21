package main

import (
	"log"
	app "src/tango/v1/app"
)

func main() {
        if err := app.Start(); err != nil{
                log.Fatal(err)
        }
}
