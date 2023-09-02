package parser

import "gitlab.com/AravindIM/goli/lexer"

type AstNode interface {
	Type() string
}

type ListNode struct {
	Value *AstNode
	Next  *AstNode
	Pos   lexer.Position
}

func (n ListNode) Type() string {
	return "list"
}

type StringNode struct {
	Value string
	Pos   lexer.Position
}

func (n StringNode) Type() string {
	return "string"
}

type NumberNode struct {
	Value string
	Pos   lexer.Position
}

func (n NumberNode) Type() string {
	return "number"
}

