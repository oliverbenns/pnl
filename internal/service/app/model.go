package app

import "time"

type Trade struct {
	Symbol   string
	Quantity string
	Price    string
	Date     time.Time
	Side     string
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
