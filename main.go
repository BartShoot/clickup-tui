package main

import (
	// "context"
	// "encoding/json"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	// "github.com/raksul/go-clickup/clickup"
)

type model struct {
	choices  []string         // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
	spinner  spinner.Model
}

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return model{
		// Our to-do list is a grocery list
		choices: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
		spinner:  s,
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// // The header
	// s := "What should we buy at the market?\n\n"
	//
	// // Iterate over our choices
	// for i, choice := range m.choices {
	//
	// 	// Is the cursor pointing at this choice?
	// 	cursor := " " // no cursor
	// 	if m.cursor == i {
	// 		cursor = ">" // cursor!
	// 	}
	//
	// 	// Is this choice selected?
	// 	checked := " " // not selected
	// 	if _, ok := m.selected[i]; ok {
	// 		checked = "x" // selected!
	// 	}
	//
	// 	// Render the row
	// 	s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	// }
	//
	// // The footer
	// s += "\nPress q to quit.\n"
	//
	// Send the UI for rendering
	s := fmt.Sprintf("\n\n   %s Loading forever...press q to quit\n\n", m.spinner.View())
	return s
}

func main() {
	// client := clickup.NewClient(nil, "insertkey")
	//
	// teams, _, _ := client.Teams.GetTeams(context.Background())
	// spaces, _, _ := client.Spaces.GetSpaces(context.Background(), teams[0].ID, false)
	// folders, _, _ := client.Folders.GetFolders(context.Background(), spaces[0].ID, false)
	// for i, folder := range folders {
	// 	fmt.Printf("%d. %s: %s\n", i, folder.Name, folder.ID)
	// }
	// s, _ := json.MarshalIndent(folders, "", "\t")
	// fmt.Print(string(s))

	// fmt.Printf("%v\n", teams)

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
