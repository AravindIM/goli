package lexer

type Position struct {
	Index int
	Start [2]int64
	End   [2]int64
}

type Token struct {
	Type   string
	Symbol string
	Pos    Position
}

type Lexer struct {
	definitions [][]string
	cursor      int64
}

func (l Lexer) New(definitions [][]string) Lexer {
	return Lexer{definitions: definitions, cursor: 0}
}

func (l Lexer) Advance(step int64) {
	l.cursor += step
}
