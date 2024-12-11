package cssprops

import (
	"bytes"
	"html/template"
	"strings"
)

const toHTMLTemplate = `
<style>
	:root {

		{{ range $k, $v := . }}{{$k}}: {{$v}};
		{{ end }}
	}
</style>
`

var templates *template.Template

func init() {
	templates = template.Must(template.New("toHTML").Parse(strings.TrimSpace(toHTMLTemplate)))
}

// CSSValue

type CSSValue interface {
	ToCSS() string
}

// Properties

type Properties map[string]CSSValue

func (p Properties) ToHTML() template.HTML {

	buf := bytes.NewBuffer([]byte{})
	templates.ExecuteTemplate(buf, "toHTML", p.toTemplateStrings())

	return template.HTML(buf.String())
}

func (p Properties) toTemplateStrings() (m map[template.CSS]template.CSS) {
	m = make(map[template.CSS]template.CSS)
	for k, v := range p {
		m[template.CSS(k)] = template.CSS(v.ToCSS())
	}
	return m
}
