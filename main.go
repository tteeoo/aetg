package main

import (
	"bufio"
	"fmt"
	"github.com/tteeoo/metg/parse"
	"os"
)

func main() {

	for {

		// Take an expression from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("metg> ")
		strExp, err := reader.ReadString('\n')
		strExp = strExp[:len(strExp)-1]
		if err != nil {
			fmt.Println("metg: error:", err)
		}

		// Parse the expression
		exp, err := parse.GetExpTree(strExp)
		if err != nil {
			fmt.Println("metg: error:", err)
		} else {

			// Evaluate the expression and print its value
			fmt.Println(exp.Eval())
		}
	}
}
