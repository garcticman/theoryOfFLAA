package PDA

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := Stack{contents: []int32{'a', 'b', 'c', 'd', 'e'}}
	fmt.Println(string(stack.Top()))
	stack = stack.Pop().Pop()
	fmt.Println(string(stack.Top()))
	stack = stack.Push('x').Push('y')
	fmt.Println(string(stack.Top()))
}
