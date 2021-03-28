package main

import (
	"io"
	"net/http"
)

func home(res http.ResponseWriter, req *http.Request)  {
	io.WriteString(res, "THIS IS HOME PAGE")
}

func about(res http.ResponseWriter, req *http.Request)  {
	io.WriteString(res, "THIS IS ABOUT PAGE")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)

	http.ListenAndServe(":8000", mux)
}
