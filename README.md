# aetg
➗ Arithmetic expression tree generator

Inspired by [this video](https://www.youtube.com/watch?v=7tCNu4CnjVc) (though not in Python).

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

`(((8.11 + -7) ^ 3) - (2 * (5.23 / 8)))`

The following operators are valid:
* `+` Addition
* `-` Subtraction
* `*` Multiplication
* `/` Division
* `^` Exponentiation

# License
All files are licensed under the MIT License
