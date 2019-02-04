package main

import (
	"net"
	"net/http"
	"time"
)

// Transport creates a default transport defining dial and tslhandshake timeout
var Transport = &http.Transport{
	Dial: (&net.Dialer{
		Timeout: 10 * time.Second,
	}).Dial,
	TLSHandshakeTimeout: 10 * time.Second,
}

// Client is a HTTP client timeout, with the above Transport
var Client = &http.Client{
	Timeout:   time.Second * 10,
	Transport: Transport,
}
