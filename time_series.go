package talib

import (
	"sync"
	"time"

	"github.com/shopspring/decimal"
)

type TimeSeries struct {
	data       []Bar
	period     uint64
	capacity   uint64
	threshold  uint64
	indicators sync.Map
	mu         sync.RWMutex
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
		data:      make([]Bar, 0, capacity),
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
}
func (series *TimeSeries) add(bar Bar) {
	if uint64(bar.Period().Seconds()) != series.period {
		return
	}
	series.mu.Lock()
	defer series.mu.Unlock()
	series.data = append(series.data, bar)
}

func (series *TimeSeries) Size() uint64 {
	series.mu.RLock()
	defer series.mu.RUnlock()
	return uint64(len(series.data))
}
func (series *TimeSeries) Offset(i uint64) Bar {
	series.mu.RLock()
	defer series.mu.RUnlock()
	cursor := series.Cursor(i)
	return series.data[cursor]
}
func (series *TimeSeries) Cursor(i uint64) uint64 {
	series.mu.RLock()
	defer series.mu.RUnlock()
	return series.Size() - i - 1
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
