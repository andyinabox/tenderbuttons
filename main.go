package main

import (
	"embed"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/andyinaobox/tenderbuttons/pkg/chains"
	"github.com/andyinaobox/tenderbuttons/pkg/handler"
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

type indexContext struct {
	Sentence string
	Tokens   []string
}

type aboutContext struct {
	Body template.HTML
}

func main() {
	var port int
	var host string

	flag.IntVar(&port, "p", 8080, "webserver port")
	flag.StringVar(&host, "h", "", "webserverhost")

	// compile templates
	tpl, err := template.New("").ParseFS(templates, "tmpl/*.tmpl")
	if err != nil {
		panic(err)
	}

	// create markov chain
	chain := chains.NewChain(2)
	chain.BuildFromString(corpus)

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
						log.Println("index handler")
						var err error
						var sentence string

						r.ParseForm()

						// create sentance from token
						if tok, ok := r.Form["token"]; ok && tok[0] != "" {
							log.Printf("generating new sentence from %q\n", tok[0])
							sentence = chain.GenerateFromToken(tok[0], maxWords)
						}

						// create sentance from skratch
						if sentence == "" {
							log.Println("generate new sentence from scratch")
							sentence = chain.Generate(maxWords)
						}

						log.Printf("sentence: %q\n", sentence)

						err = tpl.ExecuteTemplate(w, "index.html.tmpl", indexContext{
							Sentence: sentence,
							Tokens:   strings.Split(sentence, " "),
						})

						if err != nil {
							w.WriteHeader(http.StatusInternalServerError)
						}

					},
				},
				{
					Path: "/about",
					HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
						log.Println("about handler")
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

	displayHost := host
	if displayHost == "" {
		displayHost = "localhost"
	}
	log.Printf("starting server at http://%s:%d", displayHost, port)
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
