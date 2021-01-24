package NFA

import (
	"fmt"
	"testing"
)

func TestDFA(t *testing.T) {
	rulebook := DFARulebook{Rules: []DFARule{{
		State:     2,
		Character: 'b',
		NextState: 3,
	}, {
		State:     2,
		Character: 'a',
		NextState: 3,
	}, {
		State:     1,
		Character: 'b',
		NextState: 1,
	}, {
		State:     1,
		Character: 'a',
		NextState: 1,
	}, {
		State:     1,
		Character: 'b',
		NextState: 2,
	}, {
		State:     3,
		Character: 'a',
		NextState: 4,
	}, {
		State:     3,
		Character: 'b',
		NextState: 4,
	}}}

	nfa := New(1, []int32{4}, rulebook)
	nfa.ReadString("bbabb")
	fmt.Println(nfa.Accepting())
}

func PrintStructValueMap(valueKeyMap map[DFARule]struct{}) {
	for value := range valueKeyMap {
		fmt.Print(value.NextState, " ")
	}
	fmt.Println()
}

func TestNFA_OE_AE_IE_Endings(t *testing.T) {
	rulebook := DFARulebook{Rules: []DFARule{{
		State:     1,
		Character: 'а',
		NextState: 2,
	}, {
		State:     1,
		Character: 'о',
		NextState: 2,
	}, {
		State:     1,
		Character: 'и',
		NextState: 2,
	}, {
		State:     2,
		Character: 'е',
		NextState: 3,
	},
	}}

	alphabet := "абвгдеёжзийклмнопрстуфхцчшщъэьэюя"
	for _, v := range alphabet {
		rulebook.AddRule(1, int32(v), 1)
	}

	//nfa := New(1, []int32{3}, Rulebook)
	//nfa.CheckAndPrintWords([]string{"большое", "театр", "самолет", "дикие", "танцы", "трение", "дунае"})
	nfa := NFADesign{
		CurrentStates: []int32{1},
		AcceptStates:  []int32{3},
		Rulebook:      rulebook,
	}
	fmt.Println(nfa.Accepts("большое"))
	fmt.Println(nfa.Accepts("театр"))
	fmt.Println(nfa.Accepts("самолет"))
	fmt.Println(nfa.Accepts("дикие"))
	fmt.Println(nfa.Accepts("танцы"))
	fmt.Println(nfa.Accepts("трение"))
	fmt.Println(nfa.Accepts("дунае"))
}
