package PDA

import (
	"fmt"
	"testing"
)

func TestDPDA_Task04(t *testing.T) {
	rulebook := PDARulebook{rules: []PDARule{{
		state:          1,
		character:      '(',
		nextState:      2,
		popCharacter:   '$',
		pushCharacters: []int32{'b', '$'},
	}, {
		state:          2,
		character:      '(',
		nextState:      2,
		popCharacter:   'b',
		pushCharacters: []int32{'b', 'b'},
	}, {
		state:          2,
		character:      ')',
		nextState:      2,
		popCharacter:   'b',
		pushCharacters: []int32{},
	}, {
		state:          2,
		character:      -1,
		nextState:      1,
		popCharacter:   '$',
		pushCharacters: []int32{'$'},
	}}}

	dpdaDesign := DPDADesign{
		startState:      1,
		bottomCharacter: '$',
		acceptStates:    []int32{1},
		rulebook:        rulebook,
	}
	fmt.Println(dpdaDesign.Accepts("(((((((((())))))))))")) //true
	fmt.Println(dpdaDesign.Accepts("()(())((()))(()(()))")) //true
	fmt.Println(dpdaDesign.Accepts("(()(()(()()(()()))()")) //false
	fmt.Println(dpdaDesign.Accepts("()))"))                 //false
	fmt.Println(dpdaDesign.Accepts("((()))()"))             //true
}
