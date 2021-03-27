package store

import "net/http"

func StoreController(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "store.html", nil)
}
