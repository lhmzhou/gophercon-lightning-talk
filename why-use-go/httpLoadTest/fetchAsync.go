package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// TODO: Improve the code to write to log instead of console
func processAsync(hits int, url string) {
	async := len(os.Args) > 3

	if async {
		ch := make(chan string)

		for i := 1; i <= hits; i++ {

			go fetchAsync(i, ch, url)
			//fmt.Println(<-ch)
		}
		for i := 1; i <= hits; i++ {
			fmt.Println(<-ch) // receive from channel ch
		}
	} else {
		for i := 1; i <= hits; i++ {
			fetch(i, url)
		}
	}
}

func fetchAsync(i int, ch chan<- string, url string) { //}, logger *log.Logger) {

	resp, err := http.Get(url)
	if err != nil {
		//fmt.Fprintf(os.Stderr, "http error code: %v\n", resp.StatusCode)
		//fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		ch <- fmt.Sprintf("http error code: %v\n", resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)

	resp.Body.Close()
	if err != nil {
		//fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		ch <- fmt.Sprintf("fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}

	// Print output
	//logger.Printf("%d: %s", i, b)
	//logger.Printf(", HTTP Status: %v\n", resp.StatusCode)
	ch <- fmt.Sprintf("%d: %s, HTTP Status: %v", i, b, resp.StatusCode)
}
