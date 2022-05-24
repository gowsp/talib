package talib

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/shopspring/decimal"
)

type TimeSeries struct {
	data       map[uint64]Bar
	period     uint64
	latest     uint64
	capacity   uint64
	threshold  uint64
	indicators sync.Map
	mu         sync.RWMutex
	isCleaning uint32
}

func NewSeries(period time.Duration) *TimeSeries {
	return NewLimitSeries(period, 500, 625)
}
func NewLimitSeries(period time.Duration, capacity, threshold uint64) *TimeSeries {
	if threshold < capacity {
		panic("threshold must be greater than mcapacity")
	}
	return &TimeSeries{
		capacity:  capacity,
		threshold: threshold,
		data:      map[uint64]Bar{},
		period:    uint64(period.Seconds()),
	}
}
func (series *TimeSeries) LoadOrStore(id string, supplier func() Indicator) Indicator {
	if value, ok := series.indicators.Load(id); ok {
		return value.(Indicator)
	}
	indicator := supplier()
	series.indicators.Store(id, indicator)
	return indicator
}
func (series *TimeSeries) Add(bar Bar) {
	series.add(bar)
	series.reduce()
}
func (series *TimeSeries) add(bar Bar) {
	if uint64(bar.Period().Seconds()) != series.period {
		return
	}
	valueTime := uint64(bar.BeginTime().Unix())
	if valueTime%series.period != 0 {
		return
	}

	series.mu.Lock()
	defer series.mu.Unlock()
	series.data[valueTime] = bar
	if series.latest >= valueTime {
		return
	}
	series.latest = valueTime
}
func (series *TimeSeries) reduce() {
	series.mu.RLock()
	size := uint64(len(series.data))
	latest := series.latest
	series.mu.RUnlock()
	if size <= series.threshold || atomic.LoadUint32(&series.isCleaning) > 0 {
		return
	}
	go series.clean(size, latest)
}
func (series *TimeSeries) clean(size, latest uint64) {
	atomic.StoreUint32(&series.isCleaning, 1)
	expired := make([]uint64, 0, size-series.capacity)
	for i := size - 1; i >= series.capacity; i-- {
		cursor := latest - series.period*i
		series.mu.Lock()
		delete(series.data, cursor)
		series.mu.Unlock()
		expired = append(expired, cursor)
	}
	series.indicators.Range(func(key, value interface{}) bool {
		value.(Indicator).Delete(expired)
		return true
	})
	atomic.StoreUint32(&series.isCleaning, 0)
}
func (series *TimeSeries) Size() uint64 {
	series.mu.RLock()
	defer series.mu.RUnlock()
	return uint64(len(series.data))
}
func (series *TimeSeries) Offset(i uint64) Bar {
	series.mu.RLock()
	defer series.mu.RUnlock()
	cursor := series.latest - series.period*i
	return series.data[cursor]
}
func (series *TimeSeries) Cursor(i uint64) uint64 {
	series.mu.RLock()
	defer series.mu.RUnlock()
	return series.latest - series.period*i
}

func (series *TimeSeries) NewIndicator(id string, method func(Bar) decimal.Decimal) Indicator {
	return series.LoadOrStore(id, func() Indicator {
		indicator := NewCacheFrom(series, func(i Indicator, offset uint64) decimal.Decimal {
			if i.OutOfBounds(offset) {
				return decimal.Zero
			}
			bar := i.BarSeries().Offset(offset)
			if bar == nil {
				return decimal.Zero
			}
			return method(bar)
		})
		return indicator
	})
}
