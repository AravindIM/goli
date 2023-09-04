package parser

import "gitlab.com/AravindIM/goli/lexer"

type AstNode struct {
	ntype    string
	isList   bool
	list     *AstNode
	element  string
	position lexer.Position
	parent   *AstNode
	next     *AstNode
}

func NewListNode(ntype string, parent *AstNode) *AstNode {
	return &AstNode{
		ntype:  ntype,
		isList: true,
		parent: parent,
	}
}

func NewElementNode(ntype string, parent *AstNode) *AstNode {
	return &AstNode{
		ntype:  ntype,
		isList: false,
		parent: parent,
	}
}

func (n AstNode) Type() string {
	return n.ntype
}

func (n AstNode) IsList() bool {
	return n.isList
}

func (n AstNode) List() *AstNode {
	return n.list
}

func (n *AstNode) SetList(list *AstNode) {
	if n.isList == true {
		n.list = list
	}
}

func (n AstNode) Element() string {
	return n.element
}

func (n *AstNode) SetElement(element string) {
	if n.isList == false {
		n.element = element
	}
}

func (n AstNode) Position() lexer.Position {
	return n.position
}

func (n *AstNode) SetStart(start [2]int64) {
	n.position.Start = start
}

func (n *AstNode) SetEnd(end [2]int64) {
	n.position.End = end
}

func (n AstNode) Parent() *AstNode {
	return n.parent
}

func (n *AstNode) SetParent(parent *AstNode) {
	n.parent = parent
}

func (n AstNode) Next() *AstNode {
	return n.next
}

func (n *AstNode) SetNext(next *AstNode) {
	n.next = next
}
