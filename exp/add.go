package exp

import (
	"fmt"
)

// Add represents an addition operation
type Add struct {
	L Exp
	R Exp
}

// Eval returns the value of an Add
func (a Add) Eval() float64 {
	return a.L.Eval() + a.R.Eval()
}

// String returns the string representation of an Add
func (a Add) String() string {
	return fmt.Sprintf("(%s + %s)", a.L, a.R)
}
