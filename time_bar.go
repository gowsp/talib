package talib

import (
	"time"

	"github.com/shopspring/decimal"
)

type Bar interface {
	// the time period of the bar
	Period() time.Duration
	// the begin time of the bar period
	BeginTime() time.Time
	// the end time of the bar period
	EndTime() time.Time
	// the open price of the period
	OpenPrice() decimal.Decimal
	// the close price of the period
	ClosePrice() decimal.Decimal
	// the high price of the period
	HighPrice() decimal.Decimal
	// the low price of the period
	LowPrice() decimal.Decimal
	// the whole tradeNum volume in the period
	Volume() decimal.Decimal
	// the whole traded amount of the period
	Amount() decimal.Decimal
	// the number of trades in the period
	Trades() decimal.Decimal
}

type BaseBar struct {
	PeriodVal     time.Duration
	BeginTimeVal  time.Time
	EndTimeVal    time.Time
	OpenPriceVal  decimal.Decimal
	ClosePriceVal decimal.Decimal
	HighPriceVal  decimal.Decimal
	LowPriceVal   decimal.Decimal
	VolumeVal     decimal.Decimal
	AmountVal     decimal.Decimal
	TradesVal     decimal.Decimal
}

func (candle *BaseBar) Period() time.Duration {
	return candle.PeriodVal
}
func (candle *BaseBar) BeginTime() time.Time {
	return candle.BeginTimeVal
}
func (candle *BaseBar) EndTime() time.Time {
	return candle.EndTimeVal
}
func (candle *BaseBar) OpenPrice() decimal.Decimal {
	return candle.OpenPriceVal
}
func (candle *BaseBar) ClosePrice() decimal.Decimal {
	return candle.ClosePriceVal
}
func (candle *BaseBar) HighPrice() decimal.Decimal {
	return candle.HighPriceVal
}
func (candle *BaseBar) LowPrice() decimal.Decimal {
	return candle.LowPriceVal
}
func (candle *BaseBar) Volume() decimal.Decimal {
	return candle.VolumeVal
}
func (candle *BaseBar) Amount() decimal.Decimal {
	return candle.AmountVal
}
func (candle *BaseBar) Trades() decimal.Decimal {
	return candle.TradesVal
}

func NewBar(period time.Duration, begin time.Time, open, high, low, close decimal.Decimal) Bar {
	return &BaseBar{
		PeriodVal:     period,
		BeginTimeVal:  begin,
		OpenPriceVal:  open,
		HighPriceVal:  high,
		LowPriceVal:   low,
		ClosePriceVal: close,
	}
}
