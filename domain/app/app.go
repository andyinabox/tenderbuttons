package app

import (
	"html/template"
	"io/fs"

	"github.com/andyinaobox/tenderbuttons/domain/approuter"
	"github.com/andyinaobox/tenderbuttons/pkg/assetshandler"
	"github.com/andyinaobox/tenderbuttons/pkg/tlsserver"
)

type Config struct {
	Debug         bool
	ServerRunMode string
	Corpus        []byte
	Readme        []byte
	BgImageUrls   []string
	MinWords      int
	MaxWords      int
	PrefixLength  int
	Assets        fs.FS
	Templates     *template.Template
}

type App struct {
	router  *approuter.Router
	handler *assetshandler.Handler
	server  *tlsserver.Server
	cfg     *Config
}

func New(cfg *Config) *App {

	// create router
	router := approuter.New(cfg.Templates, &approuter.Config{
		Corpus:       cfg.Corpus,
		Readme:       cfg.Readme,
		BgImageUrls:  cfg.BgImageUrls,
		MinWords:     cfg.MinWords,
		MaxWords:     cfg.MaxWords,
		PrefixLength: cfg.PrefixLength,
	})

	handler := assetshandler.New(
		&assetshandler.Config{
			AssetsUrlPath:     "/assets",
			AssetsFS:          cfg.Assets,
			StripAssetsPrefix: true,
		},
	)

	handler.AddRoute("/", router.IndexHandler())
	handler.AddRoute("/readme", router.AboutHandler())

	server := tlsserver.New(&tlsserver.Config{
		Handler:      handler,
		RunMode:      tlsserver.RunMode(cfg.ServerRunMode),
		CertFilePath: ".cert/localhost.crt",
		KeyFilePath:  ".cert/localhost.key",
	})

	return &App{router, handler, server, cfg}
}

func (a *App) Start() error {
	defer a.server.Close()
	return a.server.Start()
}
