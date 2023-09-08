package parser

import "errors"

type Ast struct {
	root    *AstNode
	tail    *AstNode
	current *AstNode
}

func NewAst() *Ast {
	return new(Ast)
}

func (a Ast) Root() *AstNode {
	return a.root
}

func (a *Ast) appendExpression(expr *AstNode) {
	if a.root == nil {
		a.root = expr
		a.tail = expr
	} else {
		a.tail.SetNext(expr)
		a.tail = expr
	}
}

func (a *Ast) nextExpression() (*AstNode, error) {
	if a.current == nil {
		a.current = a.root
	} else if a.current.Next() != nil {
		a.current = a.current.Next()
	} else {
		return nil, errors.New("No new expression found")
	}
	return a.current, nil
}
