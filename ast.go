package main

// AST node types

type Node interface{}

type Program struct {
	Body []Node
}

type VarDecl struct {
	Name string
	Init Node // expression or nil
}

type ExprStmt struct {
	Expr Node
}

type NumberLiteral struct {
	Value string
}

type Ident struct {
	Name string
}

type BinaryExpr struct {
	Op    string
	Left  Node
	Right Node
}

type FuncDecl struct {
	Name   string
	Params []string
	Body   Node
}

type CallExpr struct {
	Callee Node
	Args   []Node
}

type UnaryExpr struct {
	Op   string
	Expr Node
}
