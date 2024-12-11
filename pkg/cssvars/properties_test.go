package cssvars

import (
	"bytes"
	"html/template"
	"testing"
)

func TestCssPropertyNameRegExp(t *testing.T) {
	testData := map[string]bool{
		"--test-var":  true,
		"--Test9-var": true,
		"-test":       false,
		"test-var":    false,
	}

	var result bool
	for k, expected := range testData {
		result = cssPropertyNameRegExp.Match([]byte(k))
		if result != expected {
			t.Errorf("unexpected result for %q; got %v, expected %v", k, result, expected)
		}
	}

}

func TestToHTMLTemplate(t *testing.T) {
	expected := "<style>\n\t:root {\n\n\t\t--var-one: 2px;\n\t\t--var-two: url('image.png');\n\t\t\n\t}\n</style>"

	vars := map[template.CSS]template.CSS{}
	vars[template.CSS("--var-one")] = template.CSS("2px")
	vars[template.CSS("--var-two")] = template.CSS("url('image.png')")

	buf := bytes.NewBuffer([]byte{})
	templates.ExecuteTemplate(buf, "toHTML", toHTMLTemplateContext{
		PropertyAssignments: vars,
	})

	result := buf.String()
	if result != expected {
		t.Errorf("expected:\n%s\n%q\n, got:\n%s\n%q\n", expected, expected, result, result)
	}

}

type cssValueTest string

func (v cssValueTest) ToCSS() template.CSS {
	return template.CSS(v)
}

func TestToHTML(t *testing.T) {
	expected := "<style>\n\t:root {\n\n\t\t--var-one: 2px;\n\t\t--var-two: url('image.png');\n\t\t\n\t}\n</style>"

	vars := &Properties{
		"--var-one": cssValueTest("2px"),
		"--var-two": cssValueTest("url('image.png')"),
	}

	result := string(vars.ToHTML())

	if result != expected {
		t.Errorf("expected:\n%s\n%q\n, got:\n%s\n%q\n", expected, expected, result, result)
	}

}
