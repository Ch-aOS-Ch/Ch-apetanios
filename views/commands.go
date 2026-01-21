package views

import (
	"github.com/Ch-aOS-Ch/Ch-apetanios/commands"
	tea "github.com/charmbracelet/bubbletea"
)

func FetchRolesCmd() tea.Cmd {
	return func() tea.Msg {
		tags := commands.FetchChaosRoles()
		return ChaosRolesLoadedMsg{Tags: tags, Err: nil}
	}
}
