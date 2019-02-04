package main

import (
	"context"
	"os"
	"os/signal"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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

	//dat, err := ioutil.ReadFile("FILE")
	//check(err)
}
