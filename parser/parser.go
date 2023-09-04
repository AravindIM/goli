package parser

import (
	"errors"

	"gitlab.com/AravindIM/goli/lexer"
)

func Parse(lex *lexer.Lexer) error {
	var parent *AstNode
	var previous *AstNode
	var current *AstNode

	for {
		token, err := lex.NextToken()
		if err != nil {
			if err.Error() == "Unmatched" {
				return errors.New("Unmatched symbol found")
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
			current = nil
			if parent != nil {
				parent = parent.Parent()
			}
			break
		case "symbol":
			current = NewElementNode("symbol", parent)
			current.SetElement(token.Symbol)
			break
		case "string":
			current = NewElementNode("string", parent)
			current.SetElement(token.Symbol)
			break
		case "number":
			current = NewElementNode("number", parent)
			current.SetElement(token.Symbol)
			break
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

		previous = current
	}

	if parent != nil {
		return errors.New("Missing closing of list")
	}

	return nil
}
