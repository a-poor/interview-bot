package main

import (
	_ "github.com/charmbracelet/bubbles/list"
	_ "github.com/charmbracelet/bubbles/textarea"
	_ "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func newApp() *tea.Program {
	m := newModel()
	p := tea.NewProgram(m, tea.WithAltScreen())
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
