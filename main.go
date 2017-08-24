package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT string = ":1337"

func handlePing(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "Pong")
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/ping", handlePing)

	log.Println("Listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}
