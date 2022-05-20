package talib

import (
	"github.com/gowsp/talib/internal"
	"github.com/shopspring/decimal"
)

func (series *TimeSeries) OpenPriceIndicator() Indicator {
	return series.NewIndicator(internal.OPEN, func(b Bar) decimal.Decimal { return b.OpenPrice() })
}
func (series *TimeSeries) ClosePriceIndicator() Indicator {
	return series.NewIndicator(internal.CLOSE, func(b Bar) decimal.Decimal { return b.ClosePrice() })
}
func (series *TimeSeries) HighPriceIndicator() Indicator {
	return series.NewIndicator(internal.HIGH, func(b Bar) decimal.Decimal { return b.HighPrice() })
}
func (series *TimeSeries) LowPriceIndicator() Indicator {
	return series.NewIndicator(internal.LOW, func(b Bar) decimal.Decimal { return b.LowPrice() })
}
func (series *TimeSeries) VolumeIndicator() Indicator {
	return series.NewIndicator(internal.VOLUME, func(b Bar) decimal.Decimal { return b.Volume() })
}
func (series *TimeSeries) MedianPriceIndicator() Indicator {
	low := series.LowPriceIndicator()
	high := series.HighPriceIndicator()
	return series.LoadOrStore(internal.MEDIAN, func() Indicator {
		median := NewCacheFrom(series)
		median.calculate = func(offset uint64) decimal.Decimal {
			lowPrice := low.Offset(offset)
			highPrice := high.Offset(offset)
			return highPrice.Add(lowPrice).Div(TWO)
		}
		return median
	})
}
func (series *TimeSeries) TypicalPriceIndicator() Indicator {
	num := decimal.NewFromInt(3)
	low := series.LowPriceIndicator()
	high := series.HighPriceIndicator()
	close := series.ClosePriceIndicator()
	return series.LoadOrStore(internal.TYPICAL, func() Indicator {
		typical := NewCacheFrom(series)
		typical.calculate = func(offset uint64) decimal.Decimal {
			return decimal.Sum(high.Offset(offset), low.Offset(offset), close.Offset(offset)).Div(num)
		}
		return typical
	})
}
