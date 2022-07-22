package userinterface

import (
	"fmt"
	"os"
	"strings"

	"github.com/DavidWenkemann/Masterarbeit/store"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

var statusStore = "scan"

type modelUI struct {
	tabs        []string // items on tab list
	selectedTab int

	//Store
	cartTable      table.Model //for bubble table
	scanInputStore textinput.Model
	selectedOption int

	//Reporting
	reportingTable table.Model //for bubble table
}

func initialModel() modelUI {

	siStore := textinput.New()
	siStore.Placeholder = "" //"Scan to add Item"
	siStore.Focus()
	siStore.CharLimit = 156
	siStore.Width = 20

	return modelUI{
		tabs:        []string{"Store", "Basedata", "Reporting", "Warehouse"},
		selectedTab: 0,

		//Store
		cartTable: table.New([]table.Column{
			table.NewColumn(columnKeyName, "Name", 30),
			table.NewColumn(columnKeyPrice, "Price", 8),
		}),
		scanInputStore: siStore,
		selectedOption: 0,

		//Reporting
		reportingTable: table.New([]table.Column{
			table.NewColumn(columnKeyReporting1, "EAN", 13),
			table.NewColumn(columnKeyReporting2, "Name", 25),
			table.NewColumn(columnKeyReporting3, "Price", 6),
			table.NewColumn(columnKeyReporting4, "Quantity", 5),
		}),
	}
}

func (m modelUI) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m modelUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	//Key bindings for tab store
	if m.selectedTab == 0 {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				//gets scanned input and puts it into cart
				query := strings.TrimSpace(m.scanInputStore.Value())
				//if query != "" {
				scanItem(query)
				//}
			case " ":
				if statusStore == "scan" {
					statusStore = "checkout"
				} else if statusStore == "checkout" {
					switch m.selectedOption {
					case 0:
						store.SellCart()
						statusStore = "scan"
					case 1:
						statusStore = "scan"
					case 2:
						store.ClearCart()
						statusStore = "scan"
					}
				}
			case "right":
				if m.selectedOption < 2 {
					m.selectedOption += 1
				} else {
					m.selectedOption = 0
				}
			case "left":
				if m.selectedOption > 0 {
					m.selectedOption -= 1
				} else {
					m.selectedOption = 2
				}
			}

		}
	}

	switch msg := msg.(type) {
	// Case Key Press
	case tea.KeyMsg:

		// Switch between presses
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		case "tab":
			//m.selectedTab = 1
			if m.selectedTab < len(m.tabs)-1 {
				m.selectedTab++
			} else {
				m.selectedTab = 0
			}
		}

	}

	//m.scanInputStore, cmd = m.textInput.Update(msg)

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m modelUI) View() string {
	//physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
	doc := strings.Builder{}

	if m.selectedTab == 0 {
		//StoreUI()

		return StoreUI(m)

	}

	if m.selectedTab == 1 {
		{
			// Tabs
			row := lipgloss.JoinHorizontal(
				lipgloss.Top,
				tab.Render("Store"),
				activeTab.Render("Warehouse"),
				tab.Render("Reporting"),
				tab.Render("Basedata"),
			)

			gap := tabGap.Render(strings.Repeat(" ", max(0, width-lipgloss.Width(row)-2)))
			row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)
			doc.WriteString(row + "\n\n")
		}
	}

	if m.selectedTab == 2 {
		{
			return ReportingUI(m)
		}
	}

	if m.selectedTab == 3 {
		{
			// Tabs
			row := lipgloss.JoinHorizontal(
				lipgloss.Top,
				tab.Render("Store"),
				tab.Render("Warehouse"),
				tab.Render("Reporting"),
				activeTab.Render("Basedata"),
			)

			gap := tabGap.Render(strings.Repeat(" ", max(0, width-lipgloss.Width(row)-2)))
			row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)
			doc.WriteString(row + "\n\n")
		}
	}

	return docStyle.Render(doc.String())
}

//Starting Point for main function
func StartUI() {
	p := tea.NewProgram(initialModel())

	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}
