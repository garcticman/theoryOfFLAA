package PDA

import (
	"fmt"
	"testing"
)

func TestPDARule_AppliesTo(t *testing.T) {
	rule := PDARule{
		state:          1,
		character:      '(',
		nextState:      2,
		popCharacter:   '$',
		pushCharacters: []int32{'b', '$'},
	}
	configuration := PDAConfiguration{
		state: 1,
		stack: Stack{contents: []int32{'$'}},
	}

	configuration = rule.Follow(configuration)
	fmt.Println(configuration.state)
	for _, v := range configuration.stack.contents {
		fmt.Println(string(v))
	}
}
