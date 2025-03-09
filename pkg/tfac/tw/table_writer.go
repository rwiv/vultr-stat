package tw

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func GetTable(cols []string, rows [][]string) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader(cols)
	table.AppendBulk(rows)

	return table
}
