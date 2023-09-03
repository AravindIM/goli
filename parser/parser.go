package parser

import (
	"log"
	"reflect"

	"gitlab.com/AravindIM/goli/lexer"
)

type Content interface {
	Value()
}

type Symbol string

func (s Symbol) Value() string {
	return string(s)
}

type String string

func (s String) Value() string {
	return string(s)
}

type Number float64

func (n Number) Value() float64 {
	return float64(n)
}

type AstNode struct {
	Value  *Content
	Pos    lexer.Position
	Parent *AstNode
	Next   *AstNode
}

func (n AstNode) Type() reflect.Type {
	return reflect.TypeOf(n)
}

func Parse(lex lexer.Lexer) {
	var curr AstNode

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
		if token.Type == "start" {
			curr.Next = &ListNode{
				Pos: lexer.Position{
					Start: token.Pos.Start,
				},
			}
			curr = curr.Next
		}

		if token.Type == "string" {
			prev := curr.Parent
			listNode.Value = &StringNode{
				Value: token.Symbol,
				Pos:   token.Pos,
			}
		}

		if token.Type == "number" {
			listNode := s.Prev()
			listNode.Value = &StringNode{
				Value: token.Symbol,
				Pos:   token.Pos,
			}
		}

		if token.Type == "end" {
			s = s.Prev()
		}
	}

	s = s.Prev()

	if s != nil {
		log.Printf("%d:%d: missing closing of list", s.Pos.Start[0], s.Pos.Start[1])
	}
}
