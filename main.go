package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const PORT string = ":1337"

func runCommand(command string) string {
	commandStructure := strings.Split(command, " ")

	args := commandStructure[1:]

	spacedArgs := strings.Join(args, " ")

	commandOutput, err := exec.Command(commandStructure[0], spacedArgs).Output()
	if err != nil {
		return err.Error()
	}
	return string(commandOutput)
}

func setWallpaper() {
	// gsettings get org.gnome.desktop.background picture-uri
	// 'file:///[path]/x.jpg'
	fp, _ := filepath.Abs("./")
	fp = fmt.Sprintf("'file:///%s/uploads/wallpaper.png'", fp)
	fmt.Println("setting wallpaper to ", fp)

	command := fmt.Sprintf("gsettings get org.gnome.desktop.background picture-uri %s", fp)

	fmt.Println("Executing command: ", command)
	response := runCommand(command)
	fmt.Println(response)
}

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
	f, err := os.OpenFile("./uploads/wallpaper.png", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("save error", err)
		return
	}
	defer f.Close()

	io.Copy(f, file)
	setWallpaper()
}

func main() {
	fs := http.FileServer(http.Dir("static/build/default"))
	http.Handle("/", fs)
	http.HandleFunc("/ping", handlePing)
	http.HandleFunc("/pictures", handlePictureUpload)

	log.Println("Listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}
