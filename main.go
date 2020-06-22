package main

import (
	"flag"
	"fmt"
	"github.com/chzyer/readline"
	"github.com/tteeoo/aetg/parse"
	"io"
	"io/ioutil"
	"os"
)

var version = flag.Bool("version", false, "display version information")
var verbose = flag.Bool("verbose", false, "print expressions after parsing")
var exprMode = flag.String("expr", "", "supply an expression after this flag to run instead of the shell")
var file = flag.String("file", "", "supply a filepath after this flag to read an expression from instead of the shell")

func main() {

	flag.Parse()

	// Print version information if supplied
	if *version {
		fmt.Println("aetg: version: 0.1.0")
		os.Exit(0)
	}

	// Evaluate expression from file if supplied
	if *file != "" {
		b, err := ioutil.ReadFile(*file)
		if err != nil {
			fmt.Println("aetg: error:", err)
			os.Exit(1)
		}

		// Ignore newlines
		var strExp string
		for _, char := range string(b) {
			if char != '\n' {
				strExp += string(char)
			}
		}

		exp, err := parse.GetExpTree(strExp)
		if err != nil {
			fmt.Println("aetg: error:", err)
			os.Exit(1)
		} else {
			if *verbose {
				fmt.Println(exp)
			}
			fmt.Println(exp.Eval())
			os.Exit(0)
		}
	}

	// Evaluate expression if supplied
	if *exprMode != "" {
		exp, err := parse.GetExpTree(*exprMode)
		if err != nil {
			fmt.Println("aetg: error:", err)
			os.Exit(1)
		} else {
			if *verbose {
				fmt.Println(exp)
			}
			fmt.Println(exp.Eval())
			os.Exit(0)
		}
	}

	// Use readline for arrow key movement, etc.
	l, err := readline.NewEx(&readline.Config{
		Prompt:          "aetg> ",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		fmt.Println("aetg: error:", err)
		os.Exit(1)
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
			if *verbose {
				fmt.Println(exp)
			}
			fmt.Println(exp.Eval())
		}
	}

	os.Exit(0)
}
