package views
import (
	ch_types "github.com/Ch-aOS-Ch/Ch-apetanios/types"
)

type ApplyViewModel struct {
	MenuItems []string
	Cursor		int
	FocusRight	bool

	Tags			[]string
	SelectedTag int
	CheckedTags	map[int]bool

	IsRunning bool
	Output string
	OutputError error
	IsLoadingTags bool
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

type ChaosFinishedMsg struct {
	err error
	output string
}

type ChaosRolesLoadedMsg struct {
	Tags []string
	Err   error
}
