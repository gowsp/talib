package talib

import (
	"github.com/gowsp/talib/internal"
	"github.com/shopspring/decimal"
)

// Sma function returns the moving average
func Sma(indicator Indicator, length uint64) Indicator {
	id := internal.SMA.Id(length)
	return indicator.LoadOrStore(id, func() Indicator {
		window := decimal.NewFromInt(int64(length))
		cache := NewCachedIndicator(indicator, func(i Indicator, offset uint64) decimal.Decimal {
			if i.OutOfBounds(offset + length) {
				size := indicator.BarSeries().Size()
				slice := make([]decimal.Decimal, size-offset)
				for i := offset; i < size; i++ {
					slice[i] = indicator.Offset(i)
				}
				return decimal.Avg(slice[0], slice[1:]...)
			}
			current := indicator.Offset(offset).Div(window)
			before := indicator.Offset(offset + length).Div(window)
			return i.Offset(offset + 1).Sub(before).Add(current)
		})
		return cache
	})
}

// Ema The ema function returns the exponentially weighted moving average.
// In ema weighting factors decrease exponentially
func Ema(indicator Indicator, length uint64) Indicator {
	id := internal.EMA.Id(length)
	return indicator.LoadOrStore(id, func() Indicator {
		alpha := TWO.Div(decimal.NewFromInt(int64(length)).Add(ONE))
		return BaseEma(indicator, length, alpha)
	})
}

// Rma Moving average used in RSI. It is the exponentially weighted moving average with alpha = 1 / length.
func Rma(indicator Indicator, length uint64) Indicator {
	id := internal.RMA.Id(length)
	return indicator.LoadOrStore(id, func() Indicator {
		alpha := ONE.Div(decimal.NewFromInt(int64(length)))
		return BaseEma(indicator, length, alpha)
	})
}
func BaseEma(indicator Indicator, length uint64, alpha decimal.Decimal) Indicator {
	ema := NewCachedIndicator(indicator, func(i Indicator, offset uint64) decimal.Decimal {
		if i.OutOfBounds(offset) {
			return indicator.Offset(offset)
		}
		prevValue := i.Offset(offset + 1)
		return indicator.Offset(offset).Sub(prevValue).Mul(alpha).Add(prevValue)
	})
	return ema

}
