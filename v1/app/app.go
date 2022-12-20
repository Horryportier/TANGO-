package app

import (
	utils "src/tango/v1/utils"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
)


type model struct {
        searchBar textinput.Model
        results 
}

func Start() (error) {
        return nil
}
