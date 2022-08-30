//handels UI for reporting module

package userinterface

import (
	"fmt"
	"strings"

	"github.com/DavidWenkemann/Masterarbeit/Monolith_SyncCommunication/reporting"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

const (
	columnKeyReporting1 = "Ean"
	columnKeyReporting2 = "Name"
	columnKeyReporting3 = "Preis"
	columnKeyReporting4 = "Anzahl"
)

//Generates Rows for Cart Table
func generateRowsFromStock() []table.Row {
	rows := []table.Row{}

	for i := 0; i <= len(reporting.GetAllProducts())-1; i++ {
		row := table.NewRow(table.RowData{
			columnKeyReporting1: reporting.GetAllProducts()[i].EAN,
			columnKeyReporting2: reporting.GetAllProducts()[i].Name,
			columnKeyReporting3: fmt.Sprintf("%.2f€", reporting.GetAllProducts()[i].Price),
			columnKeyReporting4: reporting.GetItemsInStockByEan(reporting.GetAllProducts()[i].EAN),
		})
		rows = append(rows, row)
	}
	return rows
}

func ReportingUI(m modelUI) string {

	doc := strings.Builder{}

	{
		// Tabs on Top
		row := lipgloss.JoinHorizontal(
			lipgloss.Top,
			tab.Render("Store"),
			tab.Render("Warehouse"),
			activeTab.Render("Reporting"),
			tab.Render("Basedata"),
		)

		gap := tabGap.Render(strings.Repeat(" ", max(0, width-lipgloss.Width(row)-2)))
		row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)
		doc.WriteString(row + "\n")
	}

	//ReportingTable

	desc := lipgloss.JoinVertical(lipgloss.Left,
		descStyle.Render("Reporting View"+divider+"All Items in Stock"),
	)
	row := lipgloss.Place(width, 3,
		lipgloss.Center, lipgloss.Center,
		lipgloss.JoinHorizontal(lipgloss.Top, desc),
	)

	doc.WriteString(row + "\n")

	m.reportingTable = m.reportingTable.WithRows(generateRowsFromStock())
	doc.WriteString(m.reportingTable.View())

	return docStyle.Render(doc.String())
}
