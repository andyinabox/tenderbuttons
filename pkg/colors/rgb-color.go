package colors

type RGBColor struct {
	R int8
	G int8
	B int8
	A float32
}

func NewRGBColor(r int8, g int8, b int8, a float32) RGBColor {
	return RGBColor{r, g, b, a}
}
