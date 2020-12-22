package DFA

type DFARule struct {
	state     int32
	character int32
	nextState int32
}

func (d DFARule) AppliesTo(state, character int32) bool {
	return d.state == state && d.character == character
}
func (d DFARule) Follow() int32 {
	return d.nextState
}

type DFARulebook struct {
	rules []DFARule
}

func (d *DFARulebook) AddRule(startState, switchCharacter, nextState int32) {
	d.rules = append(d.rules, DFARule{
		state:     startState,
		character: switchCharacter,
		nextState: nextState,
	})
}
func (d DFARulebook) NextState(state, character int32) int32 {
	return d.ruleFor(state, character).Follow()
}
func (d DFARulebook) ruleFor(state, character int32) DFARule {
	for _, v := range d.rules {
		if v.AppliesTo(state, character) {
			return v
		}
	}

	return DFARule{}
}
