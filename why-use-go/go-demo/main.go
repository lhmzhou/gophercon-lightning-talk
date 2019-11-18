package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["input"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url 'input' is missing")
		return
	}

	key := keys[0]
	fmt.Fprintf(w, "Go received: "+key)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
