package store

import (
	"html/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("store/templates/*.html"))
}