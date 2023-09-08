package parser

type Ast struct {
	root *AstNode
	tail *AstNode
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
