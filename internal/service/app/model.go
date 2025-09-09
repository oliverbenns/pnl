package app

import "time"

type Trade struct {
	Symbol   string
	Quantity float64
	Price    float64
	Date     time.Time
}

type Position struct {
	Symbol        string
	Quantity      string
	AvgCostBasis  string
	CurrentPrice  string
	MarketValue   string
	UnrealizedPnL string
	RealizedPnL   string
	TotalPnL      string
}

type Price struct {
	Symbol string
	Value  string
}
