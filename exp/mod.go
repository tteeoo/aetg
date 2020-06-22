package exp

import (
	"fmt"
	"math"
)

// Mod represents a modulus operation
type Mod struct {
	L Exp
	R Exp
}

// Eval returns the value of a Mod
func (m Mod) Eval() float64 {
	return float64(int(math.Round(m.L.Eval())) % int(math.Round(m.R.Eval())))
}

// String returns the string representation of a Mod
func (m Mod) String() string {
	return fmt.Sprintf("(%s %% %s)", m.L, m.R)
}
