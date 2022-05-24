package talib

import (
	"github.com/gowsp/talib/internal"
	"github.com/shopspring/decimal"
)

// Highest Highest value indicator
func Highest(i Indicator, length uint64) Indicator {
	id := internal.HIGHEST.Id(length)
	return i.LoadOrStore(id, func() Indicator {
		highest := NewCachedIndicator(i, func(v Indicator, offset uint64) decimal.Decimal {
			end := v.BarSeries().Size()
			if offset+length < end {
				end = offset + length
			}
			val := i.Offset(offset)
			for v := offset + 1; v < end; v++ {
				tmp := i.Offset(v)
				if tmp.GreaterThan(val) {
					val = tmp
				}
			}
			return val
		})
		return highest
	})
}
