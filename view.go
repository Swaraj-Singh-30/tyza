package main

import "fmt"

const banner = `
 /$$$$$$$$ /$$     /$$ /$$$$$$$$  /$$$$$$ 
|__  $$__/|  $$   /$$/|_____ $$  /$$__  $$
   | $$    \  $$ /$$/      /$$/ | $$  \ $$
   | $$     \  $$$$/      /$$/  | $$$$$$$$
   | $$      \  $$/      /$$/   | $$__  $$
   | $$       | $$      /$$/    | $$  | $$
   | $$       | $$     /$$$$$$$$| $$  | $$
   |__/       |__/    |________/|__/  |__/
`

func (m model) View() string {

	switch m.screen {

	case "typing":
		return fmt.Sprintf("%s\nTyping test screen\n\nPress 'b' to go back\n", banner)

	case "logs":
		return "Logs screen\n\nPress 'b' to go back\n"

	case "other":
		return "Other stuff\n\nPress 'b' to go back\n"

	default:
		return m.list.View()
	}
}
