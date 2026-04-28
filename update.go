package main

import tea "github.com/charmbracelet/bubbletea"

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch m.screen {

	case "menu":
		var cmd tea.Cmd

		switch msg := msg.(type) {

		case tea.WindowSizeMsg:
			m.list.SetSize(msg.Width, msg.Height)

		case tea.KeyMsg:
			switch msg.String() {

			case "ctrl+c", "q":
				return m, tea.Quit

			case "enter":
				selected := m.list.SelectedItem().(item)

				switch string(selected) {
				case "Test typing speed":
					m.screen = "typing"
				case "Check logs":
					m.screen = "logs"
				case "Some other stuff":
					m.screen = "other"
				}
			}
		}

		m.list, cmd = m.list.Update(msg)
		return m, cmd

	default:
		if key, ok := msg.(tea.KeyMsg); ok {
			switch key.String() {

			case "q", "ctrl+c":
				return m, tea.Quit

			case "b":
				m.screen = "menu"
			}
		}
	}

	return m, nil
}
