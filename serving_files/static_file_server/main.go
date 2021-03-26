package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8000", http.FileServer(http.Dir("."))))

	// http.Error(w, "", errorCode)
	// errorCode: http.StatusOK = 200
	// errorCode: http.StatusNotFound = 404
}