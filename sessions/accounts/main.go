package accounts

import (
	"../database"
	"html/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("accounts/templates/*.html"))
}

func CreateUser(username, firstname, lastname, role string, password []byte) database.User {
	user := database.User{
		Username: username,
		Firstname: firstname,
		Lastname: lastname,
		Password: password,
		Role: role,
	}
	database.DBUsers[username] = user
	return user
}