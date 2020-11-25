package SIMPLE

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	add := MultiplyInt{
		left:  IntOperand(5),
		right: MultiplyInt{
			left:  AddInt{
				left:  IntOperand(2),
				right: IntOperand(3),
			},
			right: IntOperand(4),
		},
	}

	fmt.Println(add.Reduce(nil))
	fmt.Println(add.Int(nil))
}