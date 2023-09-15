package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func newApp() *tea.Program {
	p := tea.NewProgram(newModel())
	return p
}

type model struct{}

func newModel() tea.Model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	return "Hello, World!"
}
