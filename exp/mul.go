package exp

import (
	"fmt"
)

// Mul represents a multiplication operation
type Mul struct {
	L Exp
	R Exp
}

// Eval returns the value of a Mul
func (m Mul) Eval() float64 {
	return m.L.Eval() * m.R.Eval()
}

// String returns the string representation of a Mul
func (m Mul) String() string {
	return fmt.Sprintf("(%s * %s)", m.L, m.R)
}
