package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type handler int

func (h handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	res.Header().Set("GoogleAPIKey", "HS&*S*GSHGSSJHGS&^%S&%SFVSSGFS")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	data := struct {
		Method string
		Submissions url.Values
		URL *url.URL
		Header http.Header
		ContentLength int64
	}{
		req.Method,
		req.Form,
		req.URL,
		req.Header,
		req.ContentLength,
	}

	tpl.ExecuteTemplate(res, "home.html", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("home.html"))
}

func main() {
	var h handler
	http.ListenAndServe(":8000", h)
}
