package SIMPLE

import "testing"

func Test_Condition(t *testing.T) {
	variable := Operand(LessThanInt{
		left: MultiplyInt{
			left:  IntOperand(4),
			right: IntOperand(2),
		},
		right: AddInt{
			left:  IntOperand(4),
			right: IntOperand(2),
		},
	})
	consequence := Operand(Assign{
		variableName: "var1",
		Operand: AddInt{
			left:  IntOperand(4),
			right: IntOperand(2),
		},
	})
	statement := Operand(If{
		condition:   variable,
		consequence: consequence,
		alternative: SentenceEnd{BoolOperand(false)},
	})

	environment := Environment{"var": IntOperand(2)}
	machine := Machine{
		Operand:     statement,
		Environment: environment,
	}
	machine.Run()
}
