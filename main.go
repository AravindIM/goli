package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gitlab.com/AravindIM/goli/compiler"
	"gitlab.com/AravindIM/goli/lexer"
	"gitlab.com/AravindIM/goli/parser"
)

func main() {
	definitions := [][2]string{
		{"start", `\(`},
		{"end", `\)`},
		{"number", `\d+(.\d+)?`},
		{"string", `".*?"|'.*?'`},
		{"symbol", `[^\(\)\s]+`},
	}
	lex := lexer.NewLexer(definitions)

	scanner := bufio.NewScanner(os.Stdin)
	log.SetFlags(0)
	log.SetPrefix("goli:")

	fmt.Print("Welcome to goli repl!\n")
	fmt.Print("Type (exit) to exit!\n\n")

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
		ast, err := parser.Parse(lex)
		if err != nil {
			log.Printf(err.Error())
		} else {
			compiler.Compile(ast)
		}
	}
}
