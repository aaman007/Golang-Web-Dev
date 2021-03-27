package core

import (
	"html/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("core/templates/*.html"))
}
