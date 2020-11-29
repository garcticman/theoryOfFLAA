package SIMPLE

import "fmt"

type Variable struct {
	Name string
}

func (variable Variable) Execute(environment Environment) Operand {
	return environment[variable.Name].Execute(environment)
}
func (variable Variable) Reducible() bool {
	return true
}
func (variable Variable) Reduce(environment Environment) (Operand, Environment) {
	return environment[variable.Name], environment
}
func (variable Variable) Int(environment Environment) int {
	return environment[variable.Name].Int(environment)
}
func (variable Variable) Bool(environment Environment) bool {
	return environment[variable.Name].Bool(environment)
}
func (variable Variable) String() string {
	return variable.Name
}

type Assign struct {
	variableName string
	Operand
}

func (assign Assign) Execute(environment Environment) Operand {
	operand, _ := assign.Operand.Reduce(environment)
	return Assign{
		variableName: assign.variableName,
		Operand:      operand,
	}
}
func (assign Assign) Reducible() bool {
	return true
}
func (assign Assign) Reduce(environment Environment) (Operand, Environment) {
	if assign.Operand.Reducible() {
		operand := assign.Execute(environment)
		return operand, environment
	} else {
		newEnvironment := CopyEnvironment(environment)
		newEnvironment[assign.variableName] = assign.Operand
		operand := SentenceEnd{
			Operand: assign.Operand,
		}
		return operand, newEnvironment
	}
}
func (assign Assign) Int(environment Environment) int {
	return assign.Execute(environment).Int(environment)
}
func (assign Assign) Bool(environment Environment) bool {
	return assign.Execute(environment).Bool(environment)
}
func (assign Assign) String() string {
	return fmt.Sprintf("%s = %s", assign.variableName, assign.Operand)
}

type SentenceEnd struct {
	Operand
}

func (end SentenceEnd) Execute(environment Environment) Operand {
	return end.Operand.Execute(environment)
}
func (end SentenceEnd) Reducible() bool {
	return false
}
func (end SentenceEnd) Reduce(environment Environment) (Operand, Environment) {
	return end, environment
}
func (end SentenceEnd) Int(environment Environment) int {
	return end.Operand.Int(environment)
}
func (end SentenceEnd) Bool(environment Environment) bool {
	return end.Operand.Bool(environment)
}
func (end SentenceEnd) String() string {
	return "end"
}
