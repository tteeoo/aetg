package exp

import (
	"fmt"
)

// Div represents a division operation
type Div struct {
	L Exp
	R Exp
}

// Eval returns the value of a Div
func (d Div) Eval() float64 {
	return d.L.Eval() / d.R.Eval()
}

// String returns the string representation of a Div
func (d Div) String() string {
	return fmt.Sprintf("(%s / %s)", d.L, d.R)
}
