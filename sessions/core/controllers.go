package core

import (
	"../session"
	"net/http"
)

func HomeController(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "home.html", session.GetUser(w, req))
}
