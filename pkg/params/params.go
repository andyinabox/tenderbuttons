package params

import "math/rand"

type Params struct {
	r *rand.Rand
}

func New(seed []byte) *Params {
	return &Params{GetSeededRandom(seed)}
}

func (p *Params) GetDegree() int16 {
	return int16(p.r.Float32() * 360)
}

func (p *Params) GetComplementaryDegrees() (int16, int16) {
	d := p.GetDegree()
	return d, 360 - d
}

func (p *Params) GetRandomHueHSLA(s, l, a float32) *ColorHSLA {
	h := p.GetDegree()
	return NewColorHSLA(h, s, l, a)
}

func (p *Params) GetComplementaryHSLAColors(s, l, a float32) (*ColorHSLA, *ColorHSLA) {
	h1, h2 := p.GetComplementaryDegrees()
	c1 := NewColorHSLA(h1, s, l, a)
	c2 := NewColorHSLA(h2, s, l, a)
	return c1, c2
}

func (p *Params) GetInt32InRange(min int32, max int32) int32 {
	return int32(p.GetFloat32InRange(float32(min), float32(max)))
}

func (p *Params) GetFloat32InRange(min float32, max float32) float32 {
	n := p.r.Float32() * (max - min)
	return n + min
}
