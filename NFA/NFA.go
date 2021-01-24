package NFA

import "fmt"

type NFA struct {
	currentStates []int32
	acceptStates  []int32
	rulebook      DFARulebook
}

func New(startState int32, acceptStates []int32, rulebook DFARulebook) NFA {
	return NFA{
		currentStates: []int32{startState},
		acceptStates:  acceptStates,
		rulebook:      rulebook,
	}
}
func (n *NFA) GetCurrentStates() []int32 {
	return n.rulebook.followFreeMoves(n.currentStates)
}
func (n *NFA) AddAcceptedState(state int32) {
	n.acceptStates = append(n.acceptStates, state)
}
func (n *NFA) RemoveAcceptedState(state int32) {
	for i, v := range n.acceptStates {
		if v == state {
			n.acceptStates = append(n.acceptStates[:i], n.acceptStates[i+1:]...)
			return
		}
	}
}

func (n *NFA) Accepting() bool {
	for _, v := range n.acceptStates {
		for _, v2 := range n.GetCurrentStates() {
			if v == v2 {
				return true
			}
		}
	}
	return false
}
func (n *NFA) ReadCharacter(character int32) {
	nextStates := n.rulebook.NextStates(n.GetCurrentStates(), character)
	n.currentStates = []int32{}
	for i := range nextStates {
		n.currentStates = append(n.currentStates, i.NextState)
	}
}
func (n *NFA) ReadString(string string) {
	for _, v := range string {
		n.ReadCharacter(v)
	}
}

func (n *NFA) CheckAndPrintWords(words []string) {
	for _, v := range words {
		newDfa := New(n.currentStates[0], n.acceptStates, n.rulebook)
		newDfa.ReadString(v)
		if newDfa.Accepting() {
			fmt.Println(v)
		}
	}
}
