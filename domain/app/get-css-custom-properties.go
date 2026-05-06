package app

import (
	"html/template"
	"strings"

	"code.andydayton.com/andy/tenderbuttons/pkg/cssprops"
	"code.andydayton.com/andy/tenderbuttons/pkg/params"
	"code.andydayton.com/andy/tenderbuttons/pkg/randparams"
)

func (a *App) getCSSCustomProperties(sentence string) template.HTML {

	rng := randparams.New([]byte(sentence))

	// value of 0-1 based on the sentence lenths in words (based on max an min values)
	sentenceSpread := float32(len(strings.Split(sentence, " "))-a.cfg.MinWords) / float32(a.cfg.MaxWords-a.cfg.MinWords)

	h1, h2, h3 := rng.ComplementaryDegreesTriad()

	rs1 := params.Percent(sentenceSpread*20. + 30.)

	la1, la2, la3 := rng.ComplementaryDegreesTriad()

	d := &cssprops.Properties{
		"--rg-stop-1": rs1,

		"--lg-angle-1": la1,
		"--lg-angle-2": la2,
		"--lg-angle-3": la3,

		"--lg-hue-1": h1,
		"--lg-hue-2": h2,
		"--lg-hue-4": h3,
	}

	return d.ToHTML()
}
