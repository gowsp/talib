package talib

import "github.com/shopspring/decimal"

// Rsi Relative strength index. It is calculated based on rma's of upward and downward change of x.
func Rsi(i Indicator, length uint64) Indicator {
	id := RSI.Id(length)
	return i.LoadOrStore(id, func() Indicator {
		gain := Rma(Gain(i), length)
		loss := Rma(Loss(i), length)
		rsi := NewCachedIndicator(i)
		rsi.calculate = func(offset uint64) decimal.Decimal {
			averageGain := gain.Offset(offset)
			averageLoss := loss.Offset(offset)
			if averageLoss.IsZero() {
				if averageGain.IsZero() {
					return decimal.Zero
				} else {
					return HUNDRED
				}
			}
			rs := averageGain.Div(averageLoss)
			return HUNDRED.Sub(HUNDRED.Div(ONE.Add(rs)))
		}
		return rsi
	})
}
