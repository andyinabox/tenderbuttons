package main

import (
	_ "embed"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/andyinaobox/tenderbuttons/pkg/chains"
)

//go:embed res/index.html.tmpl
var index string

//go:embed corpus/tb.txt
var corpus string

const tmplName = "index"

type position struct {
	X int
	Y int
}

type renderContext struct {
	Sentence string
	Tokens   []string
	Pos      position
}

func randomIntRange(min int, max int) int {
	offset := rand.Float32() * (float32(max) - float32(min))
	return int(offset) + min
}

func generatePostion() position {
	min := 30
	max := 70
	return position{
		X: randomIntRange(min, max),
		Y: randomIntRange(min, max),
	}
}

func main() {
	var err error

	// create markov chain
	chain := chains.NewChain(2)
	chain.BuildFromString(corpus)

	tmpl, err := template.New(tmplName).Parse(index)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		sentence := chain.Generate(100)

		rc := renderContext{
			Sentence: sentence,
			Tokens:   strings.Split(sentence, " "),
			Pos:      generatePostion(),
		}

		err = tmpl.ExecuteTemplate(w, tmplName, rc)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	addr := "localhost:8080"
	log.Printf("starting server at http://%s", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}
