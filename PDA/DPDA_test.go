package PDA

import (
	"fmt"
	"testing"
)

func TestDPDA(t *testing.T) {
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

	dpda := DPDA{
		currentConfig: PDAConfiguration{
			state: 1,
			stack: Stack{contents: []int32{'$'}},
		},
		acceptStates: []int32{1},
		rulebook:     rulebook,
	}

	dpda.ReadString("(()(")
	fmt.Println(dpda.Accepting())
	PrintConfig(dpda.currentConfig)

	dpda.ReadString("))()")
	fmt.Println(dpda.Accepting())
	PrintConfig(dpda.currentConfig)
}
