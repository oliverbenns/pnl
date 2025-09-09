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
	priceMap := make(map[string]string)
	for _, price := range prices {
		priceMap[price.Symbol] = price.Value
	}

	groupedTrades := s.groupTradesBySymbol(trades)

	positions := make([]Position, 0)
	for _, trades := range groupedTrades {
		position, err := calculatePosition(trades, priceMap)
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

func calculatePosition(trades []Trade, priceMap map[string]string) (Position, error) {
	qty := decimal.NewFromInt(0)

	position := Position{}
	for _, trade := range trades {
		tradeQty, err := decimal.NewFromString(trade.Quantity)
		if err != nil {
			return Position{}, err
		}

		if trade.Side == "buy" {
			qty = qty.Add(tradeQty)
		} else {
			qty = qty.Sub(tradeQty)
		}

		// quantity := 0.0
		// avgCostBasis := 0.0
		// position.MarketValue = 0
		// position.UnrealizedPnL = 0
		// position.RealizedPnL = 0
		// position.TotalPnL = 0
	}

	position.Symbol = trades[0].Symbol
	position.CurrentPrice = priceMap[trades[0].Symbol]
	position.Quantity = qty.String()
	price, err := decimal.NewFromString(priceMap[trades[0].Symbol])
	if err != nil {
		return Position{}, err
	}
	position.MarketValue = qty.Mul(price).String()

	return position, nil
}
