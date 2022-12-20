package app


import (
        jisho "github.com/Horryportier/go-jisho"
)

type item struct {
        title string
        descrition string
        data jisho.Data
}



func (i item) Titile() string { return i.title }
func (i item) Descrition() string { return i.descrition}
func (i item) FilterValue() string { return i.title }


