package strategy

import (
	"github.com/ayo-69/trading-bot/internal/data"
	"log"
)

type SMACrossover struct {
	FastPeriod, SlowPeriod int
	Fast, Slow             []float64
	LastSignal             string
}

func NewSMACrossover(fast, slow int) *SMACrossover {
	return &SMACrossover{FastPeriod: fast, SlowPeriod: slow}
}

func (s *SMACrossover) calcSMA(prices []float64, period int) float64 {
	if len(prices) < period {
		return 0
	}
	sum := 0.0
	for _, p := range prices[len(prices)-period:] {
		sum += p
	}
	return sum / float64(period)
}

func (s *SMACrossover) GenerateSignal(candles []data.Candle) string {
	closes := []float64{}
	for _, c := range candles {
		closes = append(closes, c.Close)
	}

	fast := s.calcSMA(closes, s.FastPeriod)
	slow := s.calcSMA(closes, s.SlowPeriod)

	log.Printf("fast: %f, slow: %f", fast, slow)

	if fast == 0 || slow == 0 {
		return ""
	}

	if fast > slow && s.LastSignal != "BUY" {
		s.LastSignal = "BUY"
		return "BUY"
	} else if fast < slow && s.LastSignal != "SELL" {
		s.LastSignal = "SELL"
		return "SELL"
	}

	return ""
}
