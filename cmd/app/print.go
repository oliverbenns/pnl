package main

import (
	"os"
	"sort"

	"github.com/olekukonko/tablewriter"
	"github.com/oliverbenns/pnl/internal/service/app"
)

func printPositions(positions []app.Position) {
	positionsCopy := make([]app.Position, len(positions))
	copy(positionsCopy, positions)

	sort.Slice(positionsCopy, func(i, j int) bool {
		return positionsCopy[i].Symbol < positionsCopy[j].Symbol
	})

	table := tablewriter.NewWriter(os.Stdout)

	table.Header([]string{"Symbol", "Quantity", "Avg Cost Basis", "Current Price", "Market Value", "Unrealized PNL", "Realized PNL", "Total PNL"})

	for _, position := range positionsCopy {
		row := positionToTableRow(position)
		table.Append(row)
	}

	table.Render()
}

func positionToTableRow(position app.Position) []string {
	columns := []string{
		position.Symbol,
		position.Quantity,
		position.AvgCostBasis,
		position.CurrentPrice,
		position.MarketValue,
		position.UnrealizedPnL,
		position.RealizedPnL,
		position.TotalPnL,
	}

	return columns
}
