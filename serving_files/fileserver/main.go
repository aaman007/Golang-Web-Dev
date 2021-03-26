package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./media"))))
	http.Handle("/media/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/assets/image.jpg" />`)
}

func about(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/media/image.jpg" />`)
}
