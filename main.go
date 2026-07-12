package main

import (
	"fmt"
	"log"
	"net/http"
)

func Get(url string) (resp *http.Response, err error) {
	resp, err = http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	return resp, err
}

func main() {
	resp, err := Get("https://www.google.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode)
}
