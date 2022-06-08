package basedata

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-table/table"
)

func StartUI() {

	t := textinput.NewModel()
	t.Focus()

	//Starts UI
	err := tea.NewProgram(initialModel(t), tea.WithAltScreen()).Start()
	if err != nil {
		fmt.Printf("Help, there's been an error: %v", err)
		os.Exit(1)
	}

}

//add const for columnanes in bubbletable
const (
	columnKeyEan  = "ean"
	columnKeyName = "name"
)

type model struct {
	//count    int
	choices []string // items on the to-do list
	cursor  int      // which to-do list item our cursor is pointing at
	state   string

	//for bubble table
	simpleTable table.Model
	data        []Product

	textInput textinput.Model
	typing    bool

	newEan   string
	newName  string
	newPrice float64
}

func initialModel(t textinput.Model) model {
	return model{
		// Our shopping list is a grocery list
		choices: []string{"Add Product", "Delete Product"},

		//state: "overview",
		state: "overview",

		simpleTable: table.New([]table.Column{
			table.NewColumn(columnKeyEan, "EAN", 13),
			table.NewColumn(columnKeyName, "Name", 30),
		}),
		data:      GetAllProducts(),
		textInput: t,
		typing:    true,
		newEan:    "",
		newName:   "",
		newPrice:  0.0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "left", "up":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "right", "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":

			if m.state == "overview" {

				if m.cursor == 0 {
					m.state = "addean"
				}

				if m.cursor == 1 {
					m.state = "delete"
					m.typing = true

				}

			}
			if m.typing && m.state == "delete" {
				query := strings.TrimSpace(m.textInput.Value())
				if query != "" {
					m.typing = false
					RemoveProductByEAN(query)
					m.textInput.Reset()
					m.state = "overview"
				}
			}
			if m.typing && m.state == "addean" {
				query := strings.TrimSpace(m.textInput.Value())

				if query != "" && m.newEan == "" {
					//m.typing = false
					m.newEan = query
					m.textInput.Reset()
					m.state = "addname"
				}
			}
			if m.typing && m.state == "addname" {
				query := strings.TrimSpace(m.textInput.Value())

				if query != "" {
					//m.typing = false
					m.newName = query
					m.textInput.Reset()
					m.state = "addprice"
				}
			}
			if m.typing && m.state == "addprice" {
				query := strings.TrimSpace(m.textInput.Value())

				if query != "" {
					m.typing = false

					if p, err := strconv.ParseFloat(query, 64); err == nil {
						m.newPrice = p

						//fmt.Println(p) // 3.1415927410125732
					}

					//m.newPrice = query
					AddProduct(m.newEan, m.newName, m.newPrice)
					m.textInput.Reset()
					m.newEan = ""
					m.newName = ""
					m.newPrice = 0.0
					m.state = "overview"
				}
			}
		}
	}

	if m.state == "overview" {
		//Refresh Products after adding or deleting
		m.data = GetAllProducts()
		m.typing = false
	}

	if m.state == "addean" {
		m.typing = true
		var cmd tea.Cmd

		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd
	}

	if m.state == "addname" {
		m.typing = true
		var cmd tea.Cmd

		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd
	}

	if m.state == "addprice" {
		m.typing = true
		var cmd tea.Cmd

		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd
	}

	if m.state == "delete" {
		var cmd tea.Cmd
		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {

	s := "Error occured - restart:\n\n"

	if m.state == "overview" {

		s = addUIstateOverwiew(m)

	} else if m.state == "addean" {

		if m.typing {
			//s = "Status add ean\n\n"
			return fmt.Sprintf("Type in EAN you want to add: \n%s", m.textInput.View())
		}

	} else if m.state == "addname" {

		if m.typing {
			//s = "Status add name\n\n"

			return fmt.Sprintf("Type in Name for EAN%s: \n%s", m.newEan, m.textInput.View())
		} else {
			s = "Status add:\n\n"
		}

	} else if m.state == "addprice" {

		if m.typing {
			//s = "Status add name\n\n"

			return fmt.Sprintf("Type in Name for %s: \n%s", m.newName, m.textInput.View())
		} else {
			s = "Status add:\n\n"
		}

	} else if m.state == "delete" {

		//m.textInput.Focus()
		//m.textInput.Blink()

		if m.typing {
			return fmt.Sprintf("Type in EAN you want to remove: \n%s", m.textInput.View())
		}

		s = "Status delete:\n\n"

	}

	// Send the UI for rendering
	return s

}

func generateRowsFromData(refreshedproducts []Product) []table.Row {
	rows := []table.Row{}

	for i := 0; i <= len(refreshedproducts)-1; i++ {
		row := table.NewRow(table.RowData{
			columnKeyEan:  refreshedproducts[i].EAN,
			columnKeyName: refreshedproducts[i].Name,
		})
		rows = append(rows, row)
	}

	return rows
}

func addUIstateOverwiew(m model) string {

	// The header
	s := "Welcome to the basedata configuration test:\n\n"

	//------------------
	// show all products
	//------------------

	//refreshes data and adds it as string
	m.simpleTable = m.simpleTable.WithRows(generateRowsFromData(m.data))
	body := strings.Builder{}
	body.WriteString(m.simpleTable.View())
	s += body.String()

	s += fmt.Sprintln() //new line

	//------------------
	// adds choices
	//------------------
	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Render the row
		s += fmt.Sprintf("%s %s\t", cursor, choice)
	}

	// The footer
	s += fmt.Sprintln()
	s += "\nPress ctrl+c to quit.\n"

	return s
}
