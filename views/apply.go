package views

import (
	"strings"
	gloss "github.com/charmbracelet/lipgloss"
	"fmt"
)

func NewApplyViewModel() *ApplyViewModel {
	return &ApplyViewModel{
		MenuItems: []string{"Select Tags", "Config", "Secrets", "Run"},
		Cursor:    0,
		FocusRight: false,
		Tags:      []string{"packages", "services", "network", "users", "firewall"},
		SelectedTag: 0,
		CheckedTags: make(map[int]bool),
	}
}

func (m ApplyViewModel) View(width, height int) string {
    var sidebarRows []string
    for i, item := range m.MenuItems {
        if i == m.Cursor {
            style := selectedSidebarItemStyle
            if m.FocusRight {
                style = style.Copy().Foreground(colorMuted).BorderForeground(colorMuted)
			}
            sidebarRows = append(sidebarRows, style.Render(item))
        } else {
            sidebarRows = append(sidebarRows, unselectedSidebarItemStyle.Render(item))
        }
    }
    sidebarContent := gloss.JoinVertical(gloss.Left, sidebarRows...)
    sidebarView := sidebarStyle.Height(height).Render(sidebarContent)

    var detailContent string

    switch m.Cursor {
    case 0:
      title := "  AVAILABLE TAGS"
      var list strings.Builder

      for i, tag := range m.Tags {
        cursor := ""
				if i == m.SelectedTag && m.FocusRight {
					cursor = "->"
				}

				checkbox := " "
				if m.CheckedTags[i] {
					checkbox = "+"
				}

				lineStyle := gloss.NewStyle().Foreground(colorMuted)
				if m.FocusRight && i == m.SelectedTag {
					lineStyle = gloss.NewStyle().Foreground(colorHighlight).Bold(true)
				} else if m.CheckedTags[i] {
					lineStyle = gloss.NewStyle().Foreground(gloss.Color("#00FF00")).Bold(true)
				}

				fmt.Fprintf(&list, "%s %s %s\n", cursor, lineStyle.Render(checkbox), lineStyle.Render(tag))
      }

      hint := "Press [ENTER] to focus here, [ESC] to go back"
		
      if m.FocusRight { hint = "Press [UP/DOWN] to select, [SPACE] to toggle" }

      detailContent = fmt.Sprintf("%s\n\n%s\n\n%s", 
          gloss.NewStyle().Bold(true).Render(title), 
          list.String(), 
          gloss.NewStyle().Foreground(colorMuted).Italic(true).Render(hint))

    case 1:
        detailContent = "  CONFIGURATION\n\n[ ] Verbose Mode\n[x] Dry Run\n\n(Press Enter to edit config)"
    case 2:
        detailContent = "  SECRETS MANAGER\n\nVault Status: Connected"
    case 3:
        detailContent = "  READY TO DEPLOY"
    }

    availableWidth := width - gloss.Width(sidebarView) - 4

    detailStyle := gloss.NewStyle().
        Padding(1, 2).
        Align(gloss.Left)

    if m.FocusRight {
        detailStyle = detailStyle.
            Border(gloss.RoundedBorder()).
            BorderForeground(colorHighlight).
            Width(availableWidth - 2). 
            Height(height - 2)
    } else {
        detailStyle = detailStyle.
            Width(availableWidth).
            Height(height)
    }

    detailView := detailStyle.Render(detailContent)

    return gloss.JoinHorizontal(gloss.Top, sidebarView, detailView)
}
