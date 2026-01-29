package main

import (
	"fmt"
)

type Parser struct {
	tokens []Token
	pos    int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens, pos: 0}
}

func (p *Parser) current() Token {
	if p.pos < len(p.tokens) {
		return p.tokens[p.pos]
	}
	return Token{Type: EOF, Lit: ""}
}

func (p *Parser) eat(t TokenType) Token {
	cur := p.current()
	if cur.Type != t {
		panic(fmt.Sprintf("Parser: Expected %s, got %s, at position: %v", t, cur.Type, p.pos))
	}
	p.pos++
	return cur
}

func (p *Parser) match(t TokenType) bool {
	if p.current().Type == t {
		p.pos++
		return true
	}
	return false
}

// isFuncDecl looks ahead to determine if the upcoming tokens form a function declaration
func (p *Parser) isFuncDecl() bool {
	if p.current().Type != Identifier {
		return false
	}
	if p.pos+1 >= len(p.tokens) || p.tokens[p.pos+1].Type != LParen {
		return false
	}
	depth := 0
	for i := p.pos + 1; i < len(p.tokens); i++ {
		if p.tokens[i].Type == LParen {
			depth++
		} else if p.tokens[i].Type == RParen {
			depth--
			if depth == 0 {
				if i+1 < len(p.tokens) && p.tokens[i+1].Type == Equals {
					return true
				}
				return false
			}
		}
	}
	return false
}

func (p *Parser) ParseProgram() *Program {
	body := []Node{}

	for p.current().Type != EOF {
		stmt := p.parseTopLevel()
		// semicolon optional at end of input
		if p.current().Type == Semicolon {
			p.eat(Semicolon)
		} else if p.current().Type != EOF {
			panic(fmt.Sprintf("Parser: Expected Semicolon, got %s, at position: %v", p.current().Type, p.pos))
		}
		body = append(body, stmt)
	}

	return &Program{Body: body}
}

func (p *Parser) parseTopLevel() Node {
	if p.isFuncDecl() {
		// func IDENT(params...) = expr
		nameTok := p.eat(Identifier)
		p.eat(LParen)
		params := []string{}
		if p.current().Type != RParen {
			paramTok := p.eat(Identifier)
			params = append(params, paramTok.Lit)
			for p.match(Comma) {
				paramTok := p.eat(Identifier)
				params = append(params, paramTok.Lit)
			}
		}
		p.eat(RParen)
		if p.match(Equals) {
			body := p.parseExpression()
			return &FuncDecl{Name: nameTok.Lit, Params: params, Body: body}
		} else {
			panic(fmt.Sprintf("Parser: Invalid function declaration, expected '=' after parameters at position: %v", p.pos))
		}
	}

	if p.match(Var) {
		// var IDENT = expr
		nameTok := p.eat(Identifier)
		var init Node = nil
		if p.match(Equals) {
			init = p.parseExpression()
		}
		return &VarDecl{Name: nameTok.Lit, Init: init}
	}

	// expression statement
	expr := p.parseExpression()
	return &ExprStmt{Expr: expr}
}

// Expression parsing with basic precedence
func (p *Parser) parseExpression() Node {
	left := p.parseTerm()
	for p.current().Type == Plus || p.current().Type == Minus {
		op := p.current()
		p.pos++
		right := p.parseTerm()
		left = &BinaryExpr{Op: op.Lit, Left: left, Right: right}
	}
	return left
}

func (p *Parser) parseTerm() Node {
	left := p.parseFactor()
	for p.current().Type == Times || p.current().Type == Divide || p.current().Type == Mod {
		op := p.current()
		p.pos++
		right := p.parseFactor()
		left = &BinaryExpr{Op: op.Lit, Left: left, Right: right}
	}
	return left
}

func (p *Parser) parseFactor() Node {
	// prefix sqrt '@' operator
	if p.match(Sqrt) {
		expr := p.parseFactor()
		return &UnaryExpr{Op: "sqrt", Expr: expr}
	}

	// unary minus
	if p.match(Minus) {
		expr := p.parseFactor()
		return &UnaryExpr{Op: "-", Expr: expr}
	}

	return p.parsePrimary()
}

func (p *Parser) parsePrimary() Node {
	cur := p.current()
	switch cur.Type {
	case Number:
		p.pos++
		return &NumberLiteral{Value: cur.Lit}
	case Identifier:
		p.pos++
		// function call? IDENT (args...)
		if p.current().Type == LParen {
			p.pos++
			args := []Node{}
			if p.current().Type != RParen {
				args = append(args, p.parseExpression())
				for p.match(Comma) {
					args = append(args, p.parseExpression())
				}
			}
			p.eat(RParen)
			return &CallExpr{Callee: &Ident{Name: cur.Lit}, Args: args}
		}
		return &Ident{Name: cur.Lit}
	case LParen:
		p.pos++
		expr := p.parseExpression()
		p.eat(RParen)
		return expr
	default:
		panic(fmt.Sprintf("Parser: Unexpected token in primary: %v, position: %v", cur, p.pos))
	}
}

func Parse(tokens []Token) *Program {
	p := NewParser(tokens)
	return p.ParseProgram()
}
