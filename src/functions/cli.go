package functions

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func CreateTable() *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"💱 Currency", "💸 Value"})
	table.SetColWidth(50)
	return table
}
