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
    tui "github.com/Tempost/notepad-tui/pkg/TUI"

	tea "github.com/charmbracelet/bubbletea"
)


func main() {
	p := tea.NewProgram(tui.Startup())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, [ Error ] has occured: %v", err)
		os.Exit(1)
	}
}
