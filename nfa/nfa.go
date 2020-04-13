package nfa

import (
	"ebm.bz/regex/dfa"
)

// convenience type for a set of State
type stateSet map[*State]bool

func (set stateSet) contains(state *State) bool {
	_, found := set[state]
	return found
}

func (set stateSet) add(state *State) {
	set[state] = true
}

func (set stateSet) remove(state *State) {
	delete(set, state)
}

func newStateSet(states ...*State) stateSet{
	set := make(stateSet, len(states))
	for _, state := range states {
		set.add(state)
	}
	return set
}

type State struct {
	accepting bool
	transitions map[byte]*State
	epsilonTransitions []*State
}

func ToDFA(nfa []State, start *State) dfa.State {
	// Given an NFA as a collection of States (and a ref to the start state), construct an equivalent DFA and return
	// the start State.
	//startStates := getEpsilonClosure(start)



	return dfa.State{}
}

func getEpsilonClosure(state *State) stateSet {
	// set of *State
	closure := newStateSet(state)

	// iterative graph traversal of epsilon transitions using a stack
	traversal := make([]*State, len(state.epsilonTransitions))
	copy(traversal, state.epsilonTransitions)

	for len(traversal) > 0 {
		next := traversal[0]
		traversal = traversal[1:]
		if _, found := closure[next]; !found {
			closure[next] = true
			for _, next2 := range next.epsilonTransitions {
				traversal = append(traversal, next2)
			}
		}
	}

	return closure
}

//func getAlphabet(fa []State) (transitionChars []byte) {
//	for _, state := range finiteAutomata {
//		for k := range state.nextStates {
//			transitionChars = append(transitionChars, k)
//		}
//	}
//	return
//}