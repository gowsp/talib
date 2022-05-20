package internal

import "fmt"

type IndicatorEnum string

const (
	OPEN   = "open"
	CLOSE  = "close"
	HIGH   = "high"
	LOW    = "low"
	VOLUME = "volume"

	MEDIAN  = "median"
	TYPICAL = "typical"
)
const (
	TR   = "tr"
	ATR  = "atr"
	GAIN = "gain"
	LOSS = "loss"
)
const (
	SMA IndicatorEnum = "sma"
	RMA IndicatorEnum = "rma"
	EMA IndicatorEnum = "ema"
	RSI IndicatorEnum = "rsi"
	ROC IndicatorEnum = "roc"
)

func (e IndicatorEnum) Id(params ...interface{}) string {
	if len(params) == 0 {
		return string(e)
	}
	return fmt.Sprintf("%s%v", e, params)
}
