package randparams

import "code.andydayton.com/andy/tenderbuttons/pkg/params"

func (r *Rand) Degrees() params.Degrees {
	return params.Degrees(r.Float32() * 360)
}

func (r *Rand) ComplementaryDegrees() (params.Degrees, params.Degrees) {
	d := r.Degrees()
	return d, d + 180
}

func (r *Rand) ComplementaryDegreesTriad() (params.Degrees, params.Degrees, params.Degrees) {
	d := r.Degrees()
	return d, d + 120, d + 240
}
