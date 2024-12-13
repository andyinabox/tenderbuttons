package app

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

func (a *App) AboutHandler() http.HandlerFunc {

	log.Info("parsing about content")
	about := blackfriday.Run(a.cfg.AboutContent, blackfriday.WithExtensions(blackfriday.CommonExtensions|blackfriday.Footnotes))

	return func(w http.ResponseWriter, req *http.Request) {

		log.Debug("about handler")
		err := a.cfg.Templates.ExecuteTemplate(w, "about.html.tmpl", AboutContext{
			CssProperties: a.getCSSCustomProperties("tender buttons"),
			Body:          template.HTML(about),
		})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
