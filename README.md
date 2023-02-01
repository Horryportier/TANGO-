# TANGO-単語
***(単語) Tango means word***

This Tui is an japanese dictionary for ppl that can't bother to use broswer. ;)

## !Warning
***only romaji, engilish, hiragana and katakana are supported no KANJI.***
Also some engilish words can be interpreted as romaji and searched in that way. 
Its something that i can not fix bc its an jisho api problem.
(exp. china -> ちな witch means "by the way" not country)



## How to use 
type any word into the input (hiragana, katakana, romaji or engilish) and you will be presented with list of translations.
By pressing y/cntl+c copies dietails view to ***clipboard***.
Witch is formated into markdown so you can eazliy paste it into your notes.

<img src="https://raw.githubusercontent.com/Horryportier/TANGO-/main/TANGO.gif" width=512 height=512/>

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

## Build with

- ![Bubbletea](https://github.com/charmbracelet/bubbletea)
- ![go-jisho](https://github.com/Horryportier/go-jisho) **(jisho api wrapper)**
- ![glamour](https://github.com/charmbracelet/glamour) 
- ![lipgloss](https://github.com/charmbracelet/lipgloss) 
- ![clipboard](https://github.com/atotto/clipboard)


