package talib

import (
	"github.com/gowsp/talib/internal"
	"github.com/shopspring/decimal"
)

// Loss
func Loss(i Indicator) Indicator {
	return i.LoadOrStore(string(internal.LOSS), func() Indicator {
		loss := NewCachedIndicator(i, func(i Indicator, offset uint64) decimal.Decimal {
			if i.OutOfBounds(offset) {
				return decimal.Zero
			}
			cur := i.Offset(offset)
			pre := i.Offset(offset + 1)
			if cur.LessThan(pre) {
				return pre.Sub(cur)
			}
			return decimal.Zero
		})
		return loss
	})
}
