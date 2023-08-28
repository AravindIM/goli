package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gitlab.com/AravindIM/goli/lexer"
)

func main() {
	definitions := [][2]string{
		{"start", `\(`},
		{"end", `\)`},
		{"function", `fn`},
	}
	lex := lexer.NewLexer(definitions)

	scanner := bufio.NewScanner(os.Stdin)

Repl:
	for {
		fmt.Printf(">")

		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		code := scanner.Text()

		if code == "(exit)" {
			break Repl
		}

		lex.Analyze(code)

	TokenizeLine:
		for {
			token, err := lex.NextToken()
			if err != nil {
				if err.Error() == "Unmatched" {
					log.Printf("Unmatched token at line %d and column %d\n", token.Pos.Start[0], token.Pos.Start[1])
					break
				}
				if err.Error() == "Empty" {
					log.Printf("Finished Tokenizing\n")
					break TokenizeLine
				}
			}
			log.Printf("Token: %s\n", token)
		}

	}
}
