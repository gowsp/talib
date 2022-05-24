package talib

import "github.com/shopspring/decimal"

type Indicator interface {
	// Offset value forward based on the latest value,
	// Offset(0) The latest value added, Offset(1) The value of the previous period
	Offset(offset uint64) decimal.Decimal
	// OutOfBounds returns true if the offset is out of bounds
	OutOfBounds(offset uint64) bool
	// BarSeries get time series data
	BarSeries() *TimeSeries
	// Load the associated indicator in the current indicator cache,
	// use the supply function to generate and save when it does not exist
	LoadOrStore(id string, supplier func() Indicator) Indicator
	// Delete remove expired data
	Delete(expired []uint64)
}
