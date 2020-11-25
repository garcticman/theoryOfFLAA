package SIMPLE

import "fmt"

type Machine struct {
	Operand
	Environment
}
func (machine *Machine) Step() {
	machine.Operand, machine.Environment = machine.Reduce(machine.Environment)
}
func (machine *Machine) Run() {
	for machine.Reducible() {
		fmt.Println(machine)
		machine.Step()
	}
	fmt.Println(machine)
}