package params

import (
	"html/template"
	"testing"
)

func TestColorRGBA(t *testing.T) {
	c := NewColorRGB(0, 100, 200)
	expected := template.CSS("rgb(0 100 200 / 100.00%)")
	result := c.ToCSS()
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestColorHSLA(t *testing.T) {
	c := NewColorHSL(100, 50., 50.)
	expected := template.CSS("hsl(100 50.00% 50.00% / 100.00%)")
	result := c.ToCSS()
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}
