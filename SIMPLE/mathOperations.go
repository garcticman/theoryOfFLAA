package SIMPLE

import (
	"fmt"
	"strconv"
)

type IntOperand int

func (operand IntOperand) Execute(Environment) Operand {
	return operand
}
func (operand IntOperand) Reducible() bool {
	return false
}
func (operand IntOperand) Reduce(Environment) (Operand, Environment) {
	return operand, nil
}
func (operand IntOperand) Int(Environment) int {
	return int(operand)
}
func (operand IntOperand) Bool(environment Environment) bool {
	return operand.Int(environment) > 0
}
func (operand IntOperand) String() string {
	return strconv.Itoa(int(operand))
}

type AddInt struct {
	left  Operand
	right Operand
}

func (add AddInt) Execute(environment Environment) Operand {
	return IntOperand(add.right.Int(environment) + add.left.Int(environment))
}
func (add AddInt) Reducible() bool {
	return true
}
func (add AddInt) Reduce(environment Environment) (operand Operand, resultEnvironment Environment) {
	if add.left.Reducible() {
		var leftReduced Operand
		leftReduced, resultEnvironment = add.left.Reduce(environment)

		operand = AddInt{
			left:  leftReduced,
			right: add.right,
		}
	} else if add.right.Reducible() {
		var rightReduced Operand
		rightReduced, resultEnvironment = add.right.Reduce(environment)

		operand = AddInt{
			left:  add.left,
			right: rightReduced,
		}
	} else {
		operand = add.Execute(environment)
	}
	return operand, environment
}
func (add AddInt) Int(environment Environment) int {
	return add.Execute(environment).Int(environment)
}
func (add AddInt) Bool(environment Environment) bool {
	return add.Int(environment) > 0
}
func (add AddInt) String() string {
	return fmt.Sprintf("%s + %s", add.left, add.right)
}

type MultiplyInt struct {
	left  Operand
	right Operand
}

func (multiply MultiplyInt) Execute(environment Environment) Operand {
	return IntOperand(multiply.right.Int(environment) * multiply.left.Int(environment))
}
func (multiply MultiplyInt) Reducible() bool {
	return true
}
func (multiply MultiplyInt) Reduce(environment Environment) (operand Operand, resultEnvironment Environment) {
	if multiply.left.Reducible() {
		var leftReduced Operand
		leftReduced, resultEnvironment = multiply.left.Reduce(environment)

		operand = MultiplyInt{
			left:  leftReduced,
			right: multiply.right,
		}
	} else if multiply.right.Reducible() {
		var rightReduced Operand
		rightReduced, resultEnvironment = multiply.right.Reduce(environment)

		operand = MultiplyInt{
			left:  multiply.left,
			right: rightReduced,
		}
	} else {
		operand = multiply.Execute(environment)
	}
	return
}
func (multiply MultiplyInt) Int(environment Environment) int {
	return multiply.Execute(environment).Int(environment)
}
func (multiply MultiplyInt) Bool(environment Environment) bool {
	return multiply.Int(environment) > 0
}
func (multiply MultiplyInt) String() string {
	return fmt.Sprintf("%s * %s", multiply.left, multiply.right)
}
