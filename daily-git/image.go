package main

import (
	"image"
	"image/jpeg"
	"log"
	"net/http"
)

func getImage(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()

	return jpeg.Decode(resp.Body)
}
