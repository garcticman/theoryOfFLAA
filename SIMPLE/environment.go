package SIMPLE

type Environment map[string]Operand

func CopyEnvironment(environment Environment) Environment {
	result := make(Environment, len(environment))
	for key, value := range environment {
		result[key] = value
	}
	return result
}