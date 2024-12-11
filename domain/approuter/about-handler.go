package approuter

import (
	"html/template"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/russross/blackfriday/v2"
)

type AboutContext struct {
	CssProperties template.HTML
	Body          template.HTML
}

func (r *Router) AboutHandler() http.HandlerFunc {
	log.Info("parsing readme")
	about := blackfriday.Run(r.cfg.Readme)

	return func(w http.ResponseWriter, req *http.Request) {
		log.Debug("about handler")
		err := r.tpl.ExecuteTemplate(w, "about.html.tmpl", AboutContext{
			CssProperties: r.getCSSCustomProperties("tender buttons"),
			Body:          template.HTML(about),
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
