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

func TestNFA_OE_AE_IE_Endings(t *testing.T) {
	rulebook := DFARulebook{rules: []DFARule{{
		state:     1,
		character: 'а',
		nextState: 2,
	}, {
		state:     1,
		character: 'о',
		nextState: 2,
	}, {
		state:     1,
		character: 'и',
		nextState: 2,
	}, {
		state:     2,
		character: 'е',
		nextState: 3,
	},
	}}

	alphabet := "абвгдеёжзийклмнопрстуфхцчшщъэьэюя"
	for _, v := range alphabet {
		rulebook.AddRule(1, int32(v), 1)
	}

	nfa := New(1, []int32{3}, rulebook)
	nfa.CheckAndPrintWords([]string{"большое", "театр", "самолет", "дикие", "танцы", "трение", "дунае"})
}
