# TANGO-単語
***(単語) Tango means word***

This Tui is an japanese dictionary for ppl that can't bother to use broswer. ;)


## How to use 
```sh
Tango is an simple English->Japanese cli dictionary.
   this app uses jisho.org open api
---COMMANDS--

NONE                    opens tui
-h                      print help
-r                      don't use colors
たんご/タンゴ/tango/単語  will search for the word
``` 

### Env valiables
```bash
# use to disable colors. Default is true
export TANGO_STYLE=true
# use to open tui in alternative screen. Default is false
export TANGO_ALT_SCREEN=true
# use to change theme. !! theme has to have exacly 6 hex colors seperated by ;
export TANGO_THEME="#FFFFFF;#cdd6f4;#f38ba8;#cba6f7;#a6e3a1;#313244"
```



## Instal 
make sure you have golang installed.
```bash
git clone https://github.com/Horryportier/TANGO-.git
cd TANGO-
go install -v 
```

#### Uninstall 
```bash
rm ~/go/bin/tango
```

## next 
- more flags 
- capabilty of creating anki flashcards


## Build with

- ![Bubbletea](https://github.com/charmbracelet/bubbletea)
- ![go-jisho](https://github.com/Horryportier/go-jisho) **(jisho api wrapper)**
- ![glamour](https://github.com/charmbracelet/glamour) 
- ![lipgloss](https://github.com/charmbracelet/lipgloss) 
- ![clipboard](https://github.com/atotto/clipboard)


