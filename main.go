package main

import (
	"embed"
	"flag"
	"html/template"
	"io/fs"
	"os"

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
var loadAssetsFromFs bool
var serverRunMode string

func init() {
	flag.BoolVar(&debug, "v", false, "verbose logging")
	flag.BoolVar(&loadAssetsFromFs, "a", false, "load assets from filesystem (for easier frontend development)")
	flag.StringVar(&serverRunMode, "m", "http", "server run mode. can be 'http' or 'https'")
	flag.Parse()

	if debug {
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	var assetsFs fs.FS

	// load assets from filesystem for development
	if loadAssetsFromFs {
		assetsFs = os.DirFS("assets")
		// load assets from embedded data
	} else {
		var err error
		assetsFs, err = fs.Sub(fs.FS(assets), "assets")
		if err != nil {
			log.Fatal(err)
		}
	}

	app := app.New(
		&app.Config{
			Debug:         debug,
			ServerRunMode: serverRunMode,
			Corpus:        corpus,
			Readme:        readme,
			MinWords:      5,
			MaxWords:      35,
			PrefixLength:  3,
			Assets:        assetsFs,
			Templates:     template.Must(template.New("").ParseFS(templates, "tmpl/*.tmpl")),
		},
	)

	log.Info("starting application", "debug", debug, "mode", serverRunMode, "embeddedAsssets", !loadAssetsFromFs)
	log.Fatal(app.Start())

}
