package exp

// Exp describes a mathematical expression
type Exp interface {
	Eval() float64
	String() string
}
