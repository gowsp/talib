package talib

import (
	"github.com/gowsp/talib/internal"
	"github.com/shopspring/decimal"
)

// Tr True range
func (series *TimeSeries) Tr() Indicator {
	low := series.LowPriceIndicator()
	high := series.HighPriceIndicator()
	close := series.ClosePriceIndicator()
	return series.LoadOrStore(internal.TR, func() Indicator {
		trIndicator := NewCacheFrom(series, func(i Indicator, offset uint64) decimal.Decimal {
			if i.OutOfBounds(offset + 1) {
				return decimal.Zero
			}
			lowPrice := low.Offset(offset)
			highPrice := high.Offset(offset)
			beforeClosePrice := close.Offset(offset + 1)
			return decimal.Max(
				highPrice.Sub(lowPrice),
				highPrice.Sub(beforeClosePrice).Abs(),
				lowPrice.Sub(beforeClosePrice).Abs())
		})
		return trIndicator
	})
}
