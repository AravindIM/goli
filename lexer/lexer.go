package lexer

import (
	"errors"
	"regexp"
)

type Position struct {
	Index int64
	Start [2]int64
	End   [2]int64
}

type Token struct {
	Type   string
	Symbol string
	Pos    Position
}

type Definition struct {
	Type    string
	Pattern string
}

type Lexer struct {
	definitions []Definition
	count       int64
	cursor      [2]int64
}

func (l Lexer) New(definitions []Definition) Lexer {
	return Lexer{
		definitions: definitions,
		count:       0,
		cursor:      [2]int64{0, 0},
	}
}

func (l Lexer) advanceCount(step int64) {
	l.count += step
}

func (l Lexer) advanceColumn(step int64) {
	l.cursor[0] += step
}

func (l Lexer) advanceLine(step int64) {
	l.cursor[0] += step
	l.cursor[0] = 0
}
