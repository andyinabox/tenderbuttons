package randparams

import "github.com/andyinaobox/tenderbuttons/pkg/params"

func (r *Rand) HSLARandomHue(s, l, a float32) *params.ColorHSLA {
	h := r.Degrees()
	return params.NewColorHSLA(h.Value(), s, l, a)
}

func (r *Rand) ComplementaryHSLAColors(s, l, a float32) (*params.ColorHSLA, *params.ColorHSLA) {
	h1, h2 := r.ComplementaryDegrees()
	c1 := params.NewColorHSLA(h1.Value(), s, l, a)
	c2 := params.NewColorHSLA(h2.Value(), s, l, a)
	return c1, c2
}
