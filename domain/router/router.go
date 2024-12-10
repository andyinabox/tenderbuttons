package router

import (
	"html/template"
)

type Router struct {
	tpl *template.Template
}

func New(tpl *template.Template) *Router {
	return &Router{tpl}
}
