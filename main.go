package main

import "github.com/sunzenshen/cgotchas/mpc"

func main() {
	var number mpc.Parser
	number = mpc.New("number")
	mpc.Cleanup(number)
}
