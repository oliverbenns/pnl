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
