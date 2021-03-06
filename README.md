# aetg

Inspired by [this video](https://www.youtube.com/watch?v=7tCNu4CnjVc).

`aetg` is essentially just a simple calculator with an annoying syntax.

This is mainly just a small project I made to get experience with making trees (and parsing) in Go.

# Usage
`aetg` parses arithmetic expressions into a tree of sub expressions that it then evaluates.

Since I am new to writing lexers/parsers, `aetg` uses a very strict format to notate expressions that is easy to tokenize.

Run the compiled `aetg` binary to enter a shell in which you can type expressions.

Notation is as follows:

`(<sub expr> operator <sub expr>)`

Examples:

`(5 * (3 - 6))`

`(((8.11 + -7) ^ 3) - ((2 | 4) * (5.23 / (5 % 2))))`

The following operators are supported:
* `+` Addition
* `-` Subtraction
* `*` Multiplication
* `/` Division
* `^` Exponentiation
* `%` Modulus
* `|` Radical (`(2 | number)` for square root, etc.)

### Flags
* `--version` Display version information
* `--verbose` Print expressions after parsing
* `--file <filepath>` Read and run an expression from a file
* `--expr <expression>` Run a specific expression without going into the shell

# Installation

It's just a single binary.

You can compile with `go build`.

A [Linux binary](https://github.com/tteeoo/aetg/releases/download/0.1.0/aetg-linux) is provided.

# License
All files are licensed under the MIT License
