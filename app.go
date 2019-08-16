package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	color := getColor()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1 style=\"color:%s;\">Jello</h1>", color)
	})
	port := getPort()
	panic(http.ListenAndServe(":"+port, nil))

}

func getPort() string {
	port := os.Getenv("PORT")
	i, _ := strconv.Atoi(port)
	if i == 0 {
		port = "3000"
	}

	fmt.Println("Listening on port: ", port)
	return port
}

func getColor() string {
	color := os.Getenv("COLOR")
	if color == "" {
		return "red"
	}
	return color
}
