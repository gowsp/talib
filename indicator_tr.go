package talib

import "github.com/shopspring/decimal"

// Tr True range
func (series *TimeSeries) Tr() Indicator {
	id := string(TR)
	low := series.LowPriceIndicator()
	high := series.HighPriceIndicator()
	close := series.ClosePriceIndicator()
	return series.LoadOrStore(id, func() Indicator {
		trIndicator := NewCacheFrom(series)
		trIndicator.calculate = func(offset uint64) decimal.Decimal {
			if trIndicator.OutOfBounds(offset + 1) {
				return decimal.Zero
			}
			lowPrice := low.Offset(offset)
			highPrice := high.Offset(offset)
			beforeClosePrice := close.Offset(offset + 1)
			return decimal.Max(
				highPrice.Sub(lowPrice),
				highPrice.Sub(beforeClosePrice).Abs(),
				lowPrice.Sub(beforeClosePrice).Abs())
		}
		return trIndicator
	})
}
