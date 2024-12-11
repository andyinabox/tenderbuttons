package approuter

import (
	"html/template"
)

type Config struct {
	Corpus       []byte
	Readme       []byte
	BgImageUrls  []string
	MinWords     int
	MaxWords     int
	PrefixLength int
}

type Router struct {
	tpl *template.Template
	cfg *Config
}

func New(templates *template.Template, cfg *Config) *Router {
	return &Router{templates, cfg}
}
