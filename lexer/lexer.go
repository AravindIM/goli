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
	definition map[string]string
	current    int64
}

func (l Lexer) New(definition map[string]string) Lexer {
	return Lexer{definition: make(map[string]string), current: 0}
}
