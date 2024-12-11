package router

import (
	"html/template"

	"github.com/andyinaobox/tenderbuttons/pkg/handler"
)

type Data struct {
	Corpus []byte
	Readme []byte
}

type Router struct {
	tpl  *template.Template
	data *Data
}

func New(tpl *template.Template, data *Data) *Router {
	return &Router{tpl, data}
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
