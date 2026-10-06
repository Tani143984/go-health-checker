package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

type Result struct {
	URL          string        `json:"url"`
	Status       string        `json:"status"`
	StatusCode   int           `json:"status_code"`
	ResponseTime time.Duration `json:"response_time_ms"`
	Error        string        `json:"error,omitempty"`
}

func checkURL(url string, wg *sync.WaitGroup, results chan<- Result) {
	defer wg.Done()

	start := time.Now()
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	duration := time.Since(start)

	res := Result{
		URL:          url,
		ResponseTime: duration / time.Millisecond,
	}

	if err != nil {
		res.Status = "DOWN"
		res.StatusCode = 0
		res.Error = err.Error()
		results <- res
		return
	}
	defer resp.Body.Close()

	res.Status = "UP"
	res.StatusCode = resp.StatusCode
	results <- res
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://google.com",
		"https://github.com",
		"https://httpbin.org/delay/2",
		"https://invalid-domain-test-xyz.com",
	}

	fmt.Printf("Starting concurrent health check for %d endpoints...\n", len(urls))

	var wg sync.WaitGroup
	results := make(chan Result, len(urls))

	for _, u := range urls {
		wg.Add(1)
		go checkURL(u, &wg, results)
	}

	wg.Wait()
	close(results)

	var finalReport []Result
	for r := range results {
		finalReport = append(finalReport, r)
	}

	fileData, err := json.MarshalIndent(finalReport, "", "  ")
	if err != nil {
		fmt.Printf("Failed to generate JSON report: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n--- Execution Report ---")
	fmt.Println(string(fileData))
}
