package parser

import (
	"errors"

	"gitlab.com/AravindIM/goli/lexer"
)

type AstNode struct {
	ntype    string
	isList   bool
	list     *AstNode
	length   int64
	symbol   string
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

func NewSymbolNode(ntype string, symbol string, parent *AstNode) *AstNode {
	return &AstNode{
		ntype:  ntype,
		isList: false,
		parent: parent,
		symbol: symbol,
	}
}

func (n AstNode) Type() string {
	return n.ntype
}

func (n AstNode) IsList() bool {
	return n.isList
}

func (n AstNode) Length() (int64, error) {
	if !n.IsList() {
		return 0, errors.New("Not a list!")
	}
	return n.length, nil
}

func (n *AstNode) Push(node *AstNode) error {
	if !n.IsList() {
		return errors.New("Not a list!")
	}
	if n.list == nil {
		node.SetParent(n)
		n.list = node
	} else {
		node.next = n.list
		n.list = node
	}
	n.length += 1
	return nil
}

func (n *AstNode) Pop() (*AstNode, error) {
	if !n.IsList() {
		return nil, errors.New("Not a list!")
	}
	if n.list == nil {
		return nil, errors.New("Empty list!")
	}
	node := n.list
	n.list = n.list.next
	return node, nil
}

func (n AstNode) Symbol() (string, error) {
	if n.IsList() {
		return "", errors.New("Symbol cannot be called on a list")
	}
	return n.symbol, nil
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
