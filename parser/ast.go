package parser

import "errors"

type Ast struct {
	head *AstNode
	tail *AstNode
}

func NewAst() *Ast {
	return new(Ast)
}

func (a *Ast) AppendExpression(expr *AstNode) {
	if a.head == nil {
		a.head = expr
		a.tail = expr
	} else {
		a.tail.SetNext(expr)
		a.tail = expr
	}
}

func (a *Ast) NextExpression() (*AstNode, error) {
	if a.head == nil {
		return nil, errors.New("No new expression found")
	}
	current := a.head
	a.head = a.head.Next()
	return current, nil
}
