package randparams

import "github.com/andyinaobox/tenderbuttons/pkg/params"

func (r *Rand) PercentInRange(min float32, max float32) params.Percent {
	return params.Percent(r.Float32InRange(min, max))
}
