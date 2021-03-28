package blog

import (
	"database/sql"
	"fmt"
	"github.com/aaman007/Golang-Web-Dev/postgres/config"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func ListHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if req.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	blogs, err := All()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "blog_list.gohtml", blogs)
}

func DetailsHandler(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	if req.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	blogId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	blog, err := Get(blogId)
	if err == sql.ErrNoRows {
		http.NotFound(w, req)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "blog_details.gohtml", blog)
}

func CreateFormHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if req.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	config.TPL.ExecuteTemplate(w, "blog_create.gohtml", nil)
}

func CreateHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if req.Method != http.MethodPost {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	title := req.FormValue("title")
	body := req.FormValue("body")

	if title == "" || body == "" {
		http.Error(w, "All fields must be complete.", http.StatusBadRequest)
		return
	}

	_, err := Create(title, body)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func UpdateFormHandler(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	if req.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	blogId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	blog, err := Get(blogId)
	if err == sql.ErrNoRows {
		http.NotFound(w, req)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "blog_update.gohtml", blog)
}

func UpdateHandler(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	if req.Method != http.MethodPost {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	blogId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	title := req.FormValue("title")
	body := req.FormValue("body")

	if title == "" || body == "" {
		http.Error(w, "All fields must be complete.", http.StatusBadRequest)
		return
	}

	_, err = Update(title, body, blogId)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func DeleteConfirmHandler(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	if req.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	blogId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	blog, err := Get(blogId)
	if err == sql.ErrNoRows {
		http.NotFound(w, req)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "blog_delete_confirm.gohtml", blog)
}

func DeleteHandler(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	if req.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	blogId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	err = Delete(blogId)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	http.Redirect(w, req, "/", http.StatusSeeOther)
}
