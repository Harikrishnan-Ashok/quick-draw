package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	// Styles for lists
	docStyle  = lipgloss.NewStyle().Margin(1, 1)
	listStyle = lipgloss.NewStyle()

	// Style for the active (selected) list
	activeListStyle = lipgloss.NewStyle().BorderForeground(lipgloss.Color("205")).Bold(true)
)

const (
	l1 int = iota // Constant for the first list (System Commands)
	l2            // Constant for the second list (Shortcuts)
)

// Custom item struct representing list entries.
type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

// Main model struct holding two lists and the active list index.
type model struct {
	list1      list.Model
	list2      list.Model
	activeList int // Tracks which list is currently active (l1 or l2)
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
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "tab":
			// Switch between the two lists when Tab is pressed.
			if m.activeList == l1 {
				m.activeList = l2
			} else {
				m.activeList = l1
			}

		case "enter":
			// Handle item selection in the active list.
			var selectedItem item
			if m.activeList == l1 {
				selectedItem = m.list1.SelectedItem().(item)
			} else {
				selectedItem = m.list2.SelectedItem().(item)
			}

			// Execute system commands based on selection.
			switch selectedItem.title {
			case "Shutdown":
				exec.Command("shutdown", "now").Run()
				return m, tea.Quit
			case "Reboot":
				exec.Command("reboot").Run()
				return m, tea.Quit
			case "Firefox":
				exec.Command("firefox").Run()
			case "File Manager":
				exec.Command("lf").Run() // Replace with your file manager command.
			}
		}

	case tea.WindowSizeMsg:
		// Handle resizing of the window.
		h, v := docStyle.GetFrameSize()
		listWidth := (msg.Width - h) / 2 // Divide available width between the two lists
		m.list1.SetSize(listWidth, msg.Height-v)
		m.list2.SetSize(listWidth, msg.Height-v)
	}

	// Update only the active list.
	var cmd tea.Cmd
	if m.activeList == l1 {
		m.list1, cmd = m.list1.Update(msg)
	} else {
		m.list2, cmd = m.list2.Update(msg)
	}

	return m, cmd
}

// View function renders both lists side by side with styles.
func (m model) View() string {
	// Style the active list differently to indicate it's selected.
	var list1View, list2View string
	if m.activeList == l1 {
		list1View = activeListStyle.Render(m.list1.View()) // Highlight active list1
		list2View = listStyle.Render(m.list2.View())
	} else {
		list1View = listStyle.Render(m.list1.View())
		list2View = activeListStyle.Render(m.list2.View()) // Highlight active list2
	}

	// Render the two lists side by side.
	return docStyle.Render(
		lipgloss.JoinHorizontal(lipgloss.Top, list1View, list2View),
	)
}

func main() {

	// Items for Shutdown and Reboot options.
	items := []list.Item{
				item{title: "Firefox", desc: "Open firefox in a new window"},
		item{title: "File Manager", desc: "Open File Manager"},
			}

	// Items for Shortcuts.
	shortcuts := []list.Item{
		item{title: "Shutdown", desc: "Turn off the system"},
		item{title: "Reboot", desc: "Restart the system"},

	}

	// Initialize the first list (System Commands).
	list1 := list.New(items, list.NewDefaultDelegate(), 0, 0)
	list1.Title = "Shortcuts"

	// Initialize the second list (Shortcuts).
	list2 := list.New(shortcuts, list.NewDefaultDelegate(), 0, 0)
	list2.Title = "Power Menu"
	list2.SetShowHelp(false)

	// Initialize the main model with both lists, and set the first list as active.
	m := model{list1: list1, list2: list2, activeList: l1}

	// Start the Bubble Tea program with alternate screen.
	p := tea.NewProgram(m, tea.WithAltScreen())

	// Run the program and handle any errors.
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
