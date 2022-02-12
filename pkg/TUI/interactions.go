package interactions

import (
	"fmt"
	"os"

    ops "github.com/Tempost/notepad-tui/pkg/operations"
	tea "github.com/charmbracelet/bubbletea"
)

// this package will cover all TUI operations
// including building/rendering the UI, handling key events and sending off data to the database
// through the use of the database package

func Startup() StartupMenu {
    return StartupMenu {
        choices: []string{"Projects", "Personal", "Class", "Scratchpad", "Add new note", "Save work"},
        selected: make(map[int]struct{}),
    }
}

func (s StartupMenu) Init() tea.Cmd {
    // no I/O right now please
    return nil
}

// NOTE: Avoid any use of concurrency here bubbletea doesn't like it too much
func (s StartupMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

        case tea.KeyMsg:

            switch msg.String() {

            // These keys should exit the program.
            case "ctrl+c", "q":
                // TODO(Cody): Save everything to database
                return s, tea.Quit

            // The "up" and "k" keys move the cursor up
            case "up", "k":
                if s.cursor > 0 {
                    s.cursor--
                } else if s.cursor == 0 {
                    s.cursor = len(s.choices) - 1
                }

            // The "down" and "j" keys move the cursor down
            case "down", "j":
                if s.cursor < len(s.choices) - 1 {
                    s.cursor++
                } else if s.cursor == len(s.choices) - 1{
                    s.cursor = 0
                }

            // The "enter" key and the spacebar (a literal space) toggle
            // the selected state for the item that the cursor is pointing at.
            case "enter", " ":
                _, ok := s.selected[s.cursor]
                if ok {
                    delete(s.selected, s.cursor)
                } else {
                    s.selected[s.cursor] = struct{}{} 

                }
            }
        }

        // Return the updated model to the Bubble Tea runtime for processing.
        // Note that we're not returning a command.
        return s, nil
}

func (s StartupMenu) View() string {
    // The header
    prompt :=fmt.Sprintf("Hello %s!\n\n", os.Getenv("USER")) 

    // Iterate over our choices
    for i, choice := range s.choices {

        // Is the cursor pointing at this choice?
        cursor := " " // no cursor
        if s.cursor == i {
            cursor = ">" // cursor!
        }

        // Is this choice selected?
        checked := " " // not selected
        if _, ok := s.selected[i]; ok {
            checked = "x" // selected!
        }

        // Render the row
        prompt += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
    }

    // The footer
    prompt += "\nPress q to quit and save your work.\n"

    // Send the UI for rendering
    return prompt
}
