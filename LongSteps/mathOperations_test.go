package LongSteps

import (
	"fmt"
	"testing"
)

func Test_Operations(t *testing.T) {
	fmt.Println(IntNumber(23).Evaluate(nil))
	result := LessThanInt{left: AddInt{
		left:  IntNumber(2),
		right: IntNumber(5),
	},
		right: MultiplyInt{
			left: IntNumber(2),
			right: AddInt{
				left:  IntNumber(2),
				right: IntNumber(3),
			},
		}}
	fmt.Println(result.Evaluate(nil))
}
