package xutil

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"
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

// https://github.com/jedib0t/go-pretty

func PrintArrayAsTable(header []string, rows [][]string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	row := table.Row{}
	for _, v := range header {
		row = append(row, v)
	}
	t.AppendHeader(row)

	total := 0
	for _, ar := range rows {
		if len(ar) == 0 {
			continue
		}

		row := table.Row{}
		for _, v := range ar {
			row = append(row, v)
		}
		t.AppendRow(row)
		total++
	}

	t.AppendFooter(table.Row{fmt.Sprintf("Total:%d", total), "", "", "", ""})

	t.SetStyle(table.StyleColoredDark)
	t.Render()
}

func PrintStrAsTable(header string, rows []string, sep string) {
	arr := StrToArrWithNonEmpty(header, sep)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	row := table.Row{}
	for _, v := range arr {
		row = append(row, v)
	}
	t.AppendHeader(row)

	total := 0
	for _, line := range rows {
		ar := StrToArrWithNonEmpty(line, sep)
		if len(ar) == 0 {
			continue
		}

		row := table.Row{}
		for _, v := range ar {
			row = append(row, v)
		}
		t.AppendRow(row)
		total++
	}

	t.AppendFooter(table.Row{fmt.Sprintf("Total:%d", total), "", "", "", ""})

	t.SetStyle(table.StyleColoredDark)
	t.Render()
}
