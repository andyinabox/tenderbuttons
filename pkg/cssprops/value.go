package cssprops

type Value string

func (v Value) ToCSS() string {
	return string(v)
}
