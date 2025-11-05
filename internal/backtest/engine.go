package backtest

import (
	"github.com/ayo-69/trading-bot/internal/data"
	"github.com/ayo-69/trading-bot/internal/exchange"
	"github.com/ayo-69/trading-bot/internal/risk"
	"github.com/ayo-69/trading-bot/internal/strategy"
)

type Engine struct {
	exch  *exchange.SimulatedExchange
	strat *strategy.SMACrossover
	rm    *risk.Manager
}

func NewEngine(e *exchange.SimulatedExchange, s *strategy.SMACrossover, r *risk.Manager) *Engine {
	return &Engine{e, s, r}
}

func (eng *Engine) Run(candles []data.Candle) float64 {
	for i := range candles {
		if i < eng.strat.SlowPeriod {
			continue
		}

		sig := eng.strat.GenerateSignal(candles[:i])
		price := candles[i].Close
		size := eng.rm.PostionSize(eng.exch, price)

		switch sig {
		case "BUY":
			eng.exch.Buy(candles[i], size)
		case "SELL":
			eng.exch.Sell(candles[i], size)
		}
	}

	// Close any open position at the end of the backtest
	if eng.exch.Position > 0 {
		eng.exch.Sell(candles[len(candles)-1], eng.exch.Position)
	} else if eng.exch.Position < 0 {
		eng.exch.Buy(candles[len(candles)-1], -eng.exch.Position)
	}

	return eng.exch.Equity(candles[len(candles)-1])
}
