package utils

import (
	"log"
)

func Try(args ...any) any {
	for _, arg := range args {
		if _, ok := arg.(error); ok {
			if arg != nil {
				log.Fatal(arg)
			}
		}

	}
	return args
}
