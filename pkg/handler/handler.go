package handler

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"

	"github.com/charmbracelet/log"
)

type Route struct {
	Path        string
	HandlerFunc http.HandlerFunc
}

type AssetsConfig struct {
	UrlPath string   // url path ("/assets")
	SubPath string   // the part to trim off in fs ("assets")
	FS      embed.FS // embedded fs
}

func New(assetsConfig *AssetsConfig, routes []Route) Handler {

	assetsFs, err := fs.Sub(fs.FS(assetsConfig.FS), assetsConfig.SubPath)
	if err != nil {
		panic(err)
	}

	// create router using an http.ServeMux
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.HandleFunc(route.Path, func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != route.Path {
				log.Infof("attempt to access path %q, redirecting", r.URL.Path)
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			route.HandlerFunc(w, r)
		})
	}

	return Handler{
		mux:           mux,
		assetsConfig:  assetsConfig,
		assetsHandler: http.StripPrefix(assetsConfig.UrlPath, http.FileServer(http.FS(assetsFs))),
	}
}

type Handler struct {
	mux           *http.ServeMux
	assetsConfig  *AssetsConfig
	assetsHandler http.Handler
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if strings.HasPrefix(r.URL.Path, h.assetsConfig.UrlPath) {
		log.Debugf("serve using file server: %s", r.URL)
		h.assetsHandler.ServeHTTP(w, r)
		return
	}

	log.Debugf("serve using ServeMux: %s", r.URL)
	h.mux.ServeHTTP(w, r)
}
