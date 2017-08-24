package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const PORT string = ":1337"

func handlePing(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "Pong")
}

func handlePictureUpload(responseWriter http.ResponseWriter, request *http.Request) {

	fmt.Println("I should handle the picture upload now")

	request.ParseMultipartForm(32 << 20)
	file, handler, err := request.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Fprintf(responseWriter, "%v", handler.Header)
	f, err := os.OpenFile("./uploads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("save error", err)
		return
	}
	defer f.Close()

	io.Copy(f, file)

}

func main() {
	fs := http.FileServer(http.Dir("static/build/default"))
	http.Handle("/", fs)
	http.HandleFunc("/ping", handlePing)
	http.HandleFunc("/pictures", handlePictureUpload)

	log.Println("Listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}
