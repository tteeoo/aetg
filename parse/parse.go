package parse

import (
	"errors"
	"github.com/tteeoo/aetg/exp"
	"strconv"
)

func newOp(symbol string, l exp.Exp, r exp.Exp) (exp.Exp, error) {
	switch symbol {
	case "+":
		return exp.Add{L: l, R: r}, nil
	case "-":
		return exp.Sub{L: l, R: r}, nil
	case "*":
		return exp.Mul{L: l, R: r}, nil
	case "/":
		return exp.Div{L: l, R: r}, nil
	case "^":
		return exp.Pow{L: l, R: r}, nil
	case "%":
		return exp.Mod{L: l, R: r}, nil
	case "|":
		return exp.Rad{L: l, R: r}, nil
	default:
		return nil, errors.New("invalid operator symbol: " + string(symbol))
	}
}

// GetExpTree takes a string and parses it into a tree of expression structs recursively
func GetExpTree(strExp string) (exp.Exp, error) {

	// Check if valid
	if len(strExp) < 2 {
		return nil, errors.New("no expression given")
	}
	if strExp[0] != '(' || strExp[len(strExp)-1] != ')' {
		return nil, errors.New("expressions must be surrounded by parenthesis")
	}

	// Tokenize expression
	subExp := 0
	var items []string
	var itemBuf string
	for i, char := range strExp {
		switch char {
		case '(':
			if i != 0 {
				itemBuf += string(char)
				subExp++
			}
		case ')':
			if i != len(strExp)-1 {
				itemBuf += string(char)
				subExp--
			} else if itemBuf != "" {
				items = append(items, itemBuf)
				itemBuf = ""
			}
		case ' ':
			if subExp == 0 {
				if itemBuf != "" {
					items = append(items, itemBuf)
					itemBuf = ""
				}
			} else {
				itemBuf += string(char)
			}
		default:
			itemBuf += string(char)
		}
	}

	// Further checking of validity
	switch len(items) {
	case 0:
		return nil, errors.New("empty expressions are not allowed")

	// Parse single item expression
	case 1:
		f, err := strconv.ParseFloat(items[0], 64)
		if err != nil {
			e, err := GetExpTree(items[0])
			if err != nil {
				return nil, errors.New("invalid expression: " + items[0])
			}
			return e, nil
		}
		return exp.Num{Val: f}, nil

	// Parse operator expression
	case 3:
		var left exp.Exp
		lf, err := strconv.ParseFloat(items[0], 64)
		if err != nil {
			e, err := GetExpTree(items[0])
			if err != nil {
				return nil, errors.New("invalid expression: " + items[0])
			}
			left = e
		} else {
			left = exp.Num{Val: lf}
		}

		var right exp.Exp
		rf, err := strconv.ParseFloat(items[2], 64)
		if err != nil {
			e, err := GetExpTree(items[2])
			if err != nil {
				return nil, errors.New("invalid expression: " + items[2])
			}
			right = e
		} else {
			right = exp.Num{Val: rf}
		}

		e, err := newOp(items[1], left, right)
		if err != nil {
			return nil, err
		}

		return e, nil

	default:
		return nil, errors.New("expressions can only have 1 or 3 items: " + strExp + " has " + strconv.Itoa(len(items)))
	}
}
