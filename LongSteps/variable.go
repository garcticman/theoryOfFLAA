package LongSteps

type Variable struct {
	Name string
}

func (v Variable) Evaluate(environment Environment) (Operand, Environment) {
	return environment[v.Name], nil
}
func (v Variable) Int(environment Environment) int {
	return environment[v.Name].Int(environment)
}
func (v Variable) Bool(environment Environment) bool {
	return environment[v.Name].Bool(environment)
}

type Assign struct {
	Name       string
	Expression Operand
}

func (a Assign) Evaluate(environment Environment) (Operand, Environment) {
	newEnvironment := CopyEnvironment(environment)
	newEnvironment[a.Name], _ = a.Expression.Evaluate(environment)
	return newEnvironment[a.Name], newEnvironment
}
func (a Assign) Int(environment Environment) int {
	return environment[a.Name].Int(environment)
}
func (a Assign) Bool(environment Environment) bool {
	return environment[a.Name].Bool(environment)
}
