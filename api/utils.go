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

func ReturnFirstOrDef[T any](arr []T) T {
    var t T 
    if len(arr) == 0 {
        return t
    }
    return arr[0]
}
