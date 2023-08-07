package tui

import (
	"fmt"

	jisho "github.com/Horryportier/go-jisho"
	api "github.com/Horryportier/tango/api"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"

	tea "github.com/charmbracelet/bubbletea"
)

type  DataCh  struct { 
    data jisho.WordData
    err error
}

type errMsg error

func Run() api.Err {
    p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        api.PrintErr(err, )
        return err
    }
    return nil
}

type model struct {
    input textinput.Model
    dataView textarea.Model
    help textarea.Model
    data  jisho.WordData
    datach chan DataCh
    err errMsg
 }

func initialModel() model {
    input := textinput.New()
    input.Placeholder = "search/サーチ"
    input.Focus()
	return model{
        input: input,
        data: api.DefWord(),
        err: fmt.Errorf(""),
	}
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmds []tea.Cmd
    var cmd tea.Cmd

    switch msg := msg.(type) {

    case tea.KeyMsg:

        switch msg.String() {

        case "ctrl+c", "q":
            return m, tea.Quit

        case "up", "k":

        case "down", "j":

        case "enter", " ":
            val := m.input.Value()
            go search(val ,m.datach)
        }
    case errMsg:
        m.err = msg
    }
    select{
    case d := <- m.datach:
        if d.err != nil {
            cmds = append(cmds, errMsgCmd(d.err))
        }
        m.data = d.data
    default: 
    }

    m.input, cmd = m.input.Update(msg)
    cmds = append(cmds, cmd)
    return m, tea.Batch(cmds...)
}


func (m model) View() string {
    text := api.TextFrom([]api.Line{
           api.LineFrom([]api.Span{
                api.SpanFrom("Welcome to TANGO!", api.DefStyle),
                api.SpanFrom("Try Searching.", api.ArrowStyle),
           }),
           api.LineFrom([]api.Span{
                api.SpanFrom(m.input.View(), api.ArrowStyle),
           }),
           api.LineFrom([]api.Span{
                api.SpanFrom("Result", api.DimStyle),
           }),
           api.LineFrom([]api.Span{
                api.SpanFrom(api.PrintWord(m.data, false), api.DimStyle),
           }),
           api.LineFrom([]api.Span{
                api.SpanFrom(m.err.Error(), api.ErrorStyle),
           }),
    })
    return text.Render(api.ENABLE_STYLE)
}



func search(s string, ch chan DataCh) {
    var data DataCh
    if err := data.data.Get(s); err != nil {
        data.err = err 
    }
    ch <- data
}
