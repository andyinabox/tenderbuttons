package props

import (
	"html/template"

	"github.com/andyinaobox/tenderbuttons/pkg/cssprops"
	"github.com/andyinaobox/tenderbuttons/pkg/params"
	"github.com/andyinaobox/tenderbuttons/pkg/randparams"
)

func GetCSSCustomProperties(sentence string) template.HTML {

	r := randparams.New([]byte(sentence))

	rs1 := r.PercentInRange(30., 50.)
	la1, la2 := r.ComplementaryDegrees()
	lc1, lc3 := r.ComplementaryHSLAColors(75., 75., 100.)
	lc2 := params.NewColorHSLA(lc1.H.Value(), lc1.S.Value(), lc1.L.Value(), 0.)
	lc4 := params.NewColorHSLA(lc3.H.Value(), lc3.S.Value(), lc3.L.Value(), 0.)

	d := &cssprops.Properties{
		"--radial-stop-1":  rs1,
		"--linear-angle-1": la1,
		"--linear-color-1": lc1,
		"--linear-color-2": lc2,
		"--linear-angle-2": la2,
		"--linear-color-3": lc3,
		"--linear-color-4": lc4,
	}

	return d.ToHTML()
}
