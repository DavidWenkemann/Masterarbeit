package userinterface

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/DavidWenkemann/Masterarbeit/Monolith_SyncCommunication/basedata"
	"github.com/DavidWenkemann/Masterarbeit/Monolith_SyncCommunication/database"
	"github.com/DavidWenkemann/Masterarbeit/Monolith_SyncCommunication/model"
	"github.com/DavidWenkemann/Masterarbeit/Monolith_SyncCommunication/store"
	"github.com/DavidWenkemann/Masterarbeit/Monolith_SyncCommunication/warehouse"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

var statusStore = "scan"
var statusWarehouse = "scan"
var statusBasedata = "table"

type modelUI struct {
	tabs        []string // items on tab list
	selectedTab int
	spinner     spinner.Model

	//Store
	cartTable      table.Model //for bubble table
	selectedOption int

	//Warehouse
	lastStocked model.APIProduct

	//Reporting
	reportingTable table.Model //for bubble table

	//Basedata
	basedataTable          table.Model //for bubble table
	selectedOptionBasedata int
	newProduct             model.APIProduct

	//Textinput
	textInput textinput.Model
}

func initialModel() modelUI {

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	//Store

	//Warehouse

	//Basedata

	ti := textinput.New()
	ti.Placeholder = ""
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return modelUI{
		tabs:        []string{"Store", "Basedata", "Reporting", "Warehouse"},
		selectedTab: 0,
		spinner:     s,

		//Store
		cartTable: table.New([]table.Column{
			table.NewColumn(columnKeyName, "Name", 30),
			table.NewColumn(columnKeyPrice, "Price", 8),
		}),
		selectedOption: 0,

		//Warehouse

		//Reporting
		reportingTable: table.New([]table.Column{
			table.NewColumn(columnKeyReporting1, "EAN", 15),
			table.NewColumn(columnKeyReporting2, "Name", 25),
			table.NewColumn(columnKeyReporting3, "Price", 6),
			table.NewColumn(columnKeyReporting4, "Quantity", 5),
		}),

		//Basedata
		basedataTable: table.New([]table.Column{
			table.NewColumn(columnKeyBasedata1, "EAN", 16),
			table.NewColumn(columnKeyBasedata2, "Name", 25),
			table.NewColumn(columnKeyBasedata3, "Price", 6),
		}),
		selectedOptionBasedata: 0,

		textInput: ti,
	}
}

func (m modelUI) Init() tea.Cmd {
	m.spinner.Tick()
	//m.scanInputStore.Blink()
	//return m.spinner.Tick //return nil
	//return textinput.Blink //return nil
	return textinput.Blink

}

func (m modelUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	//Key bindings for tab store
	switch m.selectedTab {
	case 0: //Store
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				//gets scanned input and puts it into cart
				query := strings.TrimSpace(m.textInput.Value())
				scanItem(query)
				m.textInput.Reset()
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
	case 1: //Warehouse
		switch msg := msg.(type) {

		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				query := strings.TrimSpace(m.textInput.Value())
				//query := truncate.Truncate(strings.TrimSpace(m.textInput.Value()), 14, "", truncate.PositionEnd) //	"github.com/aquilax/truncate"

				m.textInput.Reset()
				m.lastStocked = mapBProductToAPIProduct(database.GetProductByEan(query))

				//if ean available -> stock, if not failure
				var p model.APIProduct
				if m.lastStocked != p {
					warehouse.StockProduct(m.lastStocked.EAN)
					statusWarehouse = "alert"
				} else {
					statusWarehouse = "failure"
				}
			case " ":
				if statusWarehouse == "alert" || statusWarehouse == "failure" {
					statusWarehouse = "scan"
				}
			}
		}

	case 3:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				query := strings.TrimSpace(m.textInput.Value())
				m.textInput.Placeholder = ""

				if statusBasedata == "addEAN" {
					m.newProduct.EAN = query
					//m.newProduct.EAN = truncate.Truncate(query, 14, "", truncate.PositionEnd) //	"github.com/aquilax/truncate"

					m.textInput.Reset()
					if database.GetProductByEan(m.newProduct.EAN).ProductID != 0 {
						m.textInput.Placeholder = database.GetProductByEan(m.newProduct.EAN).Name
					}
					statusBasedata = "addName"
				} else if statusBasedata == "addName" {
					m.newProduct.Name = strings.TrimSpace(m.textInput.Value())
					m.textInput.Reset()
					if database.GetProductByEan(m.newProduct.EAN).ProductID != 0 {
						m.textInput.Placeholder = fmt.Sprintf("%.2f€", database.GetProductByEan(m.newProduct.EAN).Price)
					}
					statusBasedata = "addPrice"
				} else if statusBasedata == "addPrice" {

					query = strings.TrimSpace(m.textInput.Value())
					m.textInput.Reset()

					if s, err := strconv.ParseFloat(query, 64); err == nil {
						m.newProduct.Price = s
						basedata.AddProduct(m.newProduct.EAN, m.newProduct.Name, m.newProduct.Price)
					}

					statusBasedata = "table"
				}
				if statusBasedata == "delete" {
					query := strings.TrimSpace(m.textInput.Value())
					//query := truncate.Truncate(strings.TrimSpace(m.textInput.Value()), 14, "", truncate.PositionEnd)
					if database.GetProductByEan(query).ProductID != 0 {
						basedata.RemoveProduct(query)
					}
					m.textInput.Reset()
					statusBasedata = "table"
				}

			case " ":
				if statusBasedata == "table" {
					switch m.selectedOptionBasedata {
					case 0:
						statusBasedata = "addEAN"

					case 1:
						statusBasedata = "delete"
					}
				}
			case "right":
				if m.selectedOptionBasedata < 1 {
					m.selectedOptionBasedata += 1
				} else {
					m.selectedOptionBasedata = 0
				}
			case "left":
				if m.selectedOptionBasedata > 0 {
					m.selectedOptionBasedata -= 1
				} else {
					m.selectedOptionBasedata = 1
				}
			}
		}
	}

	switch msg := msg.(type) {
	// Case Key Press for everytime
	case tea.KeyMsg:

		// Switch between presses
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "esc":
			return m, tea.Quit

		case "tab":
			//change tab
			if m.selectedTab < len(m.tabs)-1 {
				m.selectedTab++
			} else {
				m.selectedTab = 0
			}

			//Cleanup
			if len(store.GetCart()) > 0 {
				store.ClearCart()
			}
			statusStore = "scan"
			statusWarehouse = "scan"
			statusBasedata = "table"

		}

		/*
			if statusStore == "scan" {
				//m.typing = true
				m.textInput, cmd = m.textInput.Update(msg)
				return m, cmd
			}
		*/

		/*
			default:
				var cmd tea.Cmd
				m.spinner, cmd = m.spinner.Update(msg)
				return m, cmd
		*/
	}

	//m.scanInputStore, cmd = m.textInput.Update(msg)

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	//return m, nil

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd

}

func (m modelUI) View() string {
	//physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
	doc := strings.Builder{}

	if m.selectedTab == 0 {
		return StoreUI(m)
	}
	if m.selectedTab == 1 {
		return WarehouseUI(m)
	}

	if m.selectedTab == 2 {
		return ReportingUI(m)
	}

	if m.selectedTab == 3 {
		return BasedataUI(m)
	}

	return docStyle.Render(doc.String())
}

//Starting Point for main function
func StartUI() {

	t := textinput.NewModel()
	t.Focus()

	p := tea.NewProgram(initialModel())

	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}

//Helper Functions
func mapBProductToAPIProduct(input model.BProduct) model.APIProduct {
	return model.APIProduct{EAN: input.EAN, Name: input.Name, Price: input.Price}
}
