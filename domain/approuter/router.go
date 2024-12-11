package approuter

import (
	"html/template"

	"github.com/andyinaobox/tenderbuttons/pkg/handler"
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

func (r *Router) Routes() []handler.Route {
	return []handler.Route{
		{
			Path:        "/",
			HandlerFunc: r.IndexHandler(),
		},
		{
			Path:        "/readme",
			HandlerFunc: r.AboutHandler(),
		},
	}
}
