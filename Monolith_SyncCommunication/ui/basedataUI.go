//handels UI for basedata module

package userinterface

import (
	"fmt"
	"strings"

	"github.com/DavidWenkemann/Masterarbeit/Monolith_SyncCommunication/basedata"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

const (
	columnKeyBasedata1 = "Ean"
	columnKeyBasedata2 = "Name"
	columnKeyBasedata3 = "Preis"
)

//Generates Rows for Cart Table
func generateRows() []table.Row {
	rows := []table.Row{}

	for i := 0; i <= len(basedata.GetAllProducts())-1; i++ {
		row := table.NewRow(table.RowData{
			columnKeyBasedata1: basedata.GetAllProducts()[i].EAN,
			columnKeyBasedata2: basedata.GetAllProducts()[i].Name,
			columnKeyBasedata3: fmt.Sprintf("%.2f€", basedata.GetAllProducts()[i].Price),
		})
		rows = append(rows, row)
	}
	return rows
}

func BasedataUI(m modelUI) string {

	doc := strings.Builder{}

	{
		// Tabs on Top
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

	if statusBasedata == "table" {
		//BasedataTable
		m.basedataTable = m.basedataTable.WithRows(generateRows())
		//doc.WriteString(m.basedataTable.View())

		var buttonsText [3]string
		buttonsText[0] = "Add Product"
		buttonsText[1] = "Delete Product"

		var buttons [2]string
		buttons[0] = buttonStyle.Render(buttonsText[0])
		buttons[1] = buttonStyle.Render(buttonsText[1])
		buttons[m.selectedOptionBasedata] = activeButtonStyle.Render(buttonsText[m.selectedOptionBasedata])
		combinedbuttons := lipgloss.JoinHorizontal(lipgloss.Top, buttons[0], buttons[1])
		ui := lipgloss.JoinVertical(lipgloss.Center, m.basedataTable.View(), combinedbuttons)

		doc.WriteString(ui + "\n\n")
	}

	if statusBasedata == "addEAN" {
		alert := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render(fmt.Sprintf("Scan new EAN: \n%s", m.textInput.View()))

		ui := lipgloss.JoinVertical(lipgloss.Center, alert)

		dialog := lipgloss.Place(width, 9,
			lipgloss.Center, lipgloss.Center,
			dialogBoxStyle.Render(ui),
			lipgloss.WithWhitespaceChars(""),
			lipgloss.WithWhitespaceForeground(subtle),
		)

		doc.WriteString(dialog + "\n\n")

	}
	if statusBasedata == "addName" {
		alert := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render(fmt.Sprintf("Type in Name: \n%s", m.textInput.View()))

		ui := lipgloss.JoinVertical(lipgloss.Center, alert)

		dialog := lipgloss.Place(width, 9,
			lipgloss.Center, lipgloss.Center,
			dialogBoxStyle.Render(ui),
			lipgloss.WithWhitespaceChars(""),
			lipgloss.WithWhitespaceForeground(subtle),
		)

		doc.WriteString(dialog + "\n\n")

	}
	if statusBasedata == "addPrice" {
		alert := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render(fmt.Sprintf("Type in Price for %s: \n%s", m.newProduct.Name, m.textInput.View()))

		ui := lipgloss.JoinVertical(lipgloss.Center, alert)

		dialog := lipgloss.Place(width, 9,
			lipgloss.Center, lipgloss.Center,
			dialogBoxStyle.Render(ui),
			lipgloss.WithWhitespaceChars(""),
			lipgloss.WithWhitespaceForeground(subtle),
		)

		doc.WriteString(dialog + "\n\n")

	}

	if statusBasedata == "delete" {
		alert := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render(fmt.Sprintf("Scan EAN of Product you want delete: \n%s", m.textInput.View()))

		ui := lipgloss.JoinVertical(lipgloss.Center, alert)

		dialog := lipgloss.Place(width, 9,
			lipgloss.Center, lipgloss.Center,
			dialogBoxStyle.Render(ui),
			lipgloss.WithWhitespaceChars(""),
			lipgloss.WithWhitespaceForeground(subtle),
		)

		doc.WriteString(dialog + "\n\n")

	}

	return docStyle.Render(doc.String())
}
