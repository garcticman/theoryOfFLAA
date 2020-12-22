package LongSteps

type Operand interface {
	Evaluate(environment Environment) (Operand, Environment)
	Int(Environment) int
	Bool(Environment) bool
}

type IntNumber int

func (i IntNumber) Evaluate(environment Environment) (Operand, Environment) {
	return i, nil
}
func (i IntNumber) Int(Environment) int {
	return int(i)
}
func (i IntNumber) Bool(environment Environment) bool {
	if int(i) > 0 {
		return true
	} else {
		return false
	}
}

type BoolOperand bool

func (b BoolOperand) Evaluate(environment Environment) (Operand, Environment) {
	return b, nil
}
func (b BoolOperand) Int(Environment) int {
	if b {
		return 1
	} else {
		return 0
	}
}
func (b BoolOperand) Bool(environment Environment) bool {
	return bool(b)
}
