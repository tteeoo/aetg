package main

import (
	"fmt"
	"github.com/tteeoo/metg/exp"
)

func main() {

	e := exp.Add{
		L: exp.Div{
			L: exp.Num{Val: 12.123123},
			R: exp.Mul{
				L: exp.Sub{
					L: exp.Num{Val: 932.3},
					R: exp.Num{Val: 32.553},
				},
				R: exp.Num{Val: 509},
			},
		},
		R: exp.Num{Val: 4},
	}
	fmt.Println(e)
	fmt.Println(e.Eval())
}
