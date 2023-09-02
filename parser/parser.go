package parser

import (
	"container/list"
	"log"

	"gitlab.com/AravindIM/goli/lexer"
)

type AstNode interface {
	Type() string
}

type ListNode struct {
	Value *AstNode
	Next  *AstNode
	Pos   lexer.Position
}

func (n ListNode) Type() string {
	return "list"
}

type StringNode struct {
	Value string
	Pos   lexer.Position
}

func (n StringNode) Type() string {
	return "string"
}

type NumberNode struct {
	Value string
	Pos   lexer.Position
}

func (n NumberNode) Type() string {
	return "number"
}

func Parse(lex lexer.Lexer) {
	stack := list.New()
	s := stack.Back()

	log.SetFlags(0)
	log.SetPrefix("goli:")

TokenizeLine:
	for {
		token, err := lex.NextToken()
		if err != nil {
			if err.Error() == "Unmatched" {
				break
			}
			if err.Error() == "Empty" {
				break TokenizeLine
			}
		}
		if token.Type == "start" {
			node := &ListNode{
				Pos: lexer.Position{
					Start: token.Pos.Start,
				},
			}
			stack.PushBack(node)
		}

		if token.Type == "string" {
			listNode := s.Prev()
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
