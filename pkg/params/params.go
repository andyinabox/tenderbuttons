package params

import "math/rand"

type Params struct {
	r *rand.Rand
}

func New(seed []byte) *Params {
	return &Params{GetSeededRandom(seed)}
}

func (p *Params) GetDegree() int {
	return int(p.r.Float32() * 360)
}

func (p *Params) GetComplementaryDegrees() (int, int) {
	d := p.GetDegree()
	return d, 360 - d
}
