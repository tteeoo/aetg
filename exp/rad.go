package exp

import (
	"fmt"
	"math"
)

// Rad represents a radical operation
type Rad struct {
	L Exp
	R Exp
}

// Eval returns the value of a Rad
func (r Rad) Eval() float64 {
	return math.Pow(r.R.Eval(), 1 / r.L.Eval())
}

// String returns the string representation of a Div
func (r Rad) String() string {
	return fmt.Sprintf("(%s | %s)", r.L, r.R)
}
