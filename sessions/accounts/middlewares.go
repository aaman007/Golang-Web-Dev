package accounts

import (
	"../session"
	"net/http"
)

func IsAdminUser(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		user := session.GetUser(w, req)
		if user.Role != "admin" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		h.ServeHTTP(w, req)
	})
}

func IsAuthorized(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if !session.AlreadyLoggedIn(w, req) {
			http.Redirect(w, req, "/login/", http.StatusSeeOther)
			return
		}
		h.ServeHTTP(w, req)
	})
}

func RedirectSignedInUser(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if session.AlreadyLoggedIn(w, req) {
			http.Redirect(w, req, "/", http.StatusSeeOther)
			return
		}
		h.ServeHTTP(w, req)
	})
}
