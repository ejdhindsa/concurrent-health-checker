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

type Result struct {
	URL        string
	StatusCode int
	Latency    time.Duration
	Err        error
}

func worker(jobs <-chan string, results chan<- Result) {
	for url := range jobs {
		resp, latency, err := Get(url)

		var status int

		if resp != nil {
			status = resp.StatusCode
		}

		result := Result{
			URL:        url,
			StatusCode: status,
			Latency:    latency,
			Err:        err,
		}

		results <- result
	}
}

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

	urls, err := ReadURLs()
	if err != nil {
		log.Fatal(err)
	}

	jobs := make(chan string, len(urls))
	result := make(chan Result, len(urls))

	for range concurrencyFlag {
		go worker(jobs, result)
	}

	for _, url := range urls {
		jobs <- url
	}

	close(jobs)

	for range len(urls) {
		res := <-result

		fmt.Printf("URL: %s | Status: %d | Latency: %v\n", res.URL, res.StatusCode, res.Latency)
	}
}
