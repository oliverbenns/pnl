package main

import (
	"context"
	"log"

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

	printPositions(positions)
}
