package cssprops

import (
	"bytes"
	"html/template"
	"testing"
)

func TestToHTMLTemplate(t *testing.T) {
	expected := "<style>\n\t:root {\n\n\t\t--var-one: 2px;\n\t\t--var-two: url('image.png');\n\t\t\n\t}\n</style>"

	vars := map[template.CSS]template.CSS{}
	vars[template.CSS("--var-one")] = template.CSS("2px")
	vars[template.CSS("--var-two")] = template.CSS("url('image.png')")

	buf := bytes.NewBuffer([]byte{})
	templates.ExecuteTemplate(buf, "toHTML", vars)

	result := buf.String()
	if result != expected {
		t.Errorf("expected:\n%s\n%q\n, got:\n%s\n%q\n", expected, expected, result, result)
	}

}

func TestToHTML(t *testing.T) {
	expected := "<style>\n\t:root {\n\n\t\t--var-one: 2px;\n\t\t--var-two: url('image.png');\n\t\t\n\t}\n</style>"

	vars := &Properties{
		"--var-one": Value("2px"),
		"--var-two": Value("url('image.png')"),
	}

	result := string(vars.ToHTML())

	if result != expected {
		t.Errorf("expected:\n%s\n%q\n, got:\n%s\n%q\n", expected, expected, result, result)
	}

}
