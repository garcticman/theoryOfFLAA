package NFA

import (
	"fmt"
	"testing"
)

func TestDFA(t *testing.T) {
	rulebook := DFARulebook{rules: []DFARule{{
		state:     2,
		character: 'b',
		nextState: 3,
	}, {
		state:     2,
		character: 'a',
		nextState: 3,
	}, {
		state:     1,
		character: 'b',
		nextState: 1,
	}, {
		state:     1,
		character: 'a',
		nextState: 1,
	}, {
		state:     1,
		character: 'b',
		nextState: 2,
	}, {
		state:     3,
		character: 'a',
		nextState: 4,
	}, {
		state:     3,
		character: 'b',
		nextState: 4,
	}}}

	nfa := New(1, []int32{4}, rulebook)
	nfa.ReadString("bbabb")
	fmt.Println(nfa.Accepting())
}

func PrintStructValueMap(valueKeyMap map[DFARule]struct{}) {
	for value := range valueKeyMap {
		fmt.Print(value.nextState, " ")
	}
	fmt.Println()
}
