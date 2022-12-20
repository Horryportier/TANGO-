package main

import (
	app "src/tango/v1/app"
	"src/tango/v1/utils"
)

func main() {
        utils.Try(app.Start())
}
