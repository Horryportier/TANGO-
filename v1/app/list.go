package app

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ListModel struct {
        list list.Model
}


func ListUpdate(m model, msg tea.Msg) (model, tea.Cmd){
        return m, nil
}

func ListView(m model) (string){
        return ""
}
