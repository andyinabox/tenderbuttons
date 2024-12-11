package params

import (
	"fmt"
)

type Degrees int16

func (d Degrees) ToCSS() string {
	return fmt.Sprintf("%ddeg", d)
}

func (d Degrees) Value() int16 {
	return int16(d)
}
