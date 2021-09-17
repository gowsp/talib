package talib

import (
	"sync"

	"github.com/shopspring/decimal"
)

var ONE = decimal.NewFromInt(1)
var TWO = decimal.NewFromInt(2)
var HUNDRED = decimal.NewFromInt(100)

type CachedIndicator struct {
	series     *TimeSeries
	indicators sync.Map
	results    sync.Map
	calculate  func(offset uint64) decimal.Decimal
}

func NewCacheFrom(series *TimeSeries) *CachedIndicator {
	return &CachedIndicator{
		series: series,
	}
}
func NewCachedIndicator(indicator Indicator) *CachedIndicator {
	return &CachedIndicator{
		series: indicator.BarSeries(),
	}
}

func (series *CachedIndicator) LoadOrStore(key string, supplier func() Indicator) Indicator {
	if value, ok := series.indicators.Load(key); ok {
		return value.(Indicator)
	}
	indicator := supplier()
	series.indicators.Store(key, indicator)
	return indicator
}
func (s *CachedIndicator) OutOfBounds(offset uint64) bool {
	return offset >= s.series.Size()
}
func (s *CachedIndicator) BarSeries() *TimeSeries {
	return s.series
}
func (i *CachedIndicator) Offset(offset uint64) decimal.Decimal {
	if offset == 0 || i.OutOfBounds(offset) {
		return i.calculate(offset)
	}
	cursor := i.series.Cursor(offset)
	if res, ok := i.results.Load(cursor); ok {
		return res.(decimal.Decimal)
	}
	res := i.calculate(offset)
	i.results.Store(cursor, res)
	return res
}
func (i *CachedIndicator) Delete(times []uint64) {
	for key := range times {
		i.results.Delete(key)
	}
	i.indicators.Range(func(key, value interface{}) bool {
		value.(Indicator).Delete(times)
		return true
	})
}
