package views

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	gloss "github.com/charmbracelet/lipgloss"
	ch_types "github.com/Ch-aOS-Ch/Ch-apetanios/types"
)

var activeTabBorderRight = gloss.Border{
	Top: "─",
	Left: "│",
	Right: "│",
	Bottom: "",
	TopLeft: "╭",
	TopRight: "╮",
	BottomLeft: "╯",
	BottomRight: "╰",
}
var activeTabBorderLeft = gloss.Border{
	Top: "─",
	Left: "│",
	Right: "│",
	Bottom: "",
	TopLeft: "╭",
	TopRight: "╮",
	BottomLeft: "│",
	BottomRight: "╰",
}
var activeTabBorderMiddle = gloss.Border{
	Top: "─",
	Left: "│",
	Right: "│",
	Bottom: "",
	TopLeft: "╭",
	TopRight: "╮",
	BottomLeft: "╯",
	BottomRight: "╰",
}

var (
	colorHighlight = gloss.Color("#7D56F4")
	colorMuted     = gloss.Color("#626262")

	activeTabStyle = gloss.NewStyle().
			BorderForeground(colorHighlight).
			Padding(0, 1).
			BorderBottom(false).
			Margin(1, 0, -2).
			Foreground(colorHighlight).
			Bold(true)

	inactiveTabTopStyle = gloss.NewStyle().
			Border(gloss.RoundedBorder()).
			BorderForeground(colorMuted).
			Padding(0, 1).
			Margin(2, 0, 0).
			BorderBottom(true).
			Foreground(colorMuted)

	windowStyle = gloss.NewStyle().
			Border(gloss.RoundedBorder()).
			BorderForeground(colorHighlight).
			Padding(1, 2).
			BorderTop(false)

	sidebarStyle = gloss.NewStyle().
	  Width(30).
	  Padding(1, 1).
	  Border(gloss.RoundedBorder()).
	  BorderRight(true).
    BorderLeft(false).
    BorderTop(false).
    BorderBottom(false).
    BorderForeground(colorMuted)

	selectedSidebarItemStyle = gloss.NewStyle().
		Foreground(colorHighlight).
	  PaddingLeft(1).
		Bold(true).
		Border(gloss.NormalBorder(), false, false, false, true).
		BorderForeground(colorHighlight)

	unselectedSidebarItemStyle = gloss.NewStyle().
		Foreground(colorMuted).
	  PaddingLeft(2)

	detailStyle = gloss.NewStyle().
		Padding(1, 2).
		Border(gloss.RoundedBorder()).
	  Align(gloss.Left)

	activeCurrentStyle = gloss.NewStyle().
		Border(gloss.RoundedBorder()).
		BorderForeground(colorHighlight).
		Padding(1).
		Margin(1)
)

func getActiveTabBorder(activeTabint int, totalTabs int) gloss.Border {
	switch activeTabint {
		case totalTabs - 1:
			return activeTabBorderRight
		case 0:
			return activeTabBorderLeft
	}
	return activeTabBorderMiddle
}

func renderInactiveTab(text string, isFirst bool) string {
	topPart := inactiveTabTopStyle.Render(text)

	width := gloss.Width(topPart)

	lineChar := "─"
	leftCorner := "─"
	if isFirst{
		leftCorner = "╭"
	}

	rightCorner := "─"

	bottomLine := gloss.NewStyle().
		Foreground(colorHighlight).
		Render(leftCorner + strings.Repeat(lineChar, width-2) + rightCorner)

	return gloss.JoinVertical(gloss.Left, topPart, bottomLine)
}

type Model struct {
	report    ch_types.ChaosReport
	activeTab int
	subActiveTab int
	tabs      []string
	width     int
	height    int
	loaded    bool
	err       error

	applyModel ApplyViewModel
}

func initialModel(report_location string) Model {
	file, err := os.Open(report_location)
	var report ch_types.ChaosReport
	loaded := false
	if err == nil {
		byteValue, _ := io.ReadAll(file)
		json.Unmarshal(byteValue, &report)
		loaded = true
		file.Close()
	}

	return Model{
		report:    report,
		activeTab: 0,
		tabs:      []string{"Apply", "Secrets", "Explain", "Ramble", "Team", "Stats"},
		loaded:    loaded,
		err:       err,

		applyModel: *NewApplyViewModel(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (model Model) View() string {
	if model.err != nil {
		return fmt.Sprintf("Error loading report: %v", model.err)
	}

	var renderedTabs []string

	usedBorder := getActiveTabBorder(model.activeTab, len(model.tabs))
	currentActiveStyle := activeTabStyle.Border(usedBorder)

	for index, tab := range model.tabs {
		if index == model.activeTab {
			renderedTabs = append(renderedTabs, currentActiveStyle.Render(tab))
		} else {
			isFirst := (index == 0)
			renderedTabs = append(renderedTabs, renderInactiveTab(tab, isFirst))
		}
	}

	row := gloss.JoinHorizontal(gloss.Bottom, renderedTabs...)

	row_compesation := 2
	gap := model.width - gloss.Width(row) - row_compesation

	var filler string
	if gap > 0 {
		var fillerLines string
		fillerLines = strings.Repeat("─", gap+1) + "╮"

		filler = gloss.NewStyle().
			Foreground(colorHighlight).
			Render(fillerLines)
	}

	header := gloss.JoinHorizontal(gloss.Bottom, row, filler)

	docStyle := windowStyle.
		Width(model.width - 2).
		Height(model.height - gloss.Height(header))

	contentHeight := model.height - gloss.Height(header) - 2

	var content string
	switch model.activeTab {
	case 0:
		content = model.applyModel.View(model.width - 2, contentHeight)
	case 1:
		content = "Secrets under construction..."
	case 2:
		content = "Explain under construction..."
	case 3:
		content = "Ramble under construction..."
	case 4:
		content = "Team under construction..."
	case 5:
		content = "Stats under construction..."
	}

	window := docStyle.Render(content)

	return gloss.JoinVertical(gloss.Left, header, window)
}

func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        model.width = msg.Width
        model.height = msg.Height
        return model, nil

    case tea.KeyMsg:
        if msg.String() == "ctrl+c" || msg.String() == "q" {
            return model, tea.Quit
        }

        if !model.applyModel.FocusRight {
            switch msg.String() {
            case "right", "l", "tab":
                model.activeTab++
                if model.activeTab >= len(model.tabs) {
                    model.activeTab = 0
                }
                return model, nil

            case "left", "h", "shift+tab":
                model.activeTab--
                if model.activeTab < 0 {
                    model.activeTab = len(model.tabs) - 1
                }
                return model, nil
            }
        }
    }

    switch model.activeTab {
    case 0:
        model.applyModel.Update(msg)
    }

    return model, nil
}

func MainView() {
	program := tea.NewProgram(initialModel("./chaos_report.json"), tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}
