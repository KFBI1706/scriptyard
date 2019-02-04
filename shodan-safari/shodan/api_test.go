package main

import (
	"context"
	"reflect"
	"testing"

	"gopkg.in/ns3777k/go-shodan.v3/shodan"
)

func TestSearchHosts(t *testing.T) {
	type args struct {
		ctx    context.Context
		search string
		page   int
	}
	tests := []struct {
		name      string
		args      args
		wantHosts *shodan.HostMatch
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHosts, err := SearchHosts(tt.args.ctx, tt.args.search, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchHosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotHosts, tt.wantHosts) {
				t.Errorf("SearchHosts() = %v, want %v", gotHosts, tt.wantHosts)
			}
		})
	}
}
