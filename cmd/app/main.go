package main

import (
	"context"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/oliverbenns/pnl/internal/service/app"
)

func main() {
	ctx := context.Background()

	svc := app.NewService()

	trades, err := getTrades()
	if err != nil {
		log.Fatalf("Failed to get trades: %v", err)
	}

	prices, err := getPrices()
	if err != nil {
		log.Fatalf("Failed to get prices: %v", err)
	}

	positions, err := svc.Run(ctx, trades, prices)
	if err != nil {
		log.Fatalf("Failed to run service: %v", err)
	}

	table := tablewriter.NewWriter(os.Stdout)

	table.Header([]string{"Symbol", "Quantity", "Avg Cost Basis", "Current Price", "Market Value", "Unrealized PNL", "Realized PNL", "Total PNL"})

	for _, position := range positions {
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
