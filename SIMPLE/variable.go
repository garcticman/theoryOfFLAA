package SIMPLE

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