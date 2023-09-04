package parser

import (
	"log"

	"gitlab.com/AravindIM/goli/lexer"
)

func Parse(lex lexer.Lexer) {
	var parent *AstNode
	var previous *AstNode
	var current *AstNode

	log.SetFlags(0)
	log.SetPrefix("goli:")

	for {
		token, err := lex.NextToken()
		if err != nil {
			if err.Error() == "Unmatched" {
				return
			}
			if err.Error() == "Empty" {
				break
			}
		}

		switch token.Type {
		case "start":
			current = NewListNode("list", parent)
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

		if parent != nil {
			parent.SetList(current)
		}

		if current != nil && current.IsList() {
			parent = current
		}

		if previous != nil {
			previous.SetNext(current)
		}

		previous = current
	}

}
