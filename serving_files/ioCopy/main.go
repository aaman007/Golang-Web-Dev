package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/image.jpg", image)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/image.jpg" />`)
}

func image(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("image.jpg")
	if err != nil {
		http.Error(w, "Image not found", 404)
	}
	defer f.Close()

	io.Copy(w, f)
}
