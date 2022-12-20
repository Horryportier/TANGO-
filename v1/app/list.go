package app

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
        jisho "github.com/Horryportier/go-jisho"
)

var data jisho.Data

type ListModel struct {
        list list.Model
}


func ListInit() ListModel {
        return ListModel{}
}


func ListUpdate(m model, msg tea.Msg) (model, tea.Cmd){
        return m, nil
}

func ListView(m model) (string){
        return "List"
}
