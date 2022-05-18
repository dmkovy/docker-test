package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/docker", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "<h1>Hello from Docker container!</h1>")
	})

	http.ListenAndServe(":8080", nil)
}
