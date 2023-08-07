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


