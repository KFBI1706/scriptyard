package main

import (
	"context"

	"http" // internal HTTP package

	"github.com/pkg/errors"
	"gopkg.in/ns3777k/go-shodan.v3/shodan"
)

var client *shodan.Client

func init() {
	// NewEnvClient creates new Shodan client using environment variable SHODAN_KEY as the token.
	client = shodan.NewEnvClient(http.Client)
}

// SearchHosts wraps the shodan search api, seeing if the search will return any results, and if so returns the results for the specified page
func SearchHosts(ctx context.Context, search string, page int) (hosts *shodan.HostMatch, err error) {
	query := &shodan.HostQueryOptions{Query: search, Page: page}

	hostCount, err := client.GetHostsCountForQuery(ctx, query)
	if err != nil {
		return nil, err
	}

	if hostCount.Total == 0 {
		return nil, errors.Errorf("Zero matching hosts for the query '%s' on page %d", search, page)
	}

	hosts, err = client.GetHostsForQuery(ctx, query)
	if err != nil {
		return nil, err
	}
	return

}
