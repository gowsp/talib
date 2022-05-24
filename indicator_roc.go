package talib

import (
	"github.com/gowsp/talib/internal"
	"github.com/shopspring/decimal"
)

// Roc roc (rate of change) showing the difference between current value of x and the value of x that was y days ago.
func Roc(i Indicator, length uint64) Indicator {
	id := internal.ROC.Id(length)
	return i.LoadOrStore(id, func() Indicator {
		roc := NewCachedIndicator(i, func(i Indicator, offset uint64) decimal.Decimal {
			before := i.Offset(offset + length)
			if before.Equal(decimal.Zero) {
				return decimal.Zero
			}
			return HUNDRED.Mul(i.Offset(offset).Div(before).Sub(ONE))
		})
		return roc
	})
}
