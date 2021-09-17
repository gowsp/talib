package talib

// Atr (average true range) returns the RMA of true range.
func (series *TimeSeries) Atr(length uint64) Indicator {
	return Rma(series.Tr(), length)
}
