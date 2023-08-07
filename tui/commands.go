package tui

import tea "github.com/charmbracelet/bubbletea"


func errMsgCmd(err error) tea.Cmd {
    return func() tea.Msg {
        return errMsgCmd(err)
    }
}
