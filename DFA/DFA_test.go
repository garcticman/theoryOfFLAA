package DFA

import (
	"fmt"
	"testing"
)

func TestDFA(t *testing.T) {
	rulebook := DFARulebook{rules: []DFARule{{
		state:     1,
		character: 'a',
		nextState: 2,
	}, {
		state:     2,
		character: 'a',
		nextState: 2,
	}, {
		state:     2,
		character: 'b',
		nextState: 3,
	}, {
		state:     1,
		character: 'b',
		nextState: 1,
	}, {
		state:     3,
		character: 'a',
		nextState: 3,
	}, {
		state:     3,
		character: 'b',
		nextState: 3,
	},
	}}

	dfa := New(1, []int32{3}, rulebook)
	dfa.ReadString("baba")
	fmt.Println(dfa.Accepting())

	dfa2 := New(1, []int32{3}, rulebook)
	dfa2.ReadString("baa")
	fmt.Println(dfa2.Accepting())
}
