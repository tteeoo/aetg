package exp

import (
	"fmt"
)

// Sub represents a subtraction operation
type Sub struct {
	L Exp
	R Exp
}

// Eval returns the value of a Sub
func (s Sub) Eval() float64 {
	return s.L.Eval() - s.R.Eval()
}

// String returns the string representation of a Sub
func (s Sub) String() string {
	return fmt.Sprintf("(%s - %s)", s.L, s.R)
}
