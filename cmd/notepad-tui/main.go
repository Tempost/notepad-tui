package main

// In time this package will use the database and TUI package to create the final
// notepad application. Similar to how the microsoft office notes application works
// except all through the TUI (: uses the $EDITOR env variable to open .md files
// to act as the "notepad" all formatting is markdown
// Allow notes to be printed AND formatted in the TUI
// NOTE: Maybe since most has really pretty markdown styling and highlighting we can
// "print" the notes via that command

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func initModel() model {
	return model{
		choices:  []string{"Buy carrots", "Buy green vegies", "Buy kohlrabi"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	// returning nil just means no IO at the moment please
	return nil
}

// update each render
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			} else if m.cursor == 0 {
				m.cursor = len(m.choices) - 1
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			} else if m.cursor == len(m.choices)-1 {
				m.cursor = 0
			}

		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	header := "What should we buy at the mark?\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		header += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	header += "\nPress q to quit.\n"

	return header
}

func main() {
	p := tea.NewProgram(initModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, [ Error ] has occured: %v", err)
		os.Exit(1)
	}
}
