package main

import (
	"bufio"
	"fmt"
	"github.com/sunzenshen/cgotchas/mpc"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	number := mpc.New("number")
	operator := mpc.New("operator")
	expr := mpc.New("expr")
	lispy := mpc.New("lispy")

	language := "" +
		"number   : /-?[0-9]+/                             ; " +
		"operator : '+' | '-' | '*' | '/'                  ; " +
		"expr     : <number> | '(' <operator> <expr>+ ')'  ; " +
		"lispy    : /^/ <operator> <expr>+ /$/             ; "

	mpc.MpcaLang(language, number, operator, expr, lispy)
	defer mpc.Cleanup(number, operator, expr, lispy)

	for {
		// Prompt
		fmt.Print("input> ")
		// Read a line of user input
		scanner.Scan()
		input := scanner.Text()
		// Print AST using the lispy parser
		mpc.PrintAst(input, lispy)
	}
}
