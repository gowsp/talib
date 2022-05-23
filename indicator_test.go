package talib

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func load() *TimeSeries {
	period := time.Minute * 15
	series := NewLimitSeries(period, 1000, 1250)
	file, err := os.OpenFile("testdata/kline-2022-04.csv", os.O_RDONLY, 0644)
	if err != nil {
		log.Panicln(err)
	}
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		line := scanner.Text()
		if line == "" {
			continue
		}
		// Open time, Open, High, Low, Close, Volume
		fields := strings.Split(line, ",")
		begin, _ := strconv.ParseInt(fields[0], 10, 64)
		open := decimal.RequireFromString(fields[1])
		high := decimal.RequireFromString(fields[2])
		low := decimal.RequireFromString(fields[3])
		close := decimal.RequireFromString(fields[4])
		t := time.UnixMilli(begin)
		bar := NewBar(period, t, open, high, low, close)
		series.Add(bar)
	}
	fmt.Println(series.Offset(0).BeginTime().Format("2006-01-02 15:04:05"))
	return series
}

func TestHighest(t *testing.T) {
	// plot(ta.highest(close, 100))
	series := load()
	close := series.ClosePriceIndicator()
	highest := Highest(close, 100)
	fmt.Println(close.Offset(0).StringFixed(10), highest.Offset(0).StringFixed(10))
}
func TestLowest(t *testing.T) {
	// plot(ta.lowest(close, 50))
	series := load()
	close := series.ClosePriceIndicator()
	lowest := Lowest(close, 50)
	fmt.Println(close.Offset(0).StringFixed(10), lowest.Offset(0).StringFixed(10))
}

func TestEma(t *testing.T) {
	// plot(ta.ema(close, 50))
	// plot(ta.ema(close, 100))
	series := load()
	close := series.ClosePriceIndicator()
	ema_50 := Ema(close, 50)
	ema_100 := Ema(close, 100)
	fmt.Println(ema_50.Offset(0).StringFixed(10), ema_100.Offset(0).StringFixed(10))
}

func TestAtr(t *testing.T) {
	// plot(ta.atr(100))
	series := load()
	fmt.Println(series.Atr(100).Offset(0).StringFixed(10))
}