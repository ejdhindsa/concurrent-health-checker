package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	concurrencyFlag int
	timeoutFlag     int
)

func init() {
	flag.IntVar(&concurrencyFlag, "c", 10, "Number of concurrent tasks.")
	flag.IntVar(&timeoutFlag, "timeout", 10, "Timeout before the URL is skipped.")
}

func Get(url string) (resp *http.Response, elapsed time.Duration, err error) {
	start := time.Now()
	resp, err = http.Get(url)
	if err != nil {
		log.Printf("Execution failed, %v", err)
	}
	elapsed = time.Since(start)
	return resp, elapsed, err
}

func ReadURLs() (urls []string, err error) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		// if empty line contine
		if line == "" {
			continue
		}

		urls = append(urls, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}

func main() {
	flag.Parse()

	resp, latency, err := Get("https://www.ekamjot.me")
	if err != nil {
		log.Fatal(err)
	}

	urls, err := ReadURLs()
	if err != nil {
		log.Fatal(err)
	}

	for _, value := range urls {
		fmt.Printf("URL: %v\n", value)
	}

	fmt.Println(resp.StatusCode)
	fmt.Printf("Latency time: %v\n", latency)

	fmt.Printf("Concurrency flag value, %v\n", concurrencyFlag)
}
