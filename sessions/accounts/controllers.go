package accounts

import (
	"../database"
	"../session"
	"net/http"
	"time"
)

func LoginController(w http.ResponseWriter, req *http.Request) {
	if session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req,"/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		password := req.FormValue("password")

		user, ok := database.DBUsers[username]

		if !ok {
			http.Error(w, "Username or password is incorrect", http.StatusBadRequest)
			return
		}

		// TODO: Match Hashed Passwords
		if string(user.Password) != password {
			http.Error(w, "Username or password is incorrect", http.StatusBadRequest)
			return
		}

		session.CreateSessionCookie(w, req, username)

		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "login.html", nil)
}

func SignUpController(w http.ResponseWriter, req *http.Request) {
	if session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req,"/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		firstname := req.FormValue("firstname")
		lastname := req.FormValue("lastname")
		password := req.FormValue("password")
		role := req.FormValue("role")

		if _, ok := database.DBUsers[username]; ok {
			http.Error(w, "Username is already taken", http.StatusBadRequest)
			return
		}

		// TODO: Hash Password
		hashedPassword := []byte(password)
		CreateUser(username, firstname, lastname, role, hashedPassword)

		session.CreateSessionCookie(w, req, username)

		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.html", nil)
}

func LogoutController(w http.ResponseWriter, req *http.Request) {
	if !session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req,"/login", http.StatusSeeOther)
		return
	}

	cookie, _ := req.Cookie("session")

	// Delete session from db
	delete(database.DBSessions, cookie.Value)

	// Remove Cookie
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)

	// cleanup dbSessions
	if time.Now().Sub(database.DBCleanedSession) > (time.Second * 30) {
		go session.CleanSessions()
	}
	http.Redirect(w, req,"/login", http.StatusSeeOther)
}

func DashboardController(w http.ResponseWriter, req *http.Request) {
	if !session.AlreadyLoggedIn(w, req) {
		http.Error(w, "/", http.StatusSeeOther)
		return
	}

	user := session.GetUser(w, req)
	if user.Role != "admin" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	tpl.ExecuteTemplate(w, "dashboard.html", user)
}