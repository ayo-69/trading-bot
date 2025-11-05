package data

import (
	"encoding/csv"
	"os"
	"strconv"
)

type Candle struct {
	Timestamp                      int64
	Open, High, Low, Close, Volume float64
}

func LoadCSV(path string) ([]Candle, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var candles []Candle
	for _, row := range rows[1:] { // skip header
		ts, _ := strconv.ParseInt(row[0], 10, 64)
		o, _ := strconv.ParseFloat(row[1], 64)
		h, _ := strconv.ParseFloat(row[2], 64)
		l, _ := strconv.ParseFloat(row[3], 64)
		c, _ := strconv.ParseFloat(row[4], 64)
		v, _ := strconv.ParseFloat(row[5], 64)
		candles = append(candles, Candle{ts, o, h, l, c, v})
	}

	return candles, nil
}
