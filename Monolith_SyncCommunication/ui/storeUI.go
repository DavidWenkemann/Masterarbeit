package userinterface

import (
	"fmt"
	"strings"

	"github.com/DavidWenkemann/Masterarbeit/Monolith_SyncCommunication/model"
	"github.com/DavidWenkemann/Masterarbeit/Monolith_SyncCommunication/store"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

const (
	columnKeyName  = "Produkt"
	columnKeyPrice = "Preis"
)

//Generates Rows for Cart Table
func generateRowsFromCart(refreshedcart []model.APIItem, m modelUI) []table.Row {
	rows := []table.Row{}

	//only for testing
	/*
		var products = []model.APIItem{
			{Product: model.APIProduct(database.GetProductByID(1)), ItemID: "1", ReceivingDate: time.Now()},
			{Product: model.APIProduct(database.GetProductByID(2)), ItemID: "2", ReceivingDate: time.Now()},
			{Product: model.APIProduct(database.GetProductByID(1)), ItemID: "3", ReceivingDate: time.Now()},
		}
		refreshedcart = products
	*/
	for i := 0; i <= len(refreshedcart)-1; i++ {
		row := table.NewRow(table.RowData{
			columnKeyName:  refreshedcart[i].Product.Name,
			columnKeyPrice: fmt.Sprintf("%.2f€", refreshedcart[i].Product.Price),
		})
		rows = append(rows, row)
	}

	lastrow := table.NewRow(table.RowData{
		//columnKeyName:  fmt.Sprintf(" %s", m.scanInputStore.View()),
		columnKeyName: fmt.Sprintf(" %s", m.textInput.View()),

		columnKeyPrice: "",
	})
	rows = append(rows, lastrow)

	return rows
}

func StoreUI(m modelUI) string {

	doc := strings.Builder{}

	m.textInput.SetCursorMode(2)
	/*
		// Status bar
		{
			w := lipgloss.Width

			statusKey := fishCakeStyle.Render("Master Thesis Wenkemann")
			//encoding := encodingStyle.Render("UTF-8")
			fishCake := fishCakeStyle.Render(fmt.Sprintf(time.Now().Format("15:04")))

			statusVal := statusText.Copy().
				Width(width - w(statusKey) - w(fishCake)).
				Render("")

			bar := lipgloss.JoinHorizontal(lipgloss.Top,
				statusKey,
				statusVal,
				fishCake,
			)

			doc.WriteString(statusBarStyle.Width(width).Render(bar) + "\n\n")
		}
	*/

	{
		// Tabs on Top
		row := lipgloss.JoinHorizontal(
			lipgloss.Top,
			activeTab.Render("Store"),
			tab.Render("Warehouse"),
			tab.Render("Reporting"),
			tab.Render("Basedata"),
		)

		gap := tabGap.Render(strings.Repeat(" ", max(0, width-lipgloss.Width(row)-2)))
		row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)
		doc.WriteString(row + "\n")
	}

	//Title
	{

		titleStyle = lipgloss.NewStyle().
			MarginLeft(1).
			MarginRight(5).
			Padding(0, 1).
			Italic(true).
			Foreground(lipgloss.Color("#FFF7DB")).
			SetString("Store")

		desc := lipgloss.JoinVertical(lipgloss.Left,
			descStyle.Render("Welcome to the Store"),
			infoStyle.Render("Scan for add"+divider+"Spacebar for Checkout"),
		)
		row := lipgloss.Place(width, 3,
			lipgloss.Center, lipgloss.Center,
			lipgloss.JoinHorizontal(lipgloss.Top, desc),
		)

		doc.WriteString(row + "\n")
	}

	if statusStore == "scan" {
		//Cart
		m.cartTable = m.cartTable.WithRows(generateRowsFromCart(store.GetCart(), m))

		var tableStyle = lipgloss.NewStyle().
			MarginLeft(1)

		carttable := lipgloss.Place(width, 9,
			lipgloss.Center, lipgloss.Center,
			tableStyle.Render(m.cartTable.View()),
		)

		doc.WriteString(carttable + "\n\n")

	} else if statusStore == "checkout" {
		// Dialog
		var buttonsText [3]string
		buttonsText[0] = "Pay now"
		buttonsText[1] = "Keep Scannin´"
		buttonsText[2] = "Delete Cart"

		var buttons [3]string
		buttons[0] = buttonStyle.Render(buttonsText[0])
		buttons[1] = buttonStyle.Render(buttonsText[1])
		buttons[2] = buttonStyle.Render(buttonsText[2])

		buttons[m.selectedOption] = activeButtonStyle.Render(buttonsText[m.selectedOption])

		question := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render(fmt.Sprintf("You want to checkout (%.2f€) now?", store.GetPriceOfCart()))
		combinedbuttons := lipgloss.JoinHorizontal(lipgloss.Top, buttons[0], buttons[1], buttons[2])
		ui := lipgloss.JoinVertical(lipgloss.Center, question, combinedbuttons)

		dialog := lipgloss.Place(width, 9,
			lipgloss.Center, lipgloss.Center,
			dialogBoxStyle.Render(ui),
			lipgloss.WithWhitespaceChars(""),
			lipgloss.WithWhitespaceForeground(subtle),
		)

		doc.WriteString(dialog + "\n\n")
	}

	//Send Back to UI
	return docStyle.Render(doc.String())
}

func scanItem(itemID string) {
	store.AddToCart(itemID)
}
