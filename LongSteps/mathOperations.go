package LongSteps

type AddInt struct {
	left  Operand
	right Operand
}

func (a AddInt) Evaluate(environment Environment) (Operand, Environment) {
	leftOperand, _ := a.left.Evaluate(environment)
	rightOperand, _ := a.right.Evaluate(environment)
	return IntNumber(leftOperand.Int(environment) + rightOperand.Int(environment)), environment
}
func (a AddInt) Int(environment Environment) int {
	evaluated, _ := a.Evaluate(environment)
	return evaluated.Int(environment)
}
func (a AddInt) Bool(environment Environment) bool {
	evaluated, _ := a.Evaluate(environment)
	return evaluated.Bool(environment)
}

type MultiplyInt struct {
	left  Operand
	right Operand
}

func (m MultiplyInt) Evaluate(environment Environment) (Operand, Environment) {
	leftOperand, _ := m.left.Evaluate(environment)
	rightOperand, _ := m.right.Evaluate(environment)
	return IntNumber(leftOperand.Int(environment) * rightOperand.Int(environment)), environment
}
func (m MultiplyInt) Int(environment Environment) int {
	evaluated, _ := m.Evaluate(environment)
	return evaluated.Int(environment)
}
func (m MultiplyInt) Bool(environment Environment) bool {
	evaluated, _ := m.Evaluate(environment)
	return evaluated.Bool(environment)
}

type LessThanInt struct {
	left  Operand
	right Operand
}

func (l LessThanInt) Evaluate(environment Environment) (Operand, Environment) {
	leftOperand, _ := l.left.Evaluate(environment)
	rightOperand, _ := l.right.Evaluate(environment)
	return BoolOperand(leftOperand.Int(environment) < rightOperand.Int(environment)), environment
}
func (l LessThanInt) Int(environment Environment) int {
	evaluated, _ := l.Evaluate(environment)
	return evaluated.Int(environment)
}
func (l LessThanInt) Bool(environment Environment) bool {
	evaluated, _ := l.Evaluate(environment)
	return evaluated.Bool(environment)
}
