package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello dexter morgan")
	})

	err := http.ListenAndServe("0.0.0.0:1072", nil)

	if err != nil {
		log.Fatalf("oh no! dexter morgan is dead: %v", err)
	}
}
