package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/hello", hello())
	mux.Handle("/", http.FileServer(http.Dir("./static/")))

	log.Fatal(http.ListenAndServe(":3000", mux))
}
