package compiler

import (
	"fmt"

	"gitlab.com/AravindIM/goli/parser"
)

func Compile(ast *parser.Ast) {
	for {
		expr, err := ast.NextExpression()
		if err != nil {
			break
		}

		var list *parser.AstNode

		for expr != nil {
			if expr.IsList() {
				list = expr
			} else {
				symbol, _ := expr.Symbol()
				fmt.Printf("PUSH %s\n", symbol)
			}
			if list != nil {
				expr, _ = list.Pop()
				if expr == nil {
					expr = list.Parent()
					fmt.Printf("CALL\n")
				}
			}
		}
	}
}
