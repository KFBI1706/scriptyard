package main

// Project contains all the information about a specific shodan project, which is a collection of hosts with a descriptive name containing it all
type Project struct {
	Name         string `yaml:""`
	Hosts        int    `json:"hosts"`
	SearchString string `json:"search"`
}

type Config struct {
}
