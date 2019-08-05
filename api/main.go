package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
