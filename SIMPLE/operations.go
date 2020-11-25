package SIMPLE

import (
	"fmt"
	"strconv"
)

type Operand interface {
	Execute(Environment) Operand
	Reducible() bool
	Reduce(Environment) (Operand, Environment)
	Int(Environment) int
	Bool(Environment) bool
	String() string
}

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

type BoolOperand bool

func (operand BoolOperand) Execute(Environment) Operand {
	return operand
}
func (operand BoolOperand) Reducible() bool {
	return false
}
func (operand BoolOperand) Reduce(Environment) (Operand, Environment) {
	return operand, nil
}
func (operand BoolOperand) Int(Environment) (result int) {
	if operand.Bool(nil) {
		result = 1
	}
	return
}
func (operand BoolOperand) Bool(Environment) bool {
	return bool(operand)
}
func (operand BoolOperand) String() string {
	if operand.Bool(nil) {
		return "true"
	} else {
		return "false"
	}
}

type LessThanInt struct {
	left Operand
	right Operand
}

func (lessThan LessThanInt) Execute(environment Environment) Operand {
	return BoolOperand(lessThan.left.Int(environment) < lessThan.right.Int(environment))
}
func (lessThan LessThanInt) Reducible() bool {
	return true
}
func (lessThan LessThanInt) Reduce(environment Environment) (operand Operand, resultEnvironment Environment) {
	if lessThan.left.Reducible() {
		var leftReduced Operand
		leftReduced, resultEnvironment = lessThan.left.Reduce(environment)

		operand = LessThanInt{
			left:  leftReduced,
			right: lessThan.right,
		}
	} else if lessThan.right.Reducible() {
		var rightReduced Operand
		rightReduced, resultEnvironment = lessThan.right.Reduce(environment)

		operand = LessThanInt{
			left:  lessThan.left,
			right: rightReduced,
		}
	} else {
		operand = lessThan.Execute(environment)
	}
	return
}
func (lessThan LessThanInt) Int(environment Environment) int {
	return lessThan.Execute(environment).Int(environment)
}
func (lessThan LessThanInt) Bool(environment Environment) bool {
	return lessThan.Execute(environment).Bool(environment)
}
func (lessThan LessThanInt) String() string {
	return fmt.Sprintf("%s < %s", lessThan.left, lessThan.right)
}

type GreaterThanInt struct {
	left Operand
	right Operand
}

func (greaterThan GreaterThanInt) Execute(environment Environment) Operand {
	return BoolOperand(greaterThan.left.Int(environment) > greaterThan.right.Int(environment))
}
func (greaterThan GreaterThanInt) Reducible() bool {
	return true
}
func (greaterThan GreaterThanInt) Reduce(environment Environment) (operand Operand, resultEnvironment Environment) {
	if greaterThan.left.Reducible() {
		var leftReduced Operand
		leftReduced, resultEnvironment = greaterThan.left.Reduce(environment)

		operand = GreaterThanInt{
			left:  leftReduced,
			right: greaterThan.right,
		}
	} else if greaterThan.right.Reducible() {
		var rightReduced Operand
		rightReduced, resultEnvironment = greaterThan.right.Reduce(environment)

		operand = GreaterThanInt{
			left:  greaterThan.left,
			right: rightReduced,
		}
	} else {
		operand = greaterThan.Execute(environment)
	}
	return
}
func (greaterThan GreaterThanInt) Int(environment Environment) int {
	return greaterThan.Execute(environment).Int(environment)
}
func (greaterThan GreaterThanInt) Bool(environment Environment) bool {
	return greaterThan.Execute(environment).Bool(environment)
}
func (greaterThan GreaterThanInt) String() string {
	return fmt.Sprintf("%s > %s", greaterThan.left, greaterThan.right)
}

type EqualInt struct {
	left Operand
	right Operand
}

func (equal EqualInt) Execute(Environment) Operand {
	return BoolOperand(equal.left.Int(nil) == equal.right.Int(nil))
}
func (equal EqualInt) Reducible() bool {
	return true
}
func (equal EqualInt) Reduce(environment Environment) (operand Operand, resultEnvironment Environment) {
	if equal.left.Reducible() {
		var leftReduced Operand
		leftReduced, resultEnvironment = equal.left.Reduce(environment)

		operand = EqualInt{
			left:  leftReduced,
			right: equal.right,
		}
	} else if equal.right.Reducible() {
		var rightReduced Operand
		rightReduced, resultEnvironment = equal.right.Reduce(environment)

		operand = EqualInt{
			left:  equal.left,
			right: rightReduced,
		}
	} else {
		operand = equal.Execute(environment)
	}
	return
}
func (equal EqualInt) Int(environment Environment) int {
	return equal.Execute(environment).Int(environment)
}
func (equal EqualInt) Bool(environment Environment) bool {
	return equal.Execute(environment).Bool(environment)
}
func (equal EqualInt) String() string {
	return fmt.Sprintf("%s == %s", equal.left, equal.right)
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
		operand := AssignEnd{
			Operand:     assign.Operand,
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

type AssignEnd struct {
	Operand
}

func (end AssignEnd) Execute(environment Environment) Operand {
	return end.Operand.Execute(environment)
}
func (end AssignEnd) Reducible() bool {
	return false
}
func (end AssignEnd) Reduce(environment Environment) (Operand, Environment) {
	return end, environment
}
func (end AssignEnd) Int(environment Environment) int {
	return end.Operand.Int(environment)
}
func (end AssignEnd) Bool(environment Environment) bool {
	return end.Operand.Bool(environment)
}
func (end AssignEnd) String() string {
	return end.Operand.String()
}