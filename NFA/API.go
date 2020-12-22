package NFA

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
func (d DFARulebook) NextStates(states []int32, character int32) (followRules map[DFARule]struct{}) {
	followRules = make(map[DFARule]struct{})
	for _, v := range states {
		d.followRulesFor(v, character, followRules)
	}
	return
}
func (d DFARulebook) followRulesFor(state, character int32, followRules map[DFARule]struct{}) {
	d.rulesFor(state, character, followRules)
	for rule := range followRules {
		rule.Follow()
	}
}
func (d DFARulebook) rulesFor(state, character int32, followRules map[DFARule]struct{}) {
	for _, v := range d.rules {
		if v.AppliesTo(state, character) {
			followRules[v] = struct{}{}
		}
	}
}
