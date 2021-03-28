package main

import (
	"io"
	"net/http"
)

func homeDefault(res http.ResponseWriter, req *http.Request)  {
	io.WriteString(res, "THIS IS HOME PAGE")
}

func aboutDefault(res http.ResponseWriter, req *http.Request)  {
	io.WriteString(res, "THIS IS ABOUT PAGE")
}

func main() {
	http.HandleFunc("/", homeDefault)
	http.HandleFunc("/about", aboutDefault)

	http.ListenAndServe(":8000", nil)
}
