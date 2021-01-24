package NFA

type NFADesign struct {
	CurrentStates []int32
	AcceptStates  []int32
	Rulebook      DFARulebook
}

func (n NFADesign) Accepts(sequence string) bool {
	nfa := n.ToNFA(n.CurrentStates, n.AcceptStates, n.Rulebook)
	nfa.ReadString(sequence)
	return nfa.Accepting()
}
func (n NFADesign) ToNFA(currentStates []int32, acceptState []int32, rulebook DFARulebook) NFA {
	return NFA{
		currentStates: currentStates,
		acceptStates:  acceptState,
		rulebook:      rulebook,
	}
}
