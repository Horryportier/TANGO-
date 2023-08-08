package api

import (
	"os"
	"strings"
)


var (
    ENABLE_STYLE = func () bool {
     switch os.Getenv("TANGO_STYLE") {
        case "true": 
            return true
        case "false": 
            return false
        default: 
            return true
     }
    }()
    ALTSCREEN = func () bool  {
      switch os.Getenv("TANGO_ALT_SCREEN") {
        case "true": 
            return true
        case "false": 
            return false
        default: 
            return false
     }    
    }()
    THEME = func () []string {
        env := os.Getenv("TANGO_THEME")
        col := strings.Split(env, ";")
        if len(col) != 6 {
            return []string{"#FFFFFF","#cdd6f4","#f38ba8","#cba6f7", "#a6e3a1","#313244" }
        }
        return col
    }()
)
