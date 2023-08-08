package api

import (
	jisho "github.com/Horryportier/go-jisho"
)

var (
    eStrArray = []string{""}
)


func DefWord() jisho.WordData {
    return jisho.WordData{
        Data: 
        []jisho.Data{
            {Slug: "",
            IsCommon: false,
            Tags: eStrArray,
            Jlpt: eStrArray,
            Japanese: []jisho.Japanese{
                {
                    Word: "",
                    Reading: ""},
                },
            Senses: []jisho.Senses {
                {
                EnglishDefinitions: eStrArray,
                PartsOfSpeech: eStrArray,
                },
            },
        }}}
}

func ReturnFirstOrDef[T any](arr []T, opt ...int) T {
    var index int
    index = func () int {
        for _, i := range opt { 
            return i
        }
      return 0
    }()
    var t T 
    if len(arr) == 0 || index >= len(arr) {
        return t
    }
    return arr[index]
}

func ReturnFirstOrDefSlice[T any](arr []T , indexes ...interface{}) []T {
    var res []T
    for _,t := range indexes {
        switch t.(type) {
            case int: 
                res = append(res, ReturnFirstOrDef(arr, t.(int)))
            case []int:
                for _,i := range t.([]int) {
                     res = append(res, ReturnFirstOrDef(arr, i))
                }
        }
    }
    return res
}

func ClearEmptyStr(s []string) []string {
    var n []string 
    for _,v := range s {
        if v != "" {
            n = append(n, v)
        }
    }
    return s
}


