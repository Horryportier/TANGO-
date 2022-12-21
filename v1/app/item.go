package app

import (
	"fmt"

	jisho "github.com/Horryportier/go-jisho"
)

type item struct {
        data jisho.Data
}



func (i item) Titile() string { return i.data.Slug}
func (i item) Descrition() string { return fmt.Sprint(i.data.Jlpt)}
func (i item) FilterValue() string { return i.data.Slug}


func ItemGenerator(data jisho.Data) (item){
        return item{data: data}
}
