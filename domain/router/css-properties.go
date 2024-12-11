package router

import (
	"html/template"

	"github.com/andyinaobox/tenderbuttons/pkg/cssprops"
	"github.com/andyinaobox/tenderbuttons/pkg/params"
)

type DisplayParams struct {
	RadialStop1  template.CSS
	LinearAngle1 template.CSS
	LinearColor1 template.CSS
	LinearColor2 template.CSS
	LinearAngle2 template.CSS
	LinearColor3 template.CSS
	LinearColor4 template.CSS
}

func getCSSProperties(sentence string) template.HTML {

	p := params.New([]byte(sentence))

	rs1 := p.GetPercentInRange(30., 50.)
	la1, la2 := p.GetComplementaryDegrees()
	lc1, lc3 := p.GetComplementaryHSLAColors(75., 75., 100.)
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
