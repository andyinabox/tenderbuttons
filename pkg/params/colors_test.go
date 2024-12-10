package params

import "testing"

func TestHSLColorCSS(t *testing.T) {
	c := HSLColor{180, 50., 50.}
	expected := "hsl(180deg, 50.00%, 50.00%)"

	result := c.ToCSS()
	if result != expected {
		t.Fatalf("expected %q, got %q", expected, result)
	}

}

func TestHSLAColorCSS(t *testing.T) {
	c := HSLAColor{180, 50., 50., 50.}
	expected := "hsl(180deg, 50.00%, 50.00% / 50.00%)"

	result := c.ToCSS()
	if result != expected {
		t.Fatalf("expected %q, got %q", expected, result)
	}

}
