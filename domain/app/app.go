package app

import (
	"html/template"
	"io/fs"

	"github.com/andyinaobox/tenderbuttons/pkg/assetshandler"
	"github.com/andyinaobox/tenderbuttons/pkg/tlsserver"
)

type Config struct {
	Debug         bool
	ServerRunMode string
	Corpus        []byte
	Readme        []byte
	MinWords      int
	MaxWords      int
	PrefixLength  int
	Assets        fs.FS
	Templates     *template.Template
}

type App struct {
	server *tlsserver.Server
	cfg    *Config
}

func New(cfg *Config) *App {

	app := &App{
		cfg: cfg,
	}

	handler := assetshandler.New(
		&assetshandler.Config{
			AssetsUrlPath:     "/assets",
			AssetsFS:          cfg.Assets,
			StripAssetsPrefix: true,
		},
	)

	handler.AddRoute("/", app.IndexHandler())
	handler.AddRoute("/readme", app.AboutHandler())

	app.server = tlsserver.New(&tlsserver.Config{
		Handler:      handler,
		RunMode:      tlsserver.RunMode(cfg.ServerRunMode),
		CertFilePath: ".cert/localhost.crt",
		KeyFilePath:  ".cert/localhost.key",
	})

	return app
}

func (a *App) Start() error {
	defer a.server.Close()
	return a.server.Start()
}
