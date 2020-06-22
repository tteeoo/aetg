package exp

import (
	"fmt"
	"math"
)

// Pow represents an exponential operation
type Pow struct {
	L Exp
	R Exp
}

// Eval returns the value of a Pow
func (p Pow) Eval() float64 {
	return math.Pow(p.L.Eval(), p.R.Eval())
}

// String returns the string representation of a Pow
func (p Pow) String() string {
	return fmt.Sprintf("(%s ^ %s)", p.L, p.R)
}
