package views

import(
	tea "github.com/charmbracelet/bubbletea"
)

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

		case " ", "enter":
			if model.FocusRight && model.Cursor == 0 {
					model.CheckedTags[model.SelectedTag] = !model.CheckedTags[model.SelectedTag]
			} else if !model.FocusRight {
				model.FocusRight = true
			}
		}
	}
}

