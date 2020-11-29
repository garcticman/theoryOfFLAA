package SIMPLE

import (
	"fmt"
)

type Operand interface {
	Execute(Environment) Operand
	Reducible() bool
	Reduce(Environment) (Operand, Environment)
	Int(Environment) int
	Bool(Environment) bool
	String() string
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
	left  Operand
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
	left  Operand
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
	left  Operand
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

type If struct {
	condition   Operand
	consequence Operand
	alternative Operand
}

func (i If) Execute(environment Environment) Operand {
	operand, _ := i.condition.Reduce(environment)
	return If{
		condition:   operand,
		consequence: i.consequence,
		alternative: i.alternative,
	}
}
func (i If) Reducible() bool {
	return true
}
func (i If) Reduce(environment Environment) (resultOperand Operand, resultEnvironment Environment) {
	resultEnvironment = environment
	if i.condition.Reducible() {
		resultOperand = i.Execute(environment)
	} else {
		if i.condition.Bool(environment) {
			resultOperand = i.consequence
		} else {
			resultOperand = i.alternative
		}
	}
	return
}
func (i If) Int(environment Environment) int {
	return i.Execute(environment).Int(environment)
}
func (i If) Bool(environment Environment) bool {
	return i.Execute(environment).Bool(environment)
}
func (i If) String() string {
	return fmt.Sprintf("if (%s) {%s} else {%s}", i.condition, i.consequence, i.alternative)
}

type Sequence struct {
	first  Operand
	second Operand
}

func (s Sequence) Execute(environment Environment) Operand {
	operand, _ := s.Reduce(environment)
	return operand
}
func (s Sequence) Reducible() bool {
	return true
}
func (s Sequence) Reduce(environment Environment) (Operand, Environment) {
	switch s.first.(type) {
	case SentenceEnd:
		return s.second, environment
	default:
		reducedFirst, newEnvironment := s.first.Reduce(environment)
		return Sequence{
			first:  reducedFirst,
			second: s.second,
		}, newEnvironment
	}
}
func (s Sequence) Int(environment Environment) int {
	return s.Execute(environment).Int(environment)
}
func (s Sequence) Bool(environment Environment) bool {
	return s.Execute(environment).Bool(environment)
}
func (s Sequence) String() string {
	return fmt.Sprintf("%s; %s", s.first, s.second)
}

type While struct {
	condition Operand
	body      Operand
}

func (w While) Execute(environment Environment) Operand {
	panic("implement me")
}
func (w While) Reducible() bool {
	return true
}
func (w While) Reduce(environment Environment) (Operand, Environment) {
	return If{
		condition: w.condition,
		consequence: Sequence{
			first:  w.body,
			second: w,
		},
		alternative: SentenceEnd{BoolOperand(false)},
	}, environment
}
func (w While) Int(environment Environment) int {
	panic("implement me")
}
func (w While) Bool(environment Environment) bool {
	panic("implement me")
}
func (w While) String() string {
	return fmt.Sprintf("while (%s) {%s}", w.condition, w.body)
}
