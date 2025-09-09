package app

import (
	"context"

	"github.com/shopspring/decimal"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Run(ctx context.Context, trades []Trade, prices []Price) ([]Position, error) {
	priceMap := make(map[string]decimal.Decimal)
	for _, price := range prices {
		var err error
		priceMap[price.Symbol], err = decimal.NewFromString(price.Value)
		if err != nil {
			return nil, err
		}
	}

	groupedTrades := s.groupTradesBySymbol(trades)

	positions := make([]Position, 0)
	for symbol, trades := range groupedTrades {
		price := priceMap[symbol]
		position, err := calculatePosition(trades, price)
		if err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}

	return positions, nil
}

func (s *Service) groupTradesBySymbol(trades []Trade) map[string][]Trade {
	groupedTrades := make(map[string][]Trade)
	for _, trade := range trades {
		groupedTrades[trade.Symbol] = append(groupedTrades[trade.Symbol], trade)
	}

	return groupedTrades
}

func calculatePosition(trades []Trade, price decimal.Decimal) (Position, error) {
	qty := decimal.NewFromInt(0)
	costBasis := decimal.NewFromInt(0)

	position := Position{}
	for _, trade := range trades {
		tradeQty, err := decimal.NewFromString(trade.Quantity)
		if err != nil {
			return Position{}, err
		}

		tradePrice, err := decimal.NewFromString(trade.Price)
		if err != nil {
			return Position{}, err
		}

		multiplier := getMultiplier(trade.Side)
		qtyWithMultiplier := tradeQty.Mul(multiplier)
		qty = qty.Add(qtyWithMultiplier)
		costBasis = costBasis.Add(qtyWithMultiplier.Mul(tradePrice))
	}

	position.Symbol = trades[0].Symbol
	position.CurrentPrice = price.String()
	position.Quantity = qty.String()
	position.MarketValue = qty.Mul(price).String()
	position.CostBasis = costBasis.String()
	position.AvgPrice = costBasis.Div(qty).String()

	return position, nil
}

func getMultiplier(side string) decimal.Decimal {
	if side == "buy" {
		return decimal.NewFromInt(1)
	}

	return decimal.NewFromInt(-1)
}
