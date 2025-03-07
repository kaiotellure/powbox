package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
}

func Progress() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	
	return m, nil
}

func (m model) View() string {
	s := "Working...\n\n"
	s += "ctrl+c/q close\n"
	return s
}
