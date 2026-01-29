package main

import (
	"go/ast"
)

func main() {
	input := "f(x, y) = x*y / 2; f(2, 4)"
	lexer := NewLexer(input)

	tokens := []Token{}
	for {
		tok := lexer.NextToken()
		tokens = append(tokens, tok)
		if tok.Type == EOF {
			break
		}
	}

	program := Parse(tokens)
	ast.Print(nil, program)
}
