package params

import (
	"fmt"
)

type ColorRGBA struct {
	R uint8
	G uint8
	B uint8
	A Percent
}

func (c *ColorRGBA) ToCSS() string {
	return fmt.Sprintf("rgb(%d %d %d / %s)", c.R, c.G, c.B, c.A.ToCSS())
}

func NewColorRGB(r, g, b uint8) *ColorRGBA {
	return NewColorRGBA(r, g, b, 100.0)
}

func NewColorRGBA(r, g, b uint8, a float32) *ColorRGBA {
	return &ColorRGBA{r, g, b, Percent(a)}
}

type ColorHSLA struct {
	H Degrees
	S Percent
	L Percent
	A Percent
}

func (c *ColorHSLA) ToCSS() string {
	return fmt.Sprintf("hsl(%s %s %s / %s)", c.H.ToCSS(), c.S.ToCSS(), c.L.ToCSS(), c.A.ToCSS())
}

func NewColorHSL(h int16, s, l float32) *ColorHSLA {
	return NewColorHSLA(h, s, l, 100.0)
}

func NewColorHSLA(h int16, s, l, a float32) *ColorHSLA {
	return &ColorHSLA{Degrees(h), Percent(s), Percent(l), Percent(a)}
}
