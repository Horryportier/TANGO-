package utils

import (
	"log"

	"github.com/atotto/clipboard"
)

func Try(args ...any) interface{} {
        var filter []any
	for _, arg := range args {
		if _, ok := arg.(error); ok {
			if arg != nil {
				log.Fatal(arg)
			} 
		}else{ 
                        filter = append(filter, arg)
                }

	}
        if len(filter) == 1{
                return filter[0]
        }
	return filter
}

func MakeRange(min, max int) []int{
        a := make([]int, max-min+1)
        for i := range a {
                a[i] = min + i
        }
        return a
}

func copyToClipord(content string) {
	clipboard.WriteAll(content)
}


