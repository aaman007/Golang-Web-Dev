package controllers

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

type CoreController struct {
	template *template.Template
}

func NewCoreController(template *template.Template) *CoreController {
	return &CoreController{template}
}

func (cc CoreController) Home(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	cc.template.ExecuteTemplate(w, "core_index.html", nil)
}