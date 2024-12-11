package approuter

import (
	"html/template"
	"strings"

	"github.com/andyinaobox/tenderbuttons/pkg/cssprops"
	"github.com/andyinaobox/tenderbuttons/pkg/params"
	"github.com/andyinaobox/tenderbuttons/pkg/randparams"
)

func (r *Router) getCSSCustomProperties(sentence string) template.HTML {

	rng := randparams.New([]byte(sentence))

	// value of 0-1 based on the sentence lenths in words (based on max an min values)
	sentenceSpread := float32(len(strings.Split(sentence, " "))-r.cfg.MinWords) / float32(r.cfg.MaxWords-r.cfg.MinWords)

	h2, h3 := rng.ComplementaryDegrees()

	rs1 := params.Percent(sentenceSpread*20. + 30.)

	la1, la2 := rng.ComplementaryDegrees()

	var lcSaturation float32 = 75.
	var lcLightness float32 = 75.
	var lcStartOpacity float32 = 100.
	var lcEndOpacity float32 = 30.

	lc1 := params.NewColorHSLA(h2.Value(), lcSaturation, lcLightness, lcStartOpacity)
	lc2 := params.NewColorHSLA(h2.Value(), lcSaturation, lcLightness, lcEndOpacity)
	lc3 := params.NewColorHSLA(h3.Value(), lcSaturation, lcLightness, lcStartOpacity)
	lc4 := params.NewColorHSLA(h3.Value(), lcSaturation, lcLightness, lcEndOpacity)

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
