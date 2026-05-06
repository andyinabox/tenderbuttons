package app

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/charmbracelet/log"

	"code.andydayton.com/andy/tenderbuttons/pkg/chains"
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
		var err error
		var token, prev, sentence string

		req.ParseForm()

		if len(req.Form["token"]) > 0 {
			token = req.Form["token"][0]
		}
		if len(req.Form["prev"]) > 0 {
			prev = req.Form["prev"][0]
		}

		log.Debug("index handler", "token", token, "prev", prev)

		// create sentance from token
		if token != "" {
			log.Debugf("generating new sentence from %q", token)

			for i := 0; i < 3; i++ {
				try := chain.GenerateFromToken(token)

				// try to come up with a new sentence
				if strings.TrimSpace(try) != strings.TrimSpace(prev) {
					sentence = try
					break
				}
			}
		}

		// create sentance from scratch
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
