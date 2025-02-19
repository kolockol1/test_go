package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
	"syscall"
	"time"

	"app/pkg/models/poloniex"
	"app/pkg/services/kline"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		return
	}

	// Create context that will be canceled on interrupt signal
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle interrupt signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		fmt.Printf("Received signal: %v\n", sig)
		cancel()
	}()

	poloniexClient := kline.NewPoloniexClient(
		os.Getenv("POLONIEX_KEY"),
		os.Getenv("POLONIEX_SIGNATURE"),
		os.Getenv("POLONIEX_SIGNATURE_METHOD"),
		os.Getenv("POLONIEX_SIGNATURE_VERSION"),
		os.Getenv("POLONIEX_SIGN_TIMESTAMP"),
	)

	handleBtcUsdt(poloniexClient, "BTC_USDT_PERP", "MINUTE_1")

	// TODO: Add initialization and logic

	<-ctx.Done()
	fmt.Println("Shutting down gracefully...")
}

type httpPoloniexClient interface {
	FetchKlineData(request poloniex.KlineRequest) ([]poloniex.KlineData, error)
}

func handleBtcUsdt(client httpPoloniexClient, symbol, interval string) {
	//initialTs := 1733058000
	initialTs := 1733058000000
	finalTs := int(time.Now().UnixNano() / 1000000)
	result := make([]poloniex.KlineData, 0, 1_000_000)
	for {
		klineData, err := client.FetchKlineData(poloniex.KlineRequest{
			Symbol:    symbol,   //"BTC_USDT_PERP",
			Interval:  interval, //"MINUTE_1",
			Limit:     500,      // max
			StartTime: initialTs,
			EndTime:   finalTs,
		})
		if err != nil {
			fmt.Printf("Error fetching kline data: %v\n", err)
			break
		}
		if len(klineData) == 0 {
			fmt.Printf("Done: %s for interval %s", symbol, interval)
			break
		}
		// 1733058000000
		// 1733403600000
		// 1739999040000
		// 1739999096
		// 1739999096
		// 1739999004538823000

		result = append(result, klineData...)
		initialTs = klineData[len(klineData)-1].EndTime + 1
	}

	fmt.Printf("Kline data: %v\n", result)
}
