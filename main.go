package main

import (
	"fmt"
	"github.com/chzyer/readline"
	"github.com/tteeoo/aetg/parse"
	"io"
)

func main() {

	// Use readline for arrow key movement, etc.
	l, err := readline.NewEx(&readline.Config{
		Prompt:          "aetg> ",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		fmt.Println("aetg: error:", err)
	}

	for {

		// Take an expression from stdin
		strExp, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(strExp) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		// Parse the expression
		exp, err := parse.GetExpTree(strExp)
		if err != nil {
			fmt.Println("aetg: error:", err)
		} else {

			// Evaluate the expression and print its value
			fmt.Println(exp.Eval())
		}
	}
}
