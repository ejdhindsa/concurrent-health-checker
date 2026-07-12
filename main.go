package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Get(url string) (resp *http.Response, elapsed time.Duration, err error) {
	start := time.Now()
	resp, err = http.Get(url)
	if err != nil {
		log.Printf("Execution failed, %v", err)
	}
	elapsed = time.Since(start)
	return resp, elapsed, err
}

func main() {
	resp, latency, err := Get("https://www.ekamjot.me")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode)
	fmt.Printf("Latency time: %v", latency)
}
