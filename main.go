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
	RadialStop1  template.CSS
	LinearAngle1 template.CSS
	LinearColor1 template.CSS
	LinearColor2 template.CSS
	LinearAngle2 template.CSS
	LinearColor3 template.CSS
	LinearColor4 template.CSS
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

	rs1 := p.GetFloat32InRange(30., 50.)
	la1, la2 := p.GetComplementaryDegrees()
	lc1, lc3 := p.GetComplementaryHSLAColors(75., 75., 100.)
	lc2 := params.NewColorHSLA(lc1.H, lc1.S, lc1.L, 0.)
	lc4 := params.NewColorHSLA(lc3.H, lc3.S, lc3.L, 0.)

	d := &displayParams{
		RadialStop1:  template.CSS(fmt.Sprintf("%.2f%%", rs1)),
		LinearAngle1: template.CSS(fmt.Sprintf("%ddeg", la1)),
		LinearColor1: lc1.ToCSS(),
		LinearColor2: lc2.ToCSS(),
		LinearAngle2: template.CSS(fmt.Sprintf("%ddeg", la2)),
		LinearColor3: lc3.ToCSS(),
		LinearColor4: lc4.ToCSS(),
	}

	log.Debugf("displayParams: %v", d)

	return d
}
