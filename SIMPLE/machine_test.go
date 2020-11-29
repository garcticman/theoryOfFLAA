package SIMPLE

import (
	"testing"
)

func TestMachine_Run(t *testing.T) {
	machine := Machine{MultiplyInt{
		left: IntOperand(5),
		right: MultiplyInt{
			left: AddInt{
				left:  IntOperand(2),
				right: IntOperand(-3),
			},
			right: IntOperand(4),
		},
	}, Environment{}}

	machine.Run()
}

func TestMachine_Equal(t *testing.T) {
	machine := Machine{EqualInt{
		left: Variable{Name: "X"},
		right: MultiplyInt{
			left: AddInt{
				left:  IntOperand(2),
				right: Variable{Name: "Y"},
			},
			right: GreaterThanInt{
				left: LessThanInt{
					left:  IntOperand(-1),
					right: BoolOperand(true),
				},
				right: IntOperand(0),
			},
		},
	}, Environment{
		"X": AddInt{
			left:  IntOperand(9),
			right: IntOperand(-8),
		},
		"Y": IntOperand(-1),
	}}

	machine.Run()
}

func Test_Statement(t *testing.T) {
	statement := Operand(Assign{
		variableName: "var",
		Operand: AddInt{
			left:  Variable{Name: "var"},
			right: IntOperand(1),
		},
	})
	environment := Environment{"var": IntOperand(2)}
	machine := Machine{
		Operand:     statement,
		Environment: environment,
	}
	machine.Run()
}

func Test_Sequence(t *testing.T) {
	statement1 := Operand(Assign{
		variableName: "var",
		Operand:      IntOperand(2),
	})
	statement2 := Operand(Assign{
		variableName: "var2",
		Operand:      IntOperand(4),
	})
	statement3 := Operand(Assign{
		"sum",
		AddInt{
			left:  Variable{Name: "var"},
			right: Variable{Name: "var2"},
		},
	})
	environment := Environment{}
	machine := Machine{
		Operand: Sequence{
			first: statement1,
			second: Sequence{
				first:  statement2,
				second: statement3,
			},
		},
		Environment: environment,
	}
	machine.Run()
}

func Test_While(t *testing.T) {
	statement1 := Operand(Assign{
		variableName: "var",
		Operand:      IntOperand(2),
	})
	while := While{
		condition: LessThanInt{
			left:  Variable{"var"},
			right: IntOperand(20),
		},
		body: Assign{
			variableName: "var",
			Operand: MultiplyInt{
				left:  Variable{"var"},
				right: IntOperand(2),
			},
		},
	}
	environment := Environment{}
	machine := Machine{
		Operand: Sequence{
			first:  statement1,
			second: while,
		},
		Environment: environment,
	}
	machine.Run()
}
