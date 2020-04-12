package main

import (
	"fmt"
	"os"
	
	. "../../build"
	. "../.."
)

func main() {
	prefixString := os.Args[1]
	fmt.Println("Your prefix expression:", prefixString)

	infixString, _ := PrefixToInfix(prefixString)

	fmt.Println("Your infix expression:", infixString)

	fmt.Println(BuildVersion) 
	return
}
