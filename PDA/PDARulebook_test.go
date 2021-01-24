package PDA

import (
	"fmt"
	"testing"
)

func TestPDARulebook_NextConfig(t *testing.T) {
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
	configuration := PDAConfiguration{
		state: 1,
		stack: Stack{contents: []int32{'$'}},
	}

	configuration = rulebook.NextConfig(configuration, '(')
	PrintConfig(configuration)
	configuration = rulebook.NextConfig(configuration, '(')
	PrintConfig(configuration)
	configuration = rulebook.NextConfig(configuration, ')')
	PrintConfig(configuration)

	configuration = PDAConfiguration{2, Stack{contents: []int32{'$'}}}
	configuration = rulebook.followFreeMoves(configuration)
	PrintConfig(configuration)
}

func PrintConfig(configuration PDAConfiguration) {
	fmt.Print(configuration.state, " ")
	for _, v := range configuration.stack.contents {
		fmt.Print(string(v))
	}
	fmt.Println()
}
