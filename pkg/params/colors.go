package params

import "fmt"

type RGBColor struct {
	R uint8
	G uint8
	B uint8
}

type RGBAColor struct {
	R uint8
	G uint8
	B uint8
	A float32
}

type HSLColor struct {
	H int16
	S float32
	L float32
}

func (c *HSLColor) ToCSS() string {
	return fmt.Sprintf("hsl(%ddeg, %.2f%%, %.2f%%)", c.H, c.S, c.L)
}

type HSLAColor struct {
	H int16
	S float32
	L float32
	A float32
}

func (c *HSLAColor) ToCSS() string {
	return fmt.Sprintf("hsl(%ddeg, %.2f%%, %.2f%% / %.2f%%)", c.H, c.S, c.L, c.A)
}
