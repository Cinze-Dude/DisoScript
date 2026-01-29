package main

import (
	"fmt"
	"unicode"
)

type TokenType string

const (
	Number     TokenType = "Number"
	Identifier TokenType = "Identifier"

	// Keywords
	Var TokenType = "Var"

	// Operators
	Plus   TokenType = "Plus"
	Minus  TokenType = "Minus"
	Times  TokenType = "Times"
	Divide TokenType = "Divide"
	Mod    TokenType = "Mod"
	Sqrt   TokenType = "Sqrt"

	Equals  TokenType = "Equals"
	EqEq    TokenType = "EqEq"
	NotEq   TokenType = "NotEq"
	SolveEq TokenType = "SolveEq"

	LParen    TokenType = "LParen"
	RParen    TokenType = "RParen"
	Comma     TokenType = "Comma"
	Semicolon TokenType = "Semicolon"

	EOF TokenType = "EOF"
)

type Token struct {
	Type TokenType
	Lit  string
}

type Lexer struct {
	input   []rune
	pos     int // current char position
	readPos int // next read position
	ch      rune
}

func NewLexer(s string) *Lexer {
	l := &Lexer{input: []rune(s)}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}
	l.pos = l.readPos
	l.readPos++
}

func (l *Lexer) peekChar() rune {
	if l.readPos >= len(l.input) {
		return 0
	}
	return l.input[l.readPos]
}

func (l *Lexer) skipWhitespace() {
	for l.ch != 0 && unicode.IsSpace(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	start := l.pos
	for l.ch != 0 && (unicode.IsDigit(l.ch) || l.ch == '.') {
		l.readChar()
	}
	return string(l.input[start:l.pos])
}

func (l *Lexer) readIdent() string {
	start := l.pos
	for l.ch != 0 && (unicode.IsLetter(l.ch) || unicode.IsDigit(l.ch) || l.ch == '_') {
		l.readChar()
	}
	return string(l.input[start:l.pos])
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	// helper to return single-char token and advance
	makeTok := func(t TokenType, lit string) Token {
		l.readChar()
		return Token{Type: t, Lit: lit}
	}

	// Multi-char operators first
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			// '=='
			lit := "=="
			l.readChar()
			l.readChar()
			return Token{Type: EqEq, Lit: lit}
		}
		return makeTok(Equals, string(l.ch))
	case '!':
		if l.peekChar() == '=' {
			lit := "!="
			l.readChar()
			l.readChar()
			return Token{Type: NotEq, Lit: lit}
		}
		// fallthrough to unknown
	case '?':
		if l.peekChar() == '=' {
			lit := "?="
			l.readChar()
			l.readChar()
			return Token{Type: SolveEq, Lit: lit}
		}
	}

	switch l.ch {
	case '+':
		return makeTok(Plus, string(l.ch))
	case '-':
		return makeTok(Minus, string(l.ch))
	case '*', '×':
		return makeTok(Times, string(l.ch))
	case '/', '÷':
		return makeTok(Divide, string(l.ch))
	case '%':
		return makeTok(Mod, string(l.ch))
	case '@':
		return makeTok(Sqrt, string(l.ch))
	case '(':
		return makeTok(LParen, string(l.ch))
	case ')':
		return makeTok(RParen, string(l.ch))
	case ',':
		return makeTok(Comma, string(l.ch))
	case ';':
		return makeTok(Semicolon, string(l.ch))
	case 0:
		return Token{Type: EOF, Lit: ""}
	}

	if unicode.IsDigit(l.ch) || (l.ch == '.' && unicode.IsDigit(l.peekChar())) {
		lit := l.readNumber()
		return Token{Type: Number, Lit: lit}
	}

	if unicode.IsLetter(l.ch) {
		lit := l.readIdent()
		if lit == "var" {
			return Token{Type: Var, Lit: lit}
		}
		return Token{Type: Identifier, Lit: lit}
	}

	// Unknown character
	ch := l.ch
	l.readChar()
	return Token{Type: Identifier, Lit: string(ch)}
}

// small debug helper
func (t Token) String() string {
	return fmt.Sprintf("%s(%s)", string(t.Type), t.Lit)
}
