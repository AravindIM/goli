package parser

import (
	"errors"

	"gitlab.com/AravindIM/goli/lexer"
)

func Parse(lex *lexer.Lexer) (*Ast, error) {
	var list *AstNode

	ast := NewAst()

	if lex == nil {
		return nil, errors.New("Lexer was not found!")
	}

	for {
		token, err := lex.NextToken()
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
			current := NewListNode("list", list)
			if list != nil {
				list.Push(current)
			} else {
				ast.AppendExpression(current)
			}
			list = current
			break
		case "end":
			if list == nil {
				return nil, errors.New("Extra list closing found")
			}
			list = list.Parent()
			break
		case "symbol":
			current := NewSymbolNode("symbol", token.Symbol, list)
			if list != nil {
				list.Push(current)
			} else {
				ast.AppendExpression(current)
			}
			break
		case "string":
			current := NewSymbolNode("string", token.Symbol, list)
			if list != nil {
				list.Push(current)
			} else {
				ast.AppendExpression(current)
			}
			break
		case "number":
			current := NewSymbolNode("number", token.Symbol, list)
			if list != nil {
				list.Push(current)
			} else {
				ast.AppendExpression(current)
			}
			break
		}
	}

	if list != nil {
		return nil, errors.New("Missing closing of list")
	}

	return ast, nil
}
