package exchange

import (
	"log"
	"time"

	"github.com/ayo-69/trading-bot/internal/data"
	"github.com/ayo-69/trading-bot/internal/types"
)

type SimulatedExchange struct {
	Balance  float64
	Position float64
	Trades   []types.Trade
	Symbol   string
}

func NewSimulatedExchange(balance float64, symbol string) *SimulatedExchange {
	return &SimulatedExchange{Balance: balance, Symbol: symbol, Trades: []types.Trade{}}
}

func (ex *SimulatedExchange) Buy(c data.Candle, qty float64) {
	cost := qty * c.Close
	if cost <= ex.Balance {
		ex.Balance -= cost
		ex.Position += qty
		log.Printf("BUY: balance=%.2f, position=%.8f", ex.Balance, ex.Position)
		ex.Trades = append(ex.Trades, types.Trade{
			Timestamp: time.Unix(c.Timestamp, 0),
			Symbol:    ex.Symbol,
			Side:      "BUY",
			Price:     c.Close,
			Quantity:  qty,
		})
	}
}

func (ex *SimulatedExchange) Sell(c data.Candle, qty float64) {
	if ex.Position >= qty {
		ex.Balance += qty * c.Close
		ex.Position -= qty
		log.Printf("SELL: balance=%.2f, position=%.8f", ex.Balance, ex.Position)
		ex.Trades = append(ex.Trades, types.Trade{
			Timestamp: time.Unix(c.Timestamp, 0),
			Symbol:    ex.Symbol,
			Side:      "SELL",
			Price:     c.Close,
			Quantity:  qty,
		})
	}
}

func (ex *SimulatedExchange) Equity(c data.Candle) float64 {
	return ex.Balance + ex.Position*c.Close
}
