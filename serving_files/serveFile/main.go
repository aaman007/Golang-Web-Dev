package main

import (
	"io"
	"net/http"
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
	http.ServeFile(w, req, "image.jpg")
}
