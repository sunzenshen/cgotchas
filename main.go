package main

import "github.com/sunzenshen/cgotchas/mpc"

func main() {
	number := mpc.New("number")
	operator := mpc.New("operator")
	expr := mpc.New("expr")
	lispy := mpc.New("lispy")
	mpc.Cleanup(number, operator, expr, lispy)
}
