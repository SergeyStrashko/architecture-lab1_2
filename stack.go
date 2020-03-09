package lab1
import (
    "container/list"
    "fmt"
)

type Stack struct {
    stack *list.List
}

func (c *Stack) Push(value string) {
    c.stack.PushFront(value)
}

func (c *Stack) Pop() error {
    if c.stack.Len() > 0 {
        ele := c.stack.Front()
        c.stack.Remove(ele)
    }
    return fmt.Errorf("Pop Error: Queue is empty")
}

func (c *Stack) Front() (string, error) {
    if c.stack.Len() > 0 {
        if val, ok := c.stack.Front().Value.(string); ok {
            return val, nil
        }
        return "", fmt.Errorf("Peep Error: Queue Datatype is incorrect")
    }
    return "", fmt.Errorf("Peep Error: Queue is empty")
}

func (c *Stack) Size() int {
    return c.stack.Len()
}

func (c *Stack) Empty() bool {
    return c.stack.Len() == 0
}
