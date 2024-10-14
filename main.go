package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

// Custom item struct representing list entries.
type item struct {
	title, desc string
}

func (i item) Title() string { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

// Main model struct holding the list.
type model struct {
	list list.Model
}

// Init function for Bubble Tea, nothing to initialize here.
func (m model) Init() tea.Cmd {
	return nil
}

// Update handles key presses and window resizing.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Handle quitting the program with Ctrl+C.
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		// Handle item selection with Enter/Return key.
		if msg.String() == "enter" {
			selectedItem := m.list.SelectedItem()

			// Check which item is selected and simulate a system command.
			switch selectedItem.(item).title {
			case "Shutdown":
				exec.Command("shutdown", "now").Run()
				return m,tea.Quit
			case "Reboot":
				 exec.Command("reboot").Run()
				return m,tea.Quit
			}

			// Return the model without quitting.
			return m, nil
		}

	case tea.WindowSizeMsg:
		// Handle resizing of the window.
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	// Update the list.
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

// View function renders the list view with styles.
func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func main() {

	// Items for Shutdown and Reboot options.
	items := []list.Item{
		item{title: "Shutdown", desc: "Turn off the system"},
		item{title: "Reboot", desc: "Restart the system"},
	}

	// Initialize the list with items.
	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "System Commands"

	// Start the Bubble Tea program with alternate screen.
	p := tea.NewProgram(m, tea.WithAltScreen())

	// Run the program and handle any errors.
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
