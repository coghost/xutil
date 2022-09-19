package xutil

import (
	"encoding/csv"
	"os"

	"github.com/olekukonko/tablewriter"
)

func Tablelize(s *csv.Reader) {
	table, _ := tablewriter.NewCSVReader(os.Stdout, s, true)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.FgHiGreenColor, tablewriter.Bold},
		tablewriter.Colors{tablewriter.FgHiGreenColor, tablewriter.Bold},
		tablewriter.Colors{tablewriter.FgHiGreenColor, tablewriter.Bold},
	)

	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiCyanColor},
	)
	table.Render()
}

func PlainTablelize(s *csv.Reader) {
	table, _ := tablewriter.NewCSVReader(os.Stdout, s, true)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render()
}

func SliceToTable(header *[]string, data *[][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(*header)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(*data) // Add Bulk Data
	table.Render()
}
