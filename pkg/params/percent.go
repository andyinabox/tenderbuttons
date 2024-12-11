package params

import "fmt"

type Percent float32

func (p Percent) ToCSS() string {
	return fmt.Sprintf("%.2f%%", p)
}

func (p Percent) Value() float32 {
	return float32(p)
}
