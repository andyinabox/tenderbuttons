package randparams

import "github.com/andyinaobox/tenderbuttons/pkg/params"

func (r *Rand) Degrees() params.Degrees {
	return params.Degrees(r.Float32() * 360)
}

func (r *Rand) ComplementaryDegrees() (params.Degrees, params.Degrees) {
	d := r.Degrees()
	return d, 360 - d
}
