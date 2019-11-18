package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	url := os.Args[1]
	hits, _ := strconv.Atoi(os.Args[2]) // Assume good input

	// Just in case there's no prefix
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	// Start the timer
	startTime := time.Now()

	// Synchronous loop
	for i := 1; i <= hits; i++ {
		fetch(i, url)
	}

	// End the timer and output the results
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	fmt.Printf("Completed in %v\n", elapsed)
}

// Function to retrieve HTTP response body and output the response and status code
func fetch(i int, url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http error code: %v\n", resp.StatusCode)
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	}

	b, err := ioutil.ReadAll(resp.Body)

	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}

	fmt.Printf("%d: %s, HTTP Status: %v\n", i, b, resp.StatusCode)
}
