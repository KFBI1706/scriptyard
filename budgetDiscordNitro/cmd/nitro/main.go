package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var (
	emojis []Emoji
)

// Emoji is a small struct with the basic information about an emoji
type Emoji struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	Animated bool   `json:"animated"`
}

func main() {
	res, err := http.Get("http://kfbi.xyz:11337/api/v1/emojis")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&emojis)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(emojis)

}
