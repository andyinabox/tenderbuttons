package app

import (
	"embed"
	"html/template"

	"github.com/andyinaobox/tenderbuttons/domain/router"
	"github.com/andyinaobox/tenderbuttons/pkg/handler"
	"github.com/andyinaobox/tenderbuttons/pkg/server"
	"github.com/charmbracelet/log"
)

type Config struct {
	Debug         bool
	ServerRunMode string
	Corpus        []byte
	Readme        []byte
	Assets        embed.FS
	Templates     *template.Template
}

type App struct {
	router  *router.Router
	handler *handler.Handler
	server  *server.Server
	cfg     *Config
}

func New(cfg *Config) *App {

	// create router
	r := router.New(cfg.Templates, &router.Data{
		Corpus: cfg.Corpus,
		Readme: cfg.Readme,
	})

	h := handler.New(
		&handler.AssetsConfig{
			UrlPath: "/assets",
			SubPath: "assets",
			FS:      cfg.Assets,
		},
		r.Routes(),
	)

	s := server.New(&server.Config{
		Handler: h,
		RunMode: server.RunMode(cfg.ServerRunMode),
	})

	return &App{r, h, s, cfg}
}

func (a *App) Start() error {
	log.Info("starting application", "debug", a.cfg.Debug, "mode", a.cfg.ServerRunMode)
	defer a.server.Close()
	return a.server.Start()
}
