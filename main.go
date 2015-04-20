package main

import (
	"fmt"
	"github.com/sunzenshen/cgotchas/mpc"
)

func main() {
	var number mpc.Parser
	fmt.Println(number)
	number = mpc.New("number")
	fmt.Println(number)
}
