package main

import (
	"github.com/aaman007/Golang-Web-Dev/postgres/blog"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()

	router.GET("/blogs/", blog.ListHandler)
	router.GET("/", blog.ListHandler)
	router.GET("/blogs/new/", blog.CreateFormHandler)
	router.POST("/blogs/new/", blog.CreateHandler)
	router.GET("/blogs/details/:id/", blog.DetailsHandler)
	router.GET("/blogs/update/:id/", blog.UpdateFormHandler)
	router.POST("/blogs/update/:id/", blog.UpdateHandler)
	router.GET("/blogs/delete-confirm/:id/", blog.DeleteConfirmHandler)
	router.GET("/blogs/delete/:id/", blog.DeleteHandler)

	http.ListenAndServe(":8000", router)
}