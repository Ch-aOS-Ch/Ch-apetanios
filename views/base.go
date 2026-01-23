package views

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	ch_types "github.com/Ch-aOS-Ch/Ch-apetanios/types"
	tea "github.com/charmbracelet/bubbletea"
	gloss "github.com/charmbracelet/lipgloss"
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
	return FetchRolesCmd()
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


func MainView() {
	program := tea.NewProgram(initialModel("./chaos_logbook.json"), tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}
