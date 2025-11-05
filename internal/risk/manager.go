package risk

import "github.com/ayo-69/trading-bot/internal/exchange"

type Manager struct {
	RiskFration float64
}

func NewManger(fraction float64) *Manager {
	return &Manager{RiskFration: fraction}
}

func (r *Manager) PostionSize(ex *exchange.SimulatedExchange, price float64) float64 {
	riskCapital := ex.Balance * r.RiskFration
	return riskCapital / price
}
