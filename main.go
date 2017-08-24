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

func handlePictureUpload(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("I should handle the picture upload now")
	fmt.Fprintf(responseWriter, "Got it")
}

func main() {
	fs := http.FileServer(http.Dir("static/build/default"))
	http.Handle("/", fs)
	http.HandleFunc("/ping", handlePing)
	http.HandleFunc("/pictures", handlePictureUpload)

	log.Println("Listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}
