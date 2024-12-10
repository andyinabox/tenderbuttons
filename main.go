package main

import (
	"embed"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/charmbracelet/log"

	"github.com/andyinaobox/tenderbuttons/pkg/chains"
	"github.com/andyinaobox/tenderbuttons/pkg/handler"
	"github.com/andyinaobox/tenderbuttons/pkg/params"
	"github.com/russross/blackfriday/v2"
)

//go:embed tmpl/*
var templates embed.FS

//go:embed assets/*
var assets embed.FS

//go:embed corpus/tb.txt
var corpus string

//go:embed README.md
var readme []byte

const maxWords = 100
const prefixLen = 2

type displayParams struct {
	LinearAngle1 int
	LinearAngle2 int
}
type indexContext struct {
	DisplayParams *displayParams
	Sentence      string
	Tokens        []string
}

type aboutContext struct {
	Body template.HTML
}

func main() {
	var port int
	var host string
	var debug bool

	flag.IntVar(&port, "p", 8080, "webserver port")
	flag.StringVar(&host, "h", "", "webserver hostname")
	flag.BoolVar(&debug, "v", false, "verbose logging")

	flag.Parse()

	log.Info("starting application", "port", port, "host", host, "debug", debug)

	if debug {
		log.SetLevel(log.DebugLevel)
	}

	// compile templates
	log.Info("compiling templates")
	tpl, err := template.New("").ParseFS(templates, "tmpl/*.tmpl")
	if err != nil {
		panic(err)
	}

	// create markov chain
	log.Info("building chain", "prefixLen", prefixLen)
	chain := chains.NewChain(2)
	chain.BuildFromString(corpus)

	log.Info("parsing readme")
	about := blackfriday.Run(readme)

	// create server
	server := &http.Server{
		Addr: fmt.Sprintf("%s:%d", host, port),
		Handler: handler.New(

			// config assets filesystem server
			&handler.AssetsConfig{
				UrlPath: "/assets",
				SubPath: "assets",
				FS:      assets,
			},

			// configure routes
			[]handler.Route{
				{
					Path: "/",
					HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
						log.Debug("index handler")
						var err error
						var sentence string

						r.ParseForm()

						// create sentance from token
						if tok, ok := r.Form["token"]; ok && tok[0] != "" {
							log.Debugf("generating new sentence from %q", tok[0])
							sentence = chain.GenerateFromToken(tok[0], maxWords)
						}

						// create sentance from skratch
						if sentence == "" {
							log.Debug("generate new sentence from scratch")
							sentence = chain.Generate(maxWords)
						}

						log.Infof("%q", sentence)

						err = tpl.ExecuteTemplate(w, "index.html.tmpl", indexContext{
							DisplayParams: newDisplayParams(sentence),
							Sentence:      sentence,
							Tokens:        strings.Split(sentence, " "),
						})

						if err != nil {
							w.WriteHeader(http.StatusInternalServerError)
						}

					},
				},
				{
					Path: "/readme",
					HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
						log.Debug("about handler")
						err := tpl.ExecuteTemplate(w, "about.html.tmpl", aboutContext{
							Body: template.HTML(about),
						})
						if err != nil {
							w.WriteHeader(http.StatusInternalServerError)
						}
					},
				},
			},
		),
	}

	log.Info("starting server", "host", host, "port", port)
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func newDisplayParams(sentence string) *displayParams {

	p := params.New([]byte(sentence))

	la1, la2 := p.GetComplementaryDegrees()

	return &displayParams{
		LinearAngle1: la1,
		LinearAngle2: la2,
	}
}
