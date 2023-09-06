package parser

import (
	"errors"

	"gitlab.com/AravindIM/goli/lexer"
)

type Parser struct {
	lex *lexer.Lexer
}

func NewParser() *Parser {
	return new(Parser)
}

func (p *Parser) Parse(lex *lexer.Lexer) {
	p.lex = lex
}

func (p *Parser) NextExpression() (*AstNode, error) {
	var parent *AstNode
	var previous *AstNode
	var current *AstNode

	if p.lex == nil {
		return nil, errors.New("Lexer was not found!")
	}

ParseLoop:
	for {
		token, err := p.lex.NextToken()
		if err != nil {
			if err.Error() == "Unmatched" {
				return nil, errors.New("Unmatched symbol found")
			}
			if err.Error() == "Empty" {
				break
			}
		}

		switch token.Type {
		case "start":
			current = NewListNode("list", parent)
			break
		case "end":
			if parent == nil {
				return nil, errors.New("Extra list closing found")
			}
			parent = parent.Parent()
			if parent == nil {
				break ParseLoop
			}
			current = nil
			break
		case "symbol":
			current = NewElementNode("symbol", parent)
			err = current.SetElement(token.Symbol)
			break
		case "string":
			current = NewElementNode("string", parent)
			err = current.SetElement(token.Symbol)
			break
		case "number":
			current = NewElementNode("number", parent)
			err = current.SetElement(token.Symbol)
			break
		}

		if err != nil {
			return nil, err
		}

		if current != nil {
			if parent != nil {
				parent.SetList(current)
			}

			if current.IsList() {
				parent = current
			}

			if previous != nil {
				previous.SetNext(current)
			}
		}

		if parent == nil {
			break
		}

		previous = current
	}

	if parent != nil {
		return nil, errors.New("Missing closing of list")
	}

	return current, nil
}
