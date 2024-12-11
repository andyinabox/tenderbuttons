package params

import "fmt"

type URL string

func (u URL) ToCSS() string {
	return fmt.Sprintf("url('%s')", u)
}
