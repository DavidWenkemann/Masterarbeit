//handels UI for warehouse module

package userinterface

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func WarehouseUI(m modelUI) string {

	doc := strings.Builder{}

	// Tabs on Top
	{
		row := lipgloss.JoinHorizontal(
			lipgloss.Top,
			tab.Render("Store"),
			activeTab.Render("Warehouse"),
			tab.Render("Reporting"),
			tab.Render("Basedata"),
		)

		gap := tabGap.Render(strings.Repeat(" ", max(0, width-lipgloss.Width(row)-2)))
		row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)
		doc.WriteString(row + "\n")
	}

	var alert string

	if statusWarehouse == "scan" {
		alert = lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render("Scan EAN to stock the bottle" + fmt.Sprintf(" %s", m.textInput.View()))
	} else if statusWarehouse == "alert" {
		alert = lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render(fmt.Sprintf("%s was added to stock", m.lastStocked.Name))
	} else if statusWarehouse == "failure" {
		alert = lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render(fmt.Sprintf("Could not stock item - try again! \n %s", m.spinner.View()))
	}

	ui := lipgloss.JoinVertical(lipgloss.Center, alert)

	dialog := lipgloss.Place(width, 9,
		lipgloss.Center, lipgloss.Center,
		dialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceChars(""),
		lipgloss.WithWhitespaceForeground(subtle),
	)

	doc.WriteString(dialog + "\n\n")

	return docStyle.Render(doc.String())

}
