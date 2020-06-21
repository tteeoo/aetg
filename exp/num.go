package exp

import (
	"strconv"
)

// Num represent a constant number
type Num struct {
	Val float64
}

// Eval returns the value of a Num
func (n Num) Eval() float64 {
	return n.Val
}

// String returns the string representation of a Num
func (n Num) String() string {

	// Find the index of the last non zero character
	var zeroIndex int
	str := strconv.FormatFloat(n.Val, 'f', 13, 64)
	for i, char := range str {
		if char != '0' {
			zeroIndex = i
		}
	}

	// If the string did not end with a non zero character cut of the trailing zeros
	if zeroIndex != len(str) - 1 {

		cut := str[:zeroIndex + 1]

		// Cut of the trailing dot if it exists
		if cut[len(cut) - 1] == '.' {
			return str[:zeroIndex]
		}
		return cut
	}
	return str
}
