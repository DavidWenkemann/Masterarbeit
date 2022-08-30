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

const (
	columnKeyReportingSelled1 = "Verkauft am"
	columnKeyReportingSelled2 = "ID"
	columnKeyReportingSelled3 = "Name"
	columnKeyReportingSelled4 = "Preis"

	layoutISO = "2006-01-02"
	Stamp     = "Jan _2 15:04:05"
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

//Generates Rows for Cart Table
func generateRowsFromSelledItems() []table.Row {
	rows := []table.Row{}
	selledItems := reporting.GetSelledItems()
	gesamtverdienst := 0.0

	for i := 0; i <= len(selledItems)-1; i++ {
		row := table.NewRow(table.RowData{
			columnKeyReportingSelled1: selledItems[i].SellingDate.Format(Stamp),
			columnKeyReportingSelled2: selledItems[i].ItemID,
			columnKeyReportingSelled3: reporting.GetProductByID(selledItems[i].ProductID).Name,
			columnKeyReportingSelled4: fmt.Sprintf("%.2f€", reporting.GetProductByID(selledItems[i].ProductID).Price),
		})
		gesamtverdienst += reporting.GetProductByID(selledItems[i].ProductID).Price
		rows = append(rows, row)
	}

	row := table.NewRow(table.RowData{
		columnKeyReportingSelled3: "Gesamtverdienst:",
		columnKeyReportingSelled4: fmt.Sprintf("%.2f€", gesamtverdienst),
	})

	rows = append(rows, row)

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

	if m.selectedTable {
		desc := lipgloss.JoinVertical(lipgloss.Left,
			descStyle.Render("Reporting View"+divider+"change with arrows"+divider+"All Items in Stock"),
		)
		row := lipgloss.Place(width, 3,
			lipgloss.Center, lipgloss.Center,
			lipgloss.JoinHorizontal(lipgloss.Top, desc),
		)

		doc.WriteString(row + "\n")

		m.reportingTable = m.reportingTable.WithRows(generateRowsFromStock())
	} else {
		desc := lipgloss.JoinVertical(lipgloss.Left,
			descStyle.Render("Reporting View"+divider+"change with arrows"+divider+"All Selled Items"),
		)
		row := lipgloss.Place(width, 3,
			lipgloss.Center, lipgloss.Center,
			lipgloss.JoinHorizontal(lipgloss.Top, desc),
		)

		doc.WriteString(row + "\n")

		m.reportingTable = m.reportingTableSelled.WithRows(generateRowsFromSelledItems())
	}

	doc.WriteString(m.reportingTable.View())

	return docStyle.Render(doc.String())
}
