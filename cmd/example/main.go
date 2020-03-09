package main

import (
	"fmt"
	"os"

  . "../.."
)

func main() {
	prefixString := os.Args[1]
  fmt.Println("Your prefix expression:", prefixString)

	infixString, _ := PrefixToInfix(prefixString)

	fmt.Println("Your infix expression:", infixString)
	return
}
