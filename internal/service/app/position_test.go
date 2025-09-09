package app

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalculatePosition(t *testing.T) {
	t.Run("calculates long position", func(t *testing.T) {
		trades := []Trade{
			{Symbol: "BTC-USD", Quantity: "2", Price: "100", Side: "buy"},
		}
		price := decimal.NewFromInt(101)

		position, err := calculatePosition(trades, price)
		require.NoError(t, err)

		assert.Equal(t, "BTC-USD", position.Symbol)
		assert.Equal(t, "2", position.Quantity)
		assert.Equal(t, "200", position.CostBasis)
		assert.Equal(t, "100", position.AvgPrice)
		assert.Equal(t, "101", position.CurrentPrice)
		assert.Equal(t, "202", position.MarketValue)
	})

	t.Run("calculates short position", func(t *testing.T) {
		trades := []Trade{
			{Symbol: "BTC-USD", Quantity: "2", Price: "100", Side: "sell"},
		}
		price := decimal.NewFromInt(101)

		position, err := calculatePosition(trades, price)
		require.NoError(t, err)

		assert.Equal(t, "BTC-USD", position.Symbol)
		assert.Equal(t, "-2", position.Quantity)
		assert.Equal(t, "-200", position.CostBasis)
		assert.Equal(t, "100", position.AvgPrice)
		assert.Equal(t, "101", position.CurrentPrice)
		assert.Equal(t, "-202", position.MarketValue)
	})

}
