package PDA

type DPDADesign struct {
	startState      int32
	bottomCharacter int32
	acceptStates    []int32
	rulebook        PDARulebook
}

func (dpdd DPDADesign) Accepts(sequence string) bool {
	dpda := dpdd.ToDPDA()
	dpda.ReadString(sequence)
	return dpda.Accepting()
}
func (dpdd DPDADesign) ToDPDA() DPDA {
	return DPDA{
		currentConfig: PDAConfiguration{
			state: dpdd.startState,
			stack: Stack{contents: []int32{dpdd.bottomCharacter}},
		},
		acceptStates: dpdd.acceptStates,
		rulebook:     dpdd.rulebook,
	}
}
