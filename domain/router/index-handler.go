package router

import (
	"net/http"
	"strings"

	"github.com/andyinaobox/tenderbuttons/pkg/chains"
	"github.com/charmbracelet/log"
)

const (
	maxWords  = 100
	prefixLen = 2
)

type IndexContext struct {
	DisplayParams *DisplayParams
	Sentence      string
	Tokens        []string
}

func (r *Router) IndexHandler(corpus string) http.HandlerFunc {
	// create markov chain
	log.Info("building chain", "prefixLen", prefixLen)
	chain := chains.NewChain(2)
	chain.BuildFromString(corpus)

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
			DisplayParams: NewDisplayParams(sentence),
			Sentence:      sentence,
			Tokens:        strings.Split(sentence, " "),
		})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

	}
}
