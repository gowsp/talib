package talib

import "fmt"

type IndicatorEnum string

const (
	OPEN   IndicatorEnum = "open"
	CLOSE  IndicatorEnum = "close"
	HIGH   IndicatorEnum = "high"
	LOW    IndicatorEnum = "low"
	VOLUME IndicatorEnum = "volume"

	MEDIAN  IndicatorEnum = "median"
	TYPICAL IndicatorEnum = "typical"
	ATR     IndicatorEnum = "atr"
)
const (
	SMA IndicatorEnum = "sma"
	RMA IndicatorEnum = "rma"
	EMA IndicatorEnum = "ema"
	RSI IndicatorEnum = "rsi"
	ROC IndicatorEnum = "roc"
)
const (
	TR   IndicatorEnum = "tr"
	GAIN IndicatorEnum = "gain"
	LOSS IndicatorEnum = "loss"
)

func (e IndicatorEnum) Id(params ...interface{}) string {
	if len(params) == 0 {
		return string(e)
	}
	return fmt.Sprintf("%s%v", e, params)
}
