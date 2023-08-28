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
	log.SetFlags(0)
	log.SetPrefix("lexer:")

	fmt.Println("Welcome to goli repl!")
	fmt.Println("Type (exit) to exit!\n")

Repl:
	for {
		fmt.Print("> ")

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
					log.Printf("%d:%d: Unmatched token", token.Pos.Start[0], token.Pos.Start[1])
					break
				}
				if err.Error() == "Empty" {
					log.Printf(" Success!\n")
					break TokenizeLine
				}
			}
			log.Printf("%d:%d: <%s>", token.Pos.Start[0], token.Pos.Start[1], token.Type)
		}

	}
}
