package NFA

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
		for _, v2 := range n.currentStates {
			if v == v2 {
				return true
			}
		}
	}
	return false
}
func (n *NFA) ReadCharacter(character int32) {
	nextStates := n.rulebook.NextStates(n.currentStates, character)
	n.currentStates = []int32{}
	for i := range nextStates {
		n.currentStates = append(n.currentStates, i.nextState)
	}
}
func (n *NFA) ReadString(string string) {
	for _, v := range string {
		n.ReadCharacter(v)
	}
}
