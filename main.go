package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
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
	client := http.Client{
		Timeout: time.Duration(timeoutFlag) * time.Second,
	}

	start := time.Now()
	resp, err = client.Get(url)

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

func printInformation(urls []string, result chan Result) {
	failed := false
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	_, _ = fmt.Fprintf(w, "URL\tSTATUS\tLATENCY\tERROR\n")

	for range len(urls) {
		res := <-result

		if res.StatusCode != 200 {
			failed = true
		}

		printErr := ""

		if res.Err == nil {
			printErr = "N/A"
			_, _ = fmt.Fprintf(w, "%s\t%d\t%v\t%v\n", res.URL, res.StatusCode, res.Latency, printErr)
		} else {
			_, _ = fmt.Fprintf(w, "%s\t%d\t%v\t%v\n", res.URL, res.StatusCode, res.Latency, res.Err)
		}

	}

	err := w.Flush()
	if err != nil {
		log.Fatalf("Error flushing the Tabwriter, %e", err)
	}

	if failed {
		os.Exit(1)
	}
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

	printInformation(urls, result)
}
