package main

import "github.com/charmbracelet/bubbles/list"

type item string

func (i item) Title() string       { return string(i) }
func (i item) Description() string { return "" }
func (i item) FilterValue() string { return string(i) }

type model struct {
	list   list.Model
	screen string 
}

func initialModel() model {
	items := []list.Item{
		item("Test typing speed"),
		item("Check logs"),
		item("Some other stuff"),
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Tyza CLI"

	return model{
		list:   l,
		screen: "menu",
	}
}
