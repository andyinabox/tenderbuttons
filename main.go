package main

import (
	"embed"
	"flag"
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"

	"github.com/andyinaobox/tenderbuttons/domain/app"
)

const bgImagesPath = "assets/bg"

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
	var err error
	var assetsFs fs.FS
	var bgImageUrls []string
	var imageEntries []fs.DirEntry

	// load assets from filesystem for development
	if loadAssetsFromFs {

		imageEntries, err = os.ReadDir(bgImagesPath)
		if err != nil {
			log.Fatal(err)
		}

		assetsFs = os.DirFS("assets")
		// load assets from embedded data
	} else {

		imageEntries, err = assets.ReadDir(bgImagesPath)
		if err != nil {
			log.Fatal(err)
		}

		assetsFs, err = fs.Sub(fs.FS(assets), "assets")
		if err != nil {
			log.Fatal(err)
		}
	}

	bgImageUrls = make([]string, 0)
	for _, f := range imageEntries {
		if filepath.Ext(f.Name()) == ".png" {
			bgImageUrls = append(bgImageUrls, fmt.Sprintf("/%s/%s", bgImagesPath, f.Name()))
		}
	}

	app := app.New(
		&app.Config{
			Debug:         debug,
			ServerRunMode: serverRunMode,
			Corpus:        corpus,
			Readme:        readme,
			BgImageUrls:   bgImageUrls,
			MinWords:      5,
			MaxWords:      35,
			PrefixLength:  2,
			Assets:        assetsFs,
			Templates:     template.Must(template.New("").ParseFS(templates, "tmpl/*.tmpl")),
		},
	)

	log.Info("starting application", "debug", debug, "mode", serverRunMode, "embeddedAsssets", !loadAssetsFromFs)
	log.Fatal(app.Start())

}
