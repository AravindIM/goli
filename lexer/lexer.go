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

func (t Token) String() string {
	return t.Symbol
}

type Lexer struct {
	code        string
	definitions [][2]string
	count       int64
	cursor      [2]int64
}

func NewLexer(definitions [][2]string) *Lexer {
	return &Lexer{
		definitions: definitions,
		count:       0,
		cursor:      [2]int64{0, 0},
	}
}

func (l *Lexer) Analyze(code string) {
	l.code = code
	l.count = 0
	l.cursor = [2]int64{0, 0}
}

func (l Lexer) Code() string {
	return l.code
}

func (l *Lexer) advanceCount(step int64) {
	l.count += step
}

func (l *Lexer) advanceColumn(step int64) {
	l.cursor[1] += step
}

func (l *Lexer) advanceLine(step int64) {
	l.cursor[0] += step
	l.cursor[1] = 0
}

func (l *Lexer) NextToken() (*Token, error) {
	if len(l.code) == 0 {
		return nil, errors.New("Empty")
	}

	code := []byte(l.code)
	newLine := regexp.MustCompile(`^\n`)
	whiteSpace := regexp.MustCompile(`^\s+`)
	Unmatched := regexp.MustCompile(`^[^\s]+`)

	for loc := newLine.FindIndex(code); loc != nil; {
		l.advanceLine(1)
		l.code = l.code[loc[1]+1:]
		code = []byte(l.code)
	}

	for loc := whiteSpace.FindIndex(code); loc != nil; {
		l.advanceColumn(int64(loc[1] - loc[0] + 1))
		l.code = l.code[loc[1]+1:]
		code = []byte(l.code)
	}

	for _, definition := range l.definitions {
		code = []byte(l.code)
		pattern := regexp.MustCompile(`^` + definition[1])
		if loc := pattern.FindIndex(code); loc != nil {
			start := l.cursor
			l.advanceCount(1)
			l.advanceColumn(int64(loc[1] - loc[0] + 1))
			l.code = l.code[loc[1]+1:]

			return &Token{
				Type:   definition[0],
				Symbol: string(code[loc[0]:loc[1]]),
				Pos: Position{
					Index: l.count,
					Start: start,
					End:   l.cursor,
				},
			}, nil
		}
	}

	code = []byte(l.code)
	loc := Unmatched.FindIndex(code)
	start := l.cursor
	l.advanceCount(1)
	l.advanceColumn(int64(loc[1] - loc[0] + 1))
	if loc[1] < len(l.code) {
		l.code = l.code[loc[1]+1:]
	}

	return &Token{
		Type:   "unmatched",
		Symbol: string(code[loc[0]:loc[1]]),
		Pos: Position{
			Index: l.count,
			Start: start,
			End:   l.cursor,
		},
	}, errors.New("Unmatched")
}
