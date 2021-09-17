package talib

import "github.com/shopspring/decimal"

func (series *TimeSeries) OpenPriceIndicator() Indicator {
	return series.NewIndicator(OPEN, func(b Bar) decimal.Decimal { return b.OpenPrice() })
}
func (series *TimeSeries) ClosePriceIndicator() Indicator {
	return series.NewIndicator(CLOSE, func(b Bar) decimal.Decimal { return b.ClosePrice() })
}
func (series *TimeSeries) HighPriceIndicator() Indicator {
	return series.NewIndicator(HIGH, func(b Bar) decimal.Decimal { return b.HighPrice() })
}
func (series *TimeSeries) LowPriceIndicator() Indicator {
	return series.NewIndicator(LOW, func(b Bar) decimal.Decimal { return b.LowPrice() })
}
func (series *TimeSeries) VolumeIndicator() Indicator {
	return series.NewIndicator(VOLUME, func(b Bar) decimal.Decimal { return b.Volume() })
}
func (series *TimeSeries) MedianPriceIndicator() Indicator {
	id := string(MEDIAN)
	low := series.LowPriceIndicator()
	high := series.HighPriceIndicator()
	return series.LoadOrStore(id, func() Indicator {
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
	id := string(TYPICAL)
	num := decimal.NewFromInt(3)
	low := series.LowPriceIndicator()
	high := series.HighPriceIndicator()
	close := series.ClosePriceIndicator()
	return series.LoadOrStore(id, func() Indicator {
		typical := NewCacheFrom(series)
		typical.calculate = func(offset uint64) decimal.Decimal {
			return decimal.Sum(high.Offset(offset), low.Offset(offset), close.Offset(offset)).Div(num)
		}
		return typical
	})
}
