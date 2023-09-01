package parser

type AstNode interface {
	Type() string
}

type ListNode struct {
	Value AstNode*
	Next AstNode*
}

func (n ListNode) Type () {
	return "list"
}

type StringNode struct {
	Value string
}

func (n StringNode) Type () {
	return "string"
}

type NumberNode struct {
	Value string
}

func (n NumberNode) Type () {
	return "number"
}