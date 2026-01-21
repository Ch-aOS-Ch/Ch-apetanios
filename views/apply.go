package views

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	gloss "github.com/charmbracelet/lipgloss"
	"fmt"
)

type ApplyViewModel struct {
	MenuItems []string
	Cursor		int
	FocusRight	bool
	Tags			[]string
	SelectedTag int
}

func NewApplyViewModel() *ApplyViewModel {
	return &ApplyViewModel{
		MenuItems: []string{"Select Tags", "Config", "Secrets", "Run"},
		Cursor:    0,
		FocusRight: false,
		Tags:      []string{"packages", "services", "network", "users", "firewall"},
		SelectedTag: 0,
	}
}

func (model *ApplyViewModel) Update(msg tea.Msg) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			model.FocusRight = false
		case "up", "k":
			if model.FocusRight {
				if model.Cursor == 0 && model.SelectedTag > 0 {
					model.SelectedTag--
				}
			} else if model.Cursor > 0 {
				// Navigate Menu
				model.Cursor--
			}
		case "down", "j":
			if model.FocusRight {
				if model.Cursor == 0 && model.SelectedTag < len(model.Tags)-1 {
					model.SelectedTag++
				}
			} else if model.Cursor < len(model.MenuItems)-1 {
				model.Cursor++
			}
		case "enter", " ":
			model.FocusRight = true
		}
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
            cursor := "  "
            style := gloss.NewStyle().Foreground(colorMuted)
            if i == m.SelectedTag {
                cursor = "->"
                style = gloss.NewStyle().Foreground(colorHighlight).Bold(true)
            }
            fmt.Fprintf(&list, "%s %s\n", cursor, style.Render(tag))
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
