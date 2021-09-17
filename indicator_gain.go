package talib

import "github.com/shopspring/decimal"

// Gain
func Gain(i Indicator) Indicator {
	return i.LoadOrStore(string(GAIN), func() Indicator {
		gain := NewCachedIndicator(i)
		gain.calculate = func(offset uint64) decimal.Decimal {
			if gain.OutOfBounds(offset) {
				return decimal.Zero
			}
			cur := i.Offset(offset)
			pre := i.Offset(offset + 1)
			if cur.GreaterThan(pre) {
				return cur.Sub(pre)
			}
			return decimal.Zero
		}
		return gain
	})
}
