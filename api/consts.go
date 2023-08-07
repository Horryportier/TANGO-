package api

import "os"


type Err error

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
