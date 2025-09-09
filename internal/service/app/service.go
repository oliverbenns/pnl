package app

import (
	"context"
	"strconv"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Run(ctx context.Context, trades []Trade, prices []Price) ([]Position, error) {
	positions := make([]Position, 0)
	for _, trade := range trades {
		positions = append(positions, Position{
			Symbol:        trade.Symbol,
			Quantity:      strconv.FormatFloat(trade.Quantity, 'f', -1, 64),
			AvgCostBasis:  strconv.FormatFloat(trade.Price, 'f', -1, 64),
			CurrentPrice:  strconv.FormatFloat(trade.Price, 'f', -1, 64),
			MarketValue:   strconv.FormatFloat(trade.Quantity*trade.Price, 'f', -1, 64),
			UnrealizedPnL: strconv.FormatFloat(0, 'f', -1, 64),
			RealizedPnL:   strconv.FormatFloat(0, 'f', -1, 64),
			TotalPnL:      strconv.FormatFloat(0, 'f', -1, 64),
		})
	}

	return positions, nil
}
