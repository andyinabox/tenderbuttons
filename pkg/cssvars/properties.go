package cssvars

import (
	"bytes"
	"html/template"
	"regexp"
	"strings"

	"github.com/charmbracelet/log"
)

const cssPropertyNamePattern = "--[A-z0-9-]+"

var cssPropertyNameRegExp *regexp.Regexp

const toHTMLTemplate = `
<style>
	:root {

		{{ range $k, $v := .PropertyAssignments }}{{$k}}: {{$v}};
		{{ end }}
	}
</style>
`

type toHTMLTemplateContext struct {
	PropertyAssignments map[template.CSS]template.CSS
}

var templates *template.Template

func init() {
	cssPropertyNameRegExp = regexp.MustCompile(cssPropertyNamePattern)
	templates = template.Must(template.New("toHTML").Parse(strings.TrimSpace(toHTMLTemplate)))
}

// CSSValue

type CSSValue interface {
	ToCSS() template.CSS
}

// Properties

type Properties map[string]CSSValue

func (p Properties) ToHTML() template.HTML {

	tplCtx := toHTMLTemplateContext{
		PropertyAssignments: make(map[template.CSS]template.CSS),
	}

	for k, v := range p {

		// skip badly formed var names
		if !cssPropertyNameRegExp.Match([]byte(k)) {
			log.Warnf("propery %q has an invalid format", k)
			continue
		}

		tplCtx.PropertyAssignments[template.CSS(k)] = v.ToCSS()
	}

	buf := bytes.NewBuffer([]byte{})
	templates.ExecuteTemplate(buf, "toCSS", tplCtx)

	return template.HTML(buf.String())
}
