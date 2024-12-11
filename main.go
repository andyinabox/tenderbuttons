package main

import (
	"embed"
	"flag"
	"html/template"

	"github.com/charmbracelet/log"

	"github.com/andyinaobox/tenderbuttons/domain/app"
)

//go:embed tmpl/*
var templates embed.FS

//go:embed assets/*
var assets embed.FS

//go:embed corpus/tb.txt
var corpus []byte

//go:embed README.md
var readme []byte

var debug bool
var serverRunMode string

func init() {
	flag.BoolVar(&debug, "v", false, "verbose logging")
	flag.StringVar(&serverRunMode, "m", "http", "server run mode. can be 'http', 'https-ss'")
	flag.Parse()

	if debug {
		log.SetLevel(log.DebugLevel)
	}
}

func main() {

	app := app.New(
		&app.Config{
			Debug:         debug,
			ServerRunMode: serverRunMode,
			Corpus:        corpus,
			Readme:        readme,
			Assets:        assets,
			Templates:     template.Must(template.New("").ParseFS(templates, "tmpl/*.tmpl")),
		},
	)

	log.Fatal(app.Start())

}
