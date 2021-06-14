package http

import "html/template"

func NewTpl() *template.Template {
	return template.New("").Delims("<obs>", "</obs>")
}
