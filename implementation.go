package lab1

import (
  "container/list"
  "strings"
  "errors"
)

func PrefixToInfix(input string) (string, error) {
  stack := &Stack{
    stack: list.New(),
  }

  length := len(input)

  if length == 0 {
    return "", errors.New("Input must be non empty string")
  }

  output := ""

  for i := length - 1; i >= 0; i-- {
    if strings.ContainsAny(string(input[i]), "+&-&*&/&^") {
      op1, _ := stack.Front()
      stack.Pop()
      op2, _ := stack.Front()
      stack.Pop()

      stack.Push("(" + op1 + string(input[i]) + op2 + ")")
    } else {
      stack.Push(string(input[i]))
    }
  }
  output, _ = stack.Front()
  return output, nil
}
