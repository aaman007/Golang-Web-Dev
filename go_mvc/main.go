package main

import (
	"github.com/aaman007/Golang-Web-Dev/go_mvc/controllers"
	"github.com/aaman007/Golang-Web-Dev/go_mvc/session"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/core/*.html"))
}

func main() {
	r := httprouter.New()
	mp := session.GetSession()

	// User Routes
	uc := controllers.NewUserController(mp)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	// Core Routes
	cc := controllers.NewCoreController(tpl)
	r.GET("/", cc.Home)

	http.ListenAndServe("localhost:8000", r)
}

/**
curl -X POST -H "Content-Type: application/json" -d '{"name":"Aaman","gender":"Male","age":24}' http://localhost:8000/user

curl http://localhost:8000/user/<enter-user-id-here>

curl -X DELETE http://localhost:8000/user/<enter-user-id-here>
 */