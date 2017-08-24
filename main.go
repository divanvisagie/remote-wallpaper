package main

import (
	"log"
	"net/http"
)

const PORT string = ":1337"

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	log.Println("Listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}
