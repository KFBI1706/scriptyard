package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pkg/errors"

	"gopkg.in/ns3777k/go-shodan.v3/shodan"
)

var client *shodan.Client
var netTransport = &http.Transport{
	Dial: (&net.Dialer{
		Timeout: 10 * time.Second,
	}).Dial,
	TLSHandshakeTimeout: 10 * time.Second,
}

var netClient = &http.Client{
	Timeout:   time.Second * 10,
	Transport: netTransport,
}

func init() {
	// NewEnvClient creates new Shodan client using environment variable SHODAN_KEY as the token.
	client = shodan.NewEnvClient(netClient)
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

}

// SearchHosts wraps the shodan search api, seeing if the search will return any results, and if so returns the results for the specified page
func SearchHosts(ctx context.Context, search string, page int) (hosts *shodan.HostMatch, err error) {
	query := &shodan.HostQueryOptions{Query: search, Page: page}

	hostCount, err := client.GetHostsCountForQuery(ctx, query)
	if err != nil {
		return nil, err
	}

	if hostCount.Total == 0 {
		return nil, errors.New("Zero matching hosts for that query")
	}

	hosts, err = client.GetHostsForQuery(ctx, query)
	if err != nil {
		return nil, err
	}
	return

}
