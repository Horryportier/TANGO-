package main

import "os"

var (
    ENABLE_STYLE = func () bool {
     switch os.Getenv("TANGO_NO_STYLE") {
        case "true": 
            return true
        case "false": 
            return false
        default: 
            return true
     }
    }()
)
