package templates

import (
	"text/template"
)

var TMPL *template.Template

func init() {
	TMPL, _ = template.ParseGlob("templates/*.html")
}
