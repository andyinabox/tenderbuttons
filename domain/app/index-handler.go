package app

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/charmbracelet/log"

	"github.com/andyinaobox/tenderbuttons/pkg/chains"
)

type IndexContext struct {
	CssProperties template.HTML
	Sentence      string
	Tokens        []string
}

func (a *App) IndexHandler() http.HandlerFunc {
	// create markov chain
	log.Info("building chain", "prefixLen", a.cfg.PrefixLength)
	chain := chains.NewChain(&chains.Config{
		PrefixLength: a.cfg.PrefixLength,
		MinWords:     a.cfg.MinWords,
		MaxWords:     a.cfg.MaxWords,
	})
	chain.BuildFromString(string(a.cfg.Corpus))

	return func(w http.ResponseWriter, req *http.Request) {
		log.Debug("index handler")
		var err error
		var sentence string

		req.ParseForm()

		// create sentance from token
		if tok, ok := req.Form["token"]; ok && tok[0] != "" {
			log.Debugf("generating new sentence from %q", tok[0])
			sentence = chain.GenerateFromToken(tok[0])
		}

		// create sentance from skratch
		if sentence == "" {
			log.Debug("generate new sentence from scratch")
			sentence = chain.Generate()
		}

		log.Infof("%q", sentence)

		err = a.cfg.Templates.ExecuteTemplate(w, "index.html.tmpl", IndexContext{
			CssProperties: a.getCSSCustomProperties(sentence),
			Sentence:      sentence,
			Tokens:        strings.Split(sentence, " "),
		})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

	}
}
