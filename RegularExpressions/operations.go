package RegularExpressions

import (
	"dsm/NFA"
	"fmt"
)

type Operation interface {
	fmt.Stringer
	Precedence() int
	ToNFADesign(stateIterator *StateIterator) NFA.NFADesign
}

func Bracket(operation Operation, outerPrecedence int) string {
	if operation.Precedence() < outerPrecedence {
		return "(" + operation.String() + ")"
	} else {
		return operation.String()
	}
}

type Handler struct {
	StateIterator
}

func (h *Handler) Matches(op Operation, sequence string) bool {
	return op.ToNFADesign(&h.StateIterator).Accepts(sequence)
}

type Empty struct{}

func (e Empty) String() string {
	return ""
}
func (e Empty) Precedence() int {
	return 3
}
func (e Empty) ToNFADesign(stateIterator *StateIterator) NFA.NFADesign {
	state := stateIterator.GetNewState()
	return NFA.NFADesign{
		CurrentStates: []int32{state},
		AcceptStates:  []int32{state},
		Rulebook:      NFA.DFARulebook{},
	}
}

type Literal struct {
	Character int32
}

func (l Literal) String() string {
	return string(l.Character)
}
func (l Literal) Precedence() int {
	return 3
}
func (l Literal) ToNFADesign(stateIterator *StateIterator) NFA.NFADesign {
	firstState := stateIterator.GetNewState()
	acceptedState := stateIterator.GetNewState()

	rule := NFA.DFARule{
		State:     firstState,
		Character: l.Character,
		NextState: acceptedState,
	}
	return NFA.NFADesign{
		CurrentStates: []int32{firstState},
		AcceptStates:  []int32{acceptedState},
		Rulebook:      NFA.DFARulebook{Rules: []NFA.DFARule{rule}},
	}
}

type Concatenate struct {
	Left  Operation
	Right Operation
}

func (c Concatenate) String() string {
	return Bracket(c.Left, c.Precedence()) + Bracket(c.Right, c.Precedence())
}
func (c Concatenate) Precedence() int {
	return 1
}
func (c Concatenate) ToNFADesign(stateIterator *StateIterator) NFA.NFADesign {
	leftNFA := c.Left.ToNFADesign(stateIterator)
	rightNFA := c.Right.ToNFADesign(stateIterator)
	rulebook := NFA.DFARulebook{Rules: append(leftNFA.Rulebook.Rules, rightNFA.Rulebook.Rules...)}

	extraRules := make([]NFA.DFARule, len(leftNFA.AcceptStates)*len(rightNFA.CurrentStates))
	i := 0
	for _, acceptedLeft := range leftNFA.AcceptStates {
		for _, currentRight := range rightNFA.CurrentStates {
			extraRules[i] = NFA.DFARule{
				State:     acceptedLeft,
				Character: -1,
				NextState: currentRight,
			}
			i++
		}
	}
	rulebook.Rules = append(rulebook.Rules, extraRules...)

	return NFA.NFADesign{
		CurrentStates: leftNFA.CurrentStates,
		AcceptStates:  rightNFA.AcceptStates,
		Rulebook:      rulebook,
	}
}

type Choose struct {
	Left  Operation
	Right Operation
}

func (c Choose) String() string {
	return Bracket(c.Left, c.Precedence()) + "|" + Bracket(c.Right, c.Precedence())
}
func (c Choose) Precedence() int {
	return 0
}
func (c Choose) ToNFADesign(stateIterator *StateIterator) NFA.NFADesign {
	leftNFA := c.Left.ToNFADesign(stateIterator)
	rightNFA := c.Right.ToNFADesign(stateIterator)
	rulebook := NFA.DFARulebook{Rules: append(leftNFA.Rulebook.Rules, rightNFA.Rulebook.Rules...)}

	newStartState := stateIterator.GetNewState()
	extraRules := make([]NFA.DFARule, len(leftNFA.AcceptStates)+len(rightNFA.CurrentStates))
	i := 0
	for _, currentLeft := range leftNFA.CurrentStates {
		extraRules[i] = NFA.DFARule{
			State:     newStartState,
			Character: -1,
			NextState: currentLeft,
		}
		i++
	}
	for _, currentRight := range rightNFA.CurrentStates {
		extraRules[i] = NFA.DFARule{
			State:     newStartState,
			Character: -1,
			NextState: currentRight,
		}
		i++
	}
	rulebook.Rules = append(rulebook.Rules, extraRules...)

	return NFA.NFADesign{
		CurrentStates: []int32{newStartState},
		AcceptStates:  append(leftNFA.AcceptStates, rightNFA.AcceptStates...),
		Rulebook:      rulebook,
	}
}

type Repeat struct {
	Pattern Operation
}

func (r Repeat) String() string {
	return Bracket(r.Pattern, r.Precedence()) + "*"
}
func (r Repeat) Precedence() int {
	return 2
}
func (r Repeat) ToNFADesign(stateIterator *StateIterator) NFA.NFADesign {
	newStartEndState := stateIterator.GetNewState()
	patternNFA := r.Pattern.ToNFADesign(stateIterator)
	rules := patternNFA.Rulebook.Rules
	i := 0
	for _, accepted := range patternNFA.AcceptStates {
		for _, current := range patternNFA.CurrentStates {
			rules = append(rules, NFA.DFARule{
				State:     accepted,
				Character: -1,
				NextState: current,
			})
			i++
		}
	}
	for _, current := range patternNFA.CurrentStates {
		rules = append(rules, NFA.DFARule{
			State:     newStartEndState,
			Character: -1,
			NextState: current,
		})
		i++
	}
	rulebook := NFA.DFARulebook{Rules: rules}

	return NFA.NFADesign{
		CurrentStates: []int32{newStartEndState},
		AcceptStates:  append(patternNFA.AcceptStates, newStartEndState),
		Rulebook:      rulebook,
	}
}
