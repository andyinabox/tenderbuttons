package params

import (
	"fmt"
	"html/template"
)

type ColorRGBA struct {
	R uint8
	G uint8
	B uint8
	A float32
}

func (c *ColorRGBA) ToCSS() template.CSS {
	return template.CSS(fmt.Sprintf("rgb(%d %d %d / %.2f%%)", c.R, c.G, c.B, c.A))
}

func NewColorRGB(r, g, b uint8) *ColorRGBA {
	return NewColorRGBA(r, g, b, 100.0)
}

func NewColorRGBA(r, g, b uint8, a float32) *ColorRGBA {
	return &ColorRGBA{r, g, b, a}
}

type ColorHSLA struct {
	H int16
	S float32
	L float32
	A float32
}

func (c *ColorHSLA) ToCSS() template.CSS {
	return template.CSS(fmt.Sprintf("hsl(%d %.2f%% %.2f%% / %.2f%%)", c.H, c.S, c.L, c.A))
}

func NewColorHSL(h int16, s, l float32) *ColorHSLA {
	return NewColorHSLA(h, s, l, 100.0)
}

func NewColorHSLA(h int16, s, l, a float32) *ColorHSLA {
	return &ColorHSLA{h, s, l, a}
}
