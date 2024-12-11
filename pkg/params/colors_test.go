package params

import (
	"testing"
)

func TestColorRGBA(t *testing.T) {
	c := NewColorRGB(0, 100, 200)
	expected := "rgb(0 100 200 / 100.00%)"
	result := c.ToCSS()
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestColorHSLA(t *testing.T) {
	c := NewColorHSL(100, 50., 50.)
	expected := "hsl(100deg 50.00% 50.00% / 100.00%)"
	result := c.ToCSS()
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}
