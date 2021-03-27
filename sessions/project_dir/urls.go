package project_dir

import (
	"../accounts"
	"../core"
	"../store"
	"../database"
	"net/http"
)

func getUrlPatterns() []database.Path {
	UrlPatterns := core.UrlPatterns
	UrlPatterns = append(UrlPatterns, accounts.UrlPatterns...)
	UrlPatterns = append(UrlPatterns, store.UrlPatterns...)
	return UrlPatterns
}

func SetRoutes() {
	for _, path := range getUrlPatterns() {
		http.HandleFunc(path.Route, path.HandlerFunction)
	}
}