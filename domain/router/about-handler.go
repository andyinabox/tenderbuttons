package router

import (
	"html/template"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/russross/blackfriday/v2"
)

type AboutContext struct {
	Body template.HTML
}

func (r *Router) AboutHandler(readme []byte) http.HandlerFunc {
	log.Info("parsing readme")
	about := blackfriday.Run(readme)

	return func(w http.ResponseWriter, req *http.Request) {
		log.Debug("about handler")
		err := r.tpl.ExecuteTemplate(w, "about.html.tmpl", AboutContext{
			Body: template.HTML(about),
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
