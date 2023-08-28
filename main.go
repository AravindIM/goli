package main

import (
	"fmt"
	"gitlab.com/AravindIM/goli/lexer"
)

func main() {
	definitions := [][2]string{
		{"start", `\(`},
		{"end", `\)`},
		{"function", `fn`},
	}
	lex := lexer.NewLexer(definitions)
	var code string

Repl:
	for {
		fmt.Printf(">")
		fmt.Scanf("%s", &code)
		if code == "(exit)" {
			break Repl
		}
		lex.Analyze(code)
	TokenizeLine:
		for {
			token, err := lex.NextToken()
			if err != nil {
				if err.Error() == "Unmatched" {
					fmt.Printf("Unmatched token at line %d and column %d\n", token.Pos.Start[0], token.Pos.Start[1])
					break
				}
				if err.Error() == "Empty" {
					fmt.Printf("Finished Tokenizing\n")
					break TokenizeLine
				}
				fmt.Printf("Token: %s\n", token)
			}
		}

	}
}
