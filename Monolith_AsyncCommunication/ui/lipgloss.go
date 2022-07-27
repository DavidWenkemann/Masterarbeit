package userinterface

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
	"golang.org/x/term"
)

const (
	// In real life situations we'd adjust the document to fit the width we've
	// detected. In the case of this example we're hardcoding the width, and
	// later using the detected width only to truncate in order to avoid jaggy
	// wrapping.
	width = 60

	columnWidth = 10
)

// Style definitions.
var (

	// General.

	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

	divider = lipgloss.NewStyle().
		SetString("•").
		Padding(0, 1).
		Foreground(subtle).
		String()

	url = lipgloss.NewStyle().Foreground(special).Render

	// Tabs.

	activeTabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}

	tabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}

	tab = lipgloss.NewStyle().
		Border(tabBorder, true).
		BorderForeground(highlight).
		Padding(0, 1)

	activeTab = tab.Copy().Border(activeTabBorder, true)

	tabGap = tab.Copy().
		BorderTop(false).
		BorderLeft(false).
		BorderRight(false)

	// Title.

	titleStyle = lipgloss.NewStyle().
			MarginLeft(1).
			MarginRight(5).
			Padding(0, 1).
			Italic(true).
			Foreground(lipgloss.Color("#FFF7DB")).
			SetString("Lip Gloss")

	descStyle = lipgloss.NewStyle().MarginTop(1)

	infoStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderTop(true).
			BorderForeground(subtle)

	// Dialog.

	dialogBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(1, 0).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)

	buttonStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFF7DB")).
			Background(lipgloss.Color("#888B7E")).
			Padding(0, 2).
			MarginTop(1).
			MarginRight(2)

	activeButtonStyle = buttonStyle.Copy().
				Foreground(lipgloss.Color("#FFF7DB")).
				Background(lipgloss.Color("#F25D94")).
				MarginRight(2).
				Underline(true)

	// List.

	list = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderForeground(subtle).
		MarginRight(2).
		Height(8).
		Width(columnWidth + 1)

	listHeader = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(subtle).
			MarginRight(2).
			Render

	listItem = lipgloss.NewStyle().PaddingLeft(2).Render

	checkMark = lipgloss.NewStyle().SetString("✓").
			Foreground(special).
			PaddingRight(1).
			String()

	listDone = func(s string) string {
		return checkMark + lipgloss.NewStyle().
			Strikethrough(true).
			Foreground(lipgloss.AdaptiveColor{Light: "#969B86", Dark: "#696969"}).
			Render(s)
	}

	// Paragraphs/History.

	historyStyle = lipgloss.NewStyle().
			Align(lipgloss.Left).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(highlight).
			Margin(1, 3, 0, 0).
			Padding(1, 2).
			Height(19).
			Width(columnWidth)

	// Status Bar.

	statusNugget = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Padding(0, 1)

	statusBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
			Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})

	statusStyle = lipgloss.NewStyle().
			Inherit(statusBarStyle).
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#FF5F87")).
			Padding(0, 1).
			MarginRight(1)

	encodingStyle = statusNugget.Copy().
			Background(lipgloss.Color("#A550DF")).
			Align(lipgloss.Right)

	statusText = lipgloss.NewStyle().Inherit(statusBarStyle)

	fishCakeStyle = statusNugget.Copy().Background(lipgloss.Color("#6124DF"))

	// Page.

	docStyle = lipgloss.NewStyle().Padding(1, 2, 1, 2)
)

func StartLipgloss() {
	physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
	doc := strings.Builder{}

	// Tabs
	{
		row := lipgloss.JoinHorizontal(
			lipgloss.Top,
			activeTab.Render("Lip Gloss"),
			tab.Render("Blush"),
			tab.Render("Eye Shadow"),
			tab.Render("Mascara"),
			tab.Render("Foundation"),
		)
		gap := tabGap.Render(strings.Repeat(" ", max(0, width-lipgloss.Width(row)-2)))
		row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)
		doc.WriteString(row + "\n\n")
	}

	// Title
	{
		var (
			colors = colorGrid(1, 5)
			title  strings.Builder
		)

		for i, v := range colors {
			const offset = 2
			c := lipgloss.Color(v[0])
			fmt.Fprint(&title, titleStyle.Copy().MarginLeft(i*offset).Background(c))
			if i < len(colors)-1 {
				title.WriteRune('\n')
			}
		}

		desc := lipgloss.JoinVertical(lipgloss.Left,
			descStyle.Render("Style Definitions for Nice Terminal Layouts"),
			infoStyle.Render("From Charm"+divider+url("https://github.com/charmbracelet/lipgloss")),
		)

		row := lipgloss.JoinHorizontal(lipgloss.Top, title.String(), desc)
		doc.WriteString(row + "\n\n")
	}

	// Dialog
	{
		okButton := activeButtonStyle.Render("Yes")
		cancelButton := buttonStyle.Render("Maybe")

		question := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render("Are you sure you want to eat marmalade?")
		buttons := lipgloss.JoinHorizontal(lipgloss.Top, okButton, cancelButton)
		ui := lipgloss.JoinVertical(lipgloss.Center, question, buttons)

		dialog := lipgloss.Place(width, 9,
			lipgloss.Center, lipgloss.Center,
			dialogBoxStyle.Render(ui),
			lipgloss.WithWhitespaceChars("猫咪"),
			lipgloss.WithWhitespaceForeground(subtle),
		)

		doc.WriteString(dialog + "\n\n")
	}

	// Color grid
	colors := func() string {
		colors := colorGrid(14, 8)

		b := strings.Builder{}
		for _, x := range colors {
			for _, y := range x {
				s := lipgloss.NewStyle().SetString("  ").Background(lipgloss.Color(y))
				b.WriteString(s.String())
			}
			b.WriteRune('\n')
		}

		return b.String()
	}()

	lists := lipgloss.JoinHorizontal(lipgloss.Top,
		list.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				listHeader("Citrus Fruits to Try"),
				listDone("Grapefruit"),
				listDone("Yuzu"),
				listItem("Citron"),
				listItem("Kumquat"),
				listItem("Pomelo"),
			),
		),
		list.Copy().Width(columnWidth).Render(
			lipgloss.JoinVertical(lipgloss.Left,
				listHeader("Actual Lip Gloss Vendors"),
				listItem("Glossier"),
				listItem("Claire‘s Boutique"),
				listDone("Nyx"),
				listItem("Mac"),
				listDone("Milk"),
			),
		),
	)

	doc.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, lists, colors))

	// Marmalade history
	{
		const (
			historyA = "The Romans learned from the Greeks that quinces slowly cooked with honey would “set” when cool. The Apicius gives a recipe for preserving whole quinces, stems and leaves attached, in a bath of honey diluted with defrutum: Roman marmalade. Preserves of quince and lemon appear (along with rose, apple, plum and pear) in the Book of ceremonies of the Byzantine Emperor Constantine VII Porphyrogennetos."
			historyB = "Medieval quince preserves, which went by the French name cotignac, produced in a clear version and a fruit pulp version, began to lose their medieval seasoning of spices in the 16th century. In the 17th century, La Varenne provided recipes for both thick and clear cotignac."
			historyC = "In 1524, Henry VIII, King of England, received a “box of marmalade” from Mr. Hull of Exeter. This was probably marmelada, a solid quince paste from Portugal, still made and sold in southern Europe today. It became a favourite treat of Anne Boleyn and her ladies in waiting."
		)

		doc.WriteString(lipgloss.JoinHorizontal(
			lipgloss.Top,
			historyStyle.Copy().Align(lipgloss.Right).Render(historyA),
			historyStyle.Copy().Align(lipgloss.Center).Render(historyB),
			historyStyle.Copy().MarginRight(0).Render(historyC),
		))

		doc.WriteString("\n\n")
	}

	// Status bar
	{
		w := lipgloss.Width

		statusKey := statusStyle.Render("STATUS")
		encoding := encodingStyle.Render("UTF-8")
		fishCake := fishCakeStyle.Render("🍥 Fish Cake")
		statusVal := statusText.Copy().
			Width(width - w(statusKey) - w(encoding) - w(fishCake)).
			Render("Ravishing")

		bar := lipgloss.JoinHorizontal(lipgloss.Top,
			statusKey,
			statusVal,
			encoding,
			fishCake,
		)

		doc.WriteString(statusBarStyle.Width(width).Render(bar))
	}

	if physicalWidth > 0 {
		docStyle = docStyle.MaxWidth(physicalWidth)
	}

	// Okay, let's print it
	fmt.Println(docStyle.Render(doc.String()))
}

func colorGrid(xSteps, ySteps int) [][]string {
	x0y0, _ := colorful.Hex("#F25D94")
	x1y0, _ := colorful.Hex("#EDFF82")
	x0y1, _ := colorful.Hex("#643AFF")
	x1y1, _ := colorful.Hex("#14F9D5")

	x0 := make([]colorful.Color, ySteps)
	for i := range x0 {
		x0[i] = x0y0.BlendLuv(x0y1, float64(i)/float64(ySteps))
	}

	x1 := make([]colorful.Color, ySteps)
	for i := range x1 {
		x1[i] = x1y0.BlendLuv(x1y1, float64(i)/float64(ySteps))
	}

	grid := make([][]string, ySteps)
	for x := 0; x < ySteps; x++ {
		y0 := x0[x]
		grid[x] = make([]string, xSteps)
		for y := 0; y < xSteps; y++ {
			grid[x][y] = y0.BlendLuv(x1[x], float64(y)/float64(xSteps)).Hex()
		}
	}

	return grid
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-table/table"
)

//Business/Persistenze Produkte

//Draw --> know cols --> map products to raw data -->  values to :  col/string

type value struct {
	col string
	val string
}

func NewIntVal(val int, col string) value {
	v := value{}
	v.col = col
	v.val = strconv.Itoa(val)
	return v
}
func NewStrVal(val string, col string) value {
	v := value{}
	v.col = col
	v.val = val
	return v
}

type dataRow struct {
	values []value
}

var bdb *baseDataBusiness = nil

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
	columnKeyEan   = "ean"
	columnKeyName  = "name"
	columnKeyPrice = "price"
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
			table.NewColumn(columnKeyPrice, "Price", 6),
		}),
		data:      bdb.GetAllProducts(),
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
		case "enter":

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
					bdb.RemoveProductByEAN(query)
					m.textInput.Reset()
					m.state = "overview"
				}
			}
			if m.typing && m.state == "addean" {
				query := strings.TrimSpace(m.textInput.Value())

				if query != "" && m.newEan == "" {
					m.newEan = query
					m.textInput.Reset()
					m.state = "addname"
				}
			}
			if m.typing && m.state == "addname" {
				query := strings.TrimSpace(m.textInput.Value())

				if query != "" {
					m.newName = query
					m.textInput.Reset()
					m.state = "addprice"
				}
			}
			if m.typing && m.state == "addprice" {
				query := strings.TrimSpace(m.textInput.Value())

				if query != "" {
					m.typing = false

					//Replace(query, ",", ".", 1)

					if p, err := strconv.ParseFloat(strings.Replace(query, ",", ".", 1), 64); err == nil {
						m.newPrice = p

						//fmt.Println(p) // 3.1415927410125732
					}

					//m.newPrice = query
					bdb.AddProduct(m.newEan, m.newName, m.newPrice)
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
		m.data = bdb.GetAllProducts()
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

			return fmt.Sprintf("Type in Price for %s: \n%s", m.newName, m.textInput.View())
		} else {
			s = "Status add: Failure\n\n"
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
			columnKeyEan:   refreshedproducts[i].EAN,
			columnKeyName:  refreshedproducts[i].Name,
			columnKeyPrice: fmt.Sprintf("%.2f€", refreshedproducts[i].Price),
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

*/
