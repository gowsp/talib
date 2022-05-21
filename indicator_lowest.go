package talib

import (
	"github.com/gowsp/talib/internal"
	"github.com/shopspring/decimal"
)

// Lowest Lowest value indicator
func Lowest(i Indicator, length uint64) Indicator {
	id := internal.LOWEST.Id(length)
	return i.LoadOrStore(id, func() Indicator {
		lowest := NewCachedIndicator(i)
		lowest.calculate = func(offset uint64) decimal.Decimal {
			end := lowest.BarSeries().Size()
			if offset+length < end {
				end = offset + length
			}
			val := i.Offset(offset)
			for v := offset + 1; v < end; v++ {
				tmp := i.Offset(v)
				if tmp.LessThan(val) {
					val = tmp
				}
			}
			return val
		}
		return lowest
	})
}
