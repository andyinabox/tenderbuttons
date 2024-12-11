package approuter

import (
	"html/template"
)

type Content struct {
	Corpus []byte
	Readme []byte
}

type Router struct {
	tpl     *template.Template
	content *Content
}

func New(templates *template.Template, content *Content) *Router {
	return &Router{templates, content}
}
