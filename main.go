package main

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main(){
	p := tea.NewProgram(initialModel())

	if _,err := p.Run(); err!= nil{
		println("Error: ", err.Error())
		os.Exit(1)
	}
}
