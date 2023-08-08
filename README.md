# TANGO-単語
***(単語) Tango means word***

This Tui is an japanese dictionary for ppl that can't bother to use broswer. ;)

![Made with VHS](https://vhs.charm.sh/vhs-3I9O80EeylWbYh839V2DxG.gif)

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

## Contirbuting
You are free to contribite.

> List of things that i wan't to do but you might as well help 😄
- check if its working on Mac/Windows
- automatic anki card creation
- automatic obsidian note creation 
- fix visual bug "when theres no data  should not print anything under Result"

you can ask for feature or make it if you want.
