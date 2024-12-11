package router

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/charmbracelet/log"

	"github.com/andyinaobox/tenderbuttons/domain/props"
	"github.com/andyinaobox/tenderbuttons/pkg/chains"
)

const (
	maxWords  = 100
	prefixLen = 2
)

type IndexContext struct {
	CssProperties template.HTML
	Sentence      string
	Tokens        []string
}

func (r *Router) IndexHandler() http.HandlerFunc {
	// create markov chain
	log.Info("building chain", "prefixLen", prefixLen)
	chain := chains.NewChain(2)
	chain.BuildFromString(string(r.data.Corpus))

	return func(w http.ResponseWriter, req *http.Request) {
		log.Debug("index handler")
		var err error
		var sentence string

		req.ParseForm()

		// create sentance from token
		if tok, ok := req.Form["token"]; ok && tok[0] != "" {
			log.Debugf("generating new sentence from %q", tok[0])
			sentence = chain.GenerateFromToken(tok[0], maxWords)
		}

		// create sentance from skratch
		if sentence == "" {
			log.Debug("generate new sentence from scratch")
			sentence = chain.Generate(maxWords)
		}

		log.Infof("%q", sentence)

		err = r.tpl.ExecuteTemplate(w, "index.html.tmpl", IndexContext{
			CssProperties: props.GetCSSCustomProperties(sentence),
			Sentence:      sentence,
			Tokens:        strings.Split(sentence, " "),
		})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

	}
}
