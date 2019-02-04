package main

import (
	"context"
	"os"
	"os/signal"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	// Cancel upon Ctrl+C
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		<-signalChan
		cancel()
	}()

}
