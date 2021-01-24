package NFA

type DFARule struct {
	State     int32
	Character int32
	NextState int32
}

func (d DFARule) AppliesTo(state, character int32) bool {
	return d.State == state && d.Character == character
}
func (d DFARule) Follow() int32 {
	return d.NextState
}

type DFARulebook struct {
	Rules []DFARule
}

func (d *DFARulebook) AddRule(startState, switchCharacter, nextState int32) {
	d.Rules = append(d.Rules, DFARule{
		State:     startState,
		Character: switchCharacter,
		NextState: nextState,
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
	for _, v := range d.Rules {
		if v.AppliesTo(state, character) {
			followRules[v] = struct{}{}
		}
	}
}
func (d *DFARulebook) followFreeMoves(nextStates []int32) []int32 {
	followedRules := d.NextStates(nextStates, -1)
	moreStates := make([]int32, len(followedRules))
	j := 0
	for i := range followedRules {
		moreStates[j] = i.NextState
		j++
	}
	if subset(moreStates, nextStates) {
		return nextStates
	} else {
		return d.followFreeMoves(append(nextStates, moreStates...))
	}
}

func subset(first, second []int32) bool {
	set := make(map[int32]int32)
	for _, value := range second {
		set[value] += 1
	}

	for _, value := range first {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}

	return true
}
