package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
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

	// TODO: Add initialization and logic

	<-ctx.Done()
	fmt.Println("Shutting down gracefully...")
}
