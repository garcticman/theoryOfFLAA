package LongSteps

type DoNothing struct {
	Operand
}

func (d DoNothing) Evaluate(environment Environment) (Operand, Environment) {
	return d.Operand, environment
}
func (d DoNothing) Int(environment Environment) int {
	return d.Operand.Int(environment)
}
func (d DoNothing) Bool(environment Environment) bool {
	return d.Operand.Bool(environment)
}

type If struct {
	condition   Operand
	consequence Operand
	alternative Operand
}

func (i If) Evaluate(environment Environment) (Operand, Environment) {
	if i.condition.Bool(environment) {
		return i.consequence.Evaluate(environment)
	} else {
		return i.alternative.Evaluate(environment)
	}
}
func (i If) Int(environment Environment) int {
	evaluated, _ := i.Evaluate(environment)
	return evaluated.Int(environment)
}
func (i If) Bool(environment Environment) bool {
	evaluated, _ := i.Evaluate(environment)
	return evaluated.Bool(environment)
}

type Sequence struct {
	left  Operand
	right Operand
}

func (s Sequence) Evaluate(environment Environment) (Operand, Environment) {
	_, evaluatedEnvironment := s.left.Evaluate(environment)
	return s.right.Evaluate(evaluatedEnvironment)
}
func (s Sequence) Int(environment Environment) int {
	evaluated, _ := s.Evaluate(environment)
	return evaluated.Int(environment)
}
func (s Sequence) Bool(environment Environment) bool {
	evaluated, _ := s.Evaluate(environment)
	return evaluated.Bool(environment)
}

type While struct {
	condition Operand
	body      Operand
}

func (w While) Evaluate(environment Environment) (Operand, Environment) {
	if w.condition.Bool(environment) {
		_, newEnvironment := w.body.Evaluate(environment)
		return w.Evaluate(newEnvironment)
	} else {
		return w.body, environment
	}
}
func (w While) Int(environment Environment) int {
	evaluated, _ := w.Evaluate(environment)
	return evaluated.Int(environment)
}
func (w While) Bool(environment Environment) bool {
	evaluated, _ := w.Evaluate(environment)
	return evaluated.Bool(environment)
}
