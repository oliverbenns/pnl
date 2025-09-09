package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/oliverbenns/pnl/internal/service/app"
)

type Trade struct {
	Symbol       string    `json:"symbol"`
	SecurityType string    `json:"security_type"`
	Quantity     string    `json:"quantity"`
	Price        string    `json:"price"`
	Date         time.Time `json:"date"`
	Side         string    `json:"side"`
}

func (t *Trade) toAppTrade() app.Trade {
	return app.Trade{
		Symbol:   t.Symbol,
		Quantity: t.Quantity,
		Price:    t.Price,
		Date:     t.Date,
		Side:     t.Side,
	}
}

func getTrades() ([]app.Trade, error) {
	data, err := os.ReadFile("assets/trades.json")
	if err != nil {
		return nil, err
	}

	var trades []Trade
	err = json.Unmarshal(data, &trades)
	if err != nil {
		return nil, err
	}

	var appTrades []app.Trade
	for _, trade := range trades {
		appTrades = append(appTrades, trade.toAppTrade())
	}

	return appTrades, nil
}

type Price struct {
	Symbol string `json:"symbol"`
	Value  string `json:"value"`
}

func (p *Price) toAppPrice() app.Price {
	return app.Price{
		Symbol: p.Symbol,
		Value:  p.Value,
	}
}

func getPrices() ([]app.Price, error) {
	data, err := os.ReadFile("assets/prices.json")
	if err != nil {
		return nil, err
	}

	var prices []Price
	err = json.Unmarshal(data, &prices)
	if err != nil {
		return nil, err
	}

	var appPrices []app.Price
	for _, price := range prices {
		appPrices = append(appPrices, price.toAppPrice())
	}

	return appPrices, nil
}
