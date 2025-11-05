package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ayo-69/trading-bot/internal/backtest"
	"github.com/ayo-69/trading-bot/internal/data"
	"github.com/ayo-69/trading-bot/internal/exchange"
	"github.com/ayo-69/trading-bot/internal/risk"
	"github.com/ayo-69/trading-bot/internal/strategy"
	"github.com/ayo-69/trading-bot/internal/types"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Could not load environment variables")
	}
	dataSrcFile := os.Getenv("DATA_SRC_FILE")
	if dataSrcFile == "" {
		log.Fatalf("DATA_SRC_FILE environment variable not set")
	}
	candles, err := data.LoadCSV(dataSrcFile)
	if err != nil {
		log.Fatalf("failed to load data: %v", err)
	}

	exch := exchange.NewSimulatedExchange(10000, "BTCUSDT")
	strat := strategy.NewSMACrossover(10, 30)
	rm := risk.NewManger(0.01)

	engine := backtest.NewEngine(exch, strat, rm)
	equity := engine.Run(candles)

	tradesFile := os.Getenv("TRADES_FILE")
	if tradesFile == "" {
		log.Fatalf("TRADES_FILE environment variable not set")
	}
	if err := writeTrades(tradesFile, exch.Trades); err != nil {
		log.Fatalf("failed to write trades: %v", err)
	}

	fmt.Printf("Final Equity: $%.5f\n", equity)
}

func writeTrades(filename string, trades []types.Trade) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Timestamp", "Symbol", "Side", "Price", "Quantity"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, trade := range trades {
		record := []string{
			trade.Timestamp.String(),
			trade.Symbol,
			trade.Side,
			strconv.FormatFloat(trade.Price, 'f', -1, 64),
			strconv.FormatFloat(trade.Quantity, 'f', -1, 64),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}
