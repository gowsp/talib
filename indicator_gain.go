package talib

import (
	"github.com/gowsp/talib/internal"
	"github.com/shopspring/decimal"
)

// Gain
func Gain(i Indicator) Indicator {
	return i.LoadOrStore(internal.GAIN, func() Indicator {
		gain := NewCachedIndicator(i, func(i Indicator, offset uint64) decimal.Decimal {
			if i.OutOfBounds(offset) {
				return decimal.Zero
			}
			cur := i.Offset(offset)
			pre := i.Offset(offset + 1)
			if cur.GreaterThan(pre) {
				return cur.Sub(pre)
			}
			return decimal.Zero
		})
		return gain
	})
}
