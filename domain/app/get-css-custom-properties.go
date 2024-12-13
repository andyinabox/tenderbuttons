package app

import (
	"html/template"
	"strings"

	"gitlab.com/andyinaobox/tenderbuttons/pkg/cssprops"
	"gitlab.com/andyinaobox/tenderbuttons/pkg/params"
	"gitlab.com/andyinaobox/tenderbuttons/pkg/randparams"
)

func (a *App) getCSSCustomProperties(sentence string) template.HTML {

	rng := randparams.New([]byte(sentence))

	// value of 0-1 based on the sentence lenths in words (based on max an min values)
	sentenceSpread := float32(len(strings.Split(sentence, " "))-a.cfg.MinWords) / float32(a.cfg.MaxWords-a.cfg.MinWords)

	h1, h2, h3 := rng.ComplementaryDegreesTriad()

	rs1 := params.Percent(sentenceSpread*20. + 30.)

	la1, la2, la3 := rng.ComplementaryDegreesTriad()

	var lcSaturation float32 = 50.
	var lcLightness float32 = 75.
	var lcStartOpacity float32 = 100.
	var lcEndOpacity float32 = 30.

	lc1 := params.NewColorHSLA(h1.Value(), lcSaturation, lcLightness, lcStartOpacity)
	lc2 := params.NewColorHSLA(h1.Value(), lcSaturation, lcLightness, lcEndOpacity)
	lc3 := params.NewColorHSLA(h2.Value(), lcSaturation, lcLightness, lcStartOpacity)
	lc4 := params.NewColorHSLA(h2.Value(), lcSaturation, lcLightness, lcEndOpacity)
	lc5 := params.NewColorHSLA(h3.Value(), lcSaturation, lcLightness, lcStartOpacity)
	lc6 := params.NewColorHSLA(h3.Value(), lcSaturation, lcLightness, lcEndOpacity)

	d := &cssprops.Properties{
		"--radial-stop-1": rs1,

		"--linear-angle-1": la1,
		"--linear-color-1": lc1,
		"--linear-color-2": lc2,

		"--linear-angle-2": la2,
		"--linear-color-3": lc3,
		"--linear-color-4": lc4,

		"--linear-angle-3": la3,
		"--linear-color-5": lc5,
		"--linear-color-6": lc6,
	}

	return d.ToHTML()
}
