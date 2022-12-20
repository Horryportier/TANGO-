package app

import (
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type SearchModel struct {
        input textinput.Model
        pagin paginator.Model
}


func SearchUpdate(m model, msg tea.Msg) (model, tea.Cmd){
        return m, nil
}

func SearchingUpdate(m model, msg tea.Msg) (model, tea.Cmd){
        return m, nil
}
