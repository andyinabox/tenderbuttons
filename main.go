package main

import (
	"embed"
	"flag"
	"html/template"

	"github.com/charmbracelet/log"

	"github.com/andyinaobox/tenderbuttons/domain/router"
	"github.com/andyinaobox/tenderbuttons/pkg/handler"
	"github.com/andyinaobox/tenderbuttons/pkg/server"
)

//go:embed tmpl/*
var templates embed.FS

//go:embed assets/*
var assets embed.FS

//go:embed corpus/tb.txt
var corpus string

//go:embed README.md
var readme []byte

func main() {
	var debug bool
	var serverRunMode string

	flag.BoolVar(&debug, "v", false, "verbose logging")
	flag.StringVar(&serverRunMode, "m", "http", "server run mode. can be 'http', 'https-ss', 'https-le'")

	flag.Parse()

	if debug {
		log.SetLevel(log.DebugLevel)
	}

	log.Info("starting application", "debug", debug, "mode", serverRunMode)

	// compile templates
	log.Info("compiling templates")
	tpl, err := template.New("").ParseFS(templates, "tmpl/*.tmpl")
	if err != nil {
		panic(err)
	}

	// create new router
	r := router.New(tpl)

	// create request handler
	h := handler.New(

		// config assets filesystem server
		&handler.AssetsConfig{
			UrlPath: "/assets",
			SubPath: "assets",
			FS:      assets,
		},

		// configure routes
		[]handler.Route{
			{
				Path:        "/",
				HandlerFunc: r.IndexHandler(corpus),
			},
			{
				Path:        "/readme",
				HandlerFunc: r.AboutHandler(readme),
			},
		},
	)

	// configure server
	sc := &server.Config{
		Handler: h,
		RunMode: server.RunMode(serverRunMode),
	}

	// create server
	srv := server.New(sc)

	defer srv.Close()
	log.Fatal(srv.Start())
}
