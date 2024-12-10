package router

import (
	"fmt"
	"html/template"

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

func NewDisplayParams(sentence string) *DisplayParams {

	p := params.New([]byte(sentence))

	rs1 := p.GetFloat32InRange(30., 50.)
	la1, la2 := p.GetComplementaryDegrees()
	lc1, lc3 := p.GetComplementaryHSLAColors(75., 75., 100.)
	lc2 := params.NewColorHSLA(lc1.H, lc1.S, lc1.L, 0.)
	lc4 := params.NewColorHSLA(lc3.H, lc3.S, lc3.L, 0.)

	d := &DisplayParams{
		RadialStop1:  template.CSS(fmt.Sprintf("%.2f%%", rs1)),
		LinearAngle1: template.CSS(fmt.Sprintf("%ddeg", la1)),
		LinearColor1: lc1.ToCSS(),
		LinearColor2: lc2.ToCSS(),
		LinearAngle2: template.CSS(fmt.Sprintf("%ddeg", la2)),
		LinearColor3: lc3.ToCSS(),
		LinearColor4: lc4.ToCSS(),
	}

	// log.Debugf("displayParams: %v", d)

	return d
}
