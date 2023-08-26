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
	code        string
	definitions []Definition
	count       int64
	cursor      [2]int64
}

func (l Lexer) New(definitions []Definition, code string) Lexer {
	return Lexer{
		code:        code,
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

func (l Lexer) nextToken() (*Token, error) {
	if len(l.code) == 0 {
		return nil, errors.New("Empty")
	}

	for _, definition := range l.definitions {
		pattern := regexp.MustCompile(`^` + definition.Pattern)
		if loc := pattern.FindIndex([]byte(l.code)); loc != nil {
			start := l.cursor
			l.advanceCount(1)
			l.advanceColumn(int64(loc[0]))
			code := l.code
			l.code = l.code[loc[1]:]

			return &Token{
				Type:   definition.Type,
				Symbol: code[loc[0]:loc[1]],
				Pos: Position{
					Index: l.count,
					Start: start,
					End:   l.cursor,
				},
			}, nil
		}
	}

	return nil, errors.New("Unmatched")
}