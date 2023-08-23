package tui

import (
	"fmt"

	jisho "github.com/Horryportier/go-jisho"
	api "github.com/Horryportier/tango/api"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/google/go-cmp/cmp"
    . "github.com/Horryportier/lipgloss-text"


	tea "github.com/charmbracelet/bubbletea"
)

type Result struct {jisho.WordData; error}
type Mode int

const (
    Input Mode = iota 
    List
)

var (
    Future chan Result = make(chan Result)
    Finshed chan bool = make(chan bool)
    Searching bool = false
    Help bool = false
    options []tea.ProgramOption
)

type errMsg error

func Run() error {
    if api.ALTSCREEN {
        options = append(options, tea.WithAltScreen())
    }
    p := tea.NewProgram(initialModel(), options...)
    if _, err := p.Run(); err != nil {
        api.PrintErr(err, )
        return err
    }
    return nil
}


type model struct {
    mode Mode

    input textinput.Model
    dataView textarea.Model
    spin spinner.Model

    help textarea.Model

    data  jisho.WordData
    err errMsg

    pagin paginator.Model
    index int
    length int
    
 }

func initialModel() model {
    input := textinput.New()
    input.Placeholder = "search/サーチ"
    input.Focus()
    
    spin := spinner.New()
    spin.Spinner = spinner.Dot
    if api.ENABLE_STYLE {
        spin.Style = api.AcentStyle
    }

    pagin := paginator.New()
    pagin.SetTotalPages(1)
    pagin.Type = paginator.Dots
	return model{
        input: input,
        err: fmt.Errorf(""),
        spin:  spin,
        pagin:  pagin,
	}
}

func (m model) Init() tea.Cmd {
    return  nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmds []tea.Cmd
    var cmd tea.Cmd

    switch msg := msg.(type) {
    case tea.KeyMsg:

        switch msg.String() {

        case "ctrl+c", "q":
            return m, tea.Quit

        case "?":
            Help = !Help
        case "tab":
            m.mode = func (m Mode) Mode {
                if m == Input {return List}
                if m == List {return Input}
                return Input
            }(m.mode)
        case "up", "k":
            m.pagin.PrevPage()
            if m.index > 0 { m.index -= 1; return m, nil}

        case "down", "j":
            m.pagin.NextPage()
            if m.index < len(m.data.Data) { m.index += 1; return m, nil }
        case "enter":
            input := m.input.Value()
            if input != "" {
                go search(input)
                Searching = true
                cmds = append(cmds, m.spin.Tick)
            }
        }
    case errMsg:
        m.err = msg
    }
    select{
    case res := <- Future:
        if res.error != nil {
            m.err = res.error
            return m, nil
        }
        m.index = 0
        m.data = res.WordData
        m.pagin.SetTotalPages(len(m.data.Data))
        m.mode = List
        return m, nil
    default:
    }

    if m.mode == Input {
    m.input, cmd = m.input.Update(msg)
    cmds = append(cmds, cmd)
    }
    m.spin, cmd = m.spin.Update(msg)
    cmds = append(cmds, cmd)
    m.pagin, cmd = m.pagin.Update(msg)
    cmds = append(cmds, cmd)
    return m, tea.Batch(cmds...)
}


func (m model) View() string {
    word := func () string  {
        var t string
        if cmp.Equal(m.data, jisho.WordData{})    {
            t = "no data"
        } 
        t = api.PrintWord(api.ReturnFirstOrDef(m.data.Data, m.index), false)
        if m.mode != List {
            return api.FaintStyle.Render(t)
        }
        return t
    }()

    spinView := func () string {
    select {
    case _ = <- Finshed:
        Searching = false
        return ""
    default:
        if Searching {
            return m.spin.View()
        }
        return ""
    }
    }()
    inputView := func () string  {
        if m.mode != Input {
            return  api.FaintStyle.Render(m.input.View())
        }
        return m.input.View()
    }()
        
    text := TextFrom([]Line{
           LineFrom([]Span{
                SpanFrom("Welcome to TANGO!", api.DefStyle),
                SpanFrom("Try Searching.", api.AcentStyle),
           }),
           LineFrom([]string{inputView, spinView}, api.AcentStyle),
           LineFrom("Result:", api.DimStyle),
           LineFrom(m.pagin.View(), api.AcentStyle),
           LineFrom(word, api.DimStyle),
           LineFrom(m.err.Error(), api.ErrorStyle),
          
    })
    text = append(text, PrintKeys(Help)...)
    return text.Render(Styled(api.ENABLE_STYLE)) 
}


func search(input string) {
   var word jisho.WordData 
   var res Result
   err := word.Get(input)
   if err != nil {
        res.error =  err
        Future <- res
        Finshed <- true
        return
   }
    res.WordData = word
    Future <- res
    Finshed <- true
    return
}




func PrintKeys(help bool) Text {
    var text Text
    if help {
        text = TextFrom([]Line{LineFrom([]string{"j/down", "k/up", "Enter: search"}, api.DimStyle), 
        LineFrom([]string{"Tab: switch input/list", "cntl+c/esc: exit",}, api.DimStyle)})
    } else {
        text  = TextFrom("pres ? to see keys", api.DimStyle)
    }
    return text
}
