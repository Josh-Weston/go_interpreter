package main

// import (
// 	"fmt"
// 	"os"
// 	"os/user"

// 	"github.com/josh-weston/go_interpreter/repl"
// )

// func main() {
// 	user, err := user.Current()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
// 		user.Username)
// 	fmt.Printf("Feel free to type in commands\n")
// 	repl.Start(os.Stdin, os.Stdout)
// }

import (
	"fmt"

	"github.com/josh-weston/go_interpreter/lexer"
	"github.com/josh-weston/go_interpreter/parser"
)

func main() {
	input := `add(1, 2 * 3, add(4, 5))`
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	fmt.Printf("Number of statements: %d\n", len(program.Statements))
	for _, s := range program.Statements {
		fmt.Printf("%+v\n", s.TokenLiteral())
		fmt.Printf("%+v\n", s.String())
	}
}
