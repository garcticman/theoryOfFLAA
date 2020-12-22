package LongSteps

import (
	"fmt"
	"testing"
)

func TestSequence_Evaluate(t *testing.T) {
	sequence, environment := Sequence{
		left: Assign{
			Name: "var",
			Expression: AddInt{
				left:  IntNumber(2),
				right: IntNumber(-1),
			},
		},
		right: Assign{
			Name: "var2",
			Expression: AddInt{
				left:  Variable{Name: "var"},
				right: IntNumber(4),
			},
		},
	}.Evaluate(Environment{})

	fmt.Println(sequence, environment)
}

func TestWhile_Evaluate(t *testing.T) {
	statement1 := Assign{
		Name:       "var",
		Expression: IntNumber(2),
	}
	while := While{
		condition: LessThanInt{
			left:  Variable{"var"},
			right: IntNumber(20),
		},
		body: Assign{
			Name: "var",
			Expression: MultiplyInt{
				left:  Variable{"var"},
				right: IntNumber(2),
			},
		},
	}

	result, env := Sequence{
		left:  statement1,
		right: while,
	}.Evaluate(Environment{})
	fmt.Println(result, env)
}
