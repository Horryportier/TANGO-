package tui

import (
	"fmt"

	jisho "github.com/Horryportier/go-jisho"
	api "github.com/Horryportier/tango/api"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/google/go-cmp/cmp"

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

    index int
    length int
 }

func initialModel() model {
    input := textinput.New()
    input.Placeholder = "search/サーチ"
    input.Focus()
    spin := spinner.New()
    spin.Spinner = spinner.Dot
    spin.Style = api.AcentStyle
	return model{
        input: input,
        err: fmt.Errorf(""),
        spin:  spin,
	}
}

func (m model) Init() tea.Cmd {
    return  nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    m.index = len(m.data.Data)
    var cmds []tea.Cmd
    var cmd tea.Cmd

    switch msg := msg.(type) {

    case tea.KeyMsg:

        switch msg.String() {

        case "ctrl+c", "q":
            return m, tea.Quit

        case "tab":
            m.mode = func (m Mode) Mode {
                if m == Input {return List}
                if m == List {return Input}
                return Input
            }(m.mode)
        case "up", "k":

        case "down", "j":

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
        m.data = res.WordData
        return m, nil
    default:
    }

    if m.mode == Input {
    m.input, cmd = m.input.Update(msg)
    cmds = append(cmds, cmd)
    }
    m.spin, cmd = m.spin.Update(msg)
    cmds = append(cmds, cmd)
    return m, tea.Batch(cmds...)
}


func (m model) View() string {
    word := func () string  {
        if cmp.Equal(m.data, jisho.WordData{})    {
            return "no data"
        } 
        return api.PrintWord(m.data, false)
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
    
        
    text := api.TextFrom([]api.Line{
           api.LineFrom([]api.Span{
                api.SpanFrom("Welcome to TANGO!", api.DefStyle),
                api.SpanFrom("Try Searching.", api.AcentStyle),
           }),
           api.LineFrom([]string{m.input.View(), spinView}, api.AcentStyle),
           api.LineFrom("Result:", api.DimStyle),
           api.LineFrom(word, api.DimStyle),
           api.LineFrom(m.err.Error(), api.ErrorStyle),
          
    })
    return text.Render(api.ENABLE_STYLE)
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



