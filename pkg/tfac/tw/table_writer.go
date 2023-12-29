package tw

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func GetTable(cols []string, rows [][]string) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader(cols)
	table.AppendBulk(rows)

	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)

	return table
}
