package main

import "github.com/sunzenshen/cgotchas/mpc"

func main() {
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
	mpc.Cleanup(number, operator, expr, lispy)
}
