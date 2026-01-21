package views

import (
	gloss "github.com/charmbracelet/lipgloss"
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

