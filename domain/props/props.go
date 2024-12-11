package props

import (
	"html/template"
	"strings"

	"github.com/andyinaobox/tenderbuttons/pkg/cssprops"
	"github.com/andyinaobox/tenderbuttons/pkg/params"
	"github.com/andyinaobox/tenderbuttons/pkg/randparams"
)

func GetCSSCustomProperties(sentence string, minWords int, maxWords int) template.HTML {

	r := randparams.New([]byte(sentence))

	// value of 0-1 based on the sentence lenths in words (based on max an min values)
	sentenceSpread := float32(len(strings.Split(sentence, " "))-minWords) / float32(maxWords-minWords)

	// this should give us a value between 30 and 50 based on
	// size of the sentence in number of words
	rs1 := params.Percent(sentenceSpread*20. + 30.)

	// rs1 := r.PercentInRange(30., 50.)
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
