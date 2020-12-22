package DFA

type DFA struct {
	currentState int32
	acceptStates []int32
	rulebook     DFARulebook
}

func New(startState int32, acceptStates []int32, rulebook DFARulebook) DFA {
	return DFA{
		currentState: startState,
		acceptStates: acceptStates,
		rulebook:     rulebook,
	}
}
func (d *DFA) AddAcceptedState(state int32) {
	d.acceptStates = append(d.acceptStates, state)
}
func (d *DFA) RemoveAcceptedState(state int32) {
	for i, v := range d.acceptStates {
		if v == state {
			d.acceptStates = append(d.acceptStates[:i], d.acceptStates[i+1:]...)
			return
		}
	}
}
func (d DFA) Accepting() bool {
	for _, v := range d.acceptStates {
		return v == d.currentState
	}
	return false
}
func (d *DFA) ReadCharacter(character int32) {
	d.currentState = d.rulebook.NextState(d.currentState, character)
}
func (d *DFA) ReadString(string string) {
	for _, v := range string {
		d.ReadCharacter(v)
	}
}
