package app

import "github.com/shopspring/decimal"

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
