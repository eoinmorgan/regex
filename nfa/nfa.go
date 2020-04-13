package nfa

import (
	"ebm.bz/regex/dfa"
)

type State struct {
	accepting bool
	transitions map[byte]*State
	epsilonTransitions []*State
}

func ToDFA(nfa []State, start *State) dfa.State {
	// Given an NFA as a collection of States (and a ref to the start state), construct an equivalent DFA and return
	// the start State.




	return dfa.State{}
}

func getEpsilonClosure(state *State) map[*State]bool {
	closureSet := map[*State]bool{state: true}

	traversal := make([]*State, len(state.epsilonTransitions))
	copy(traversal, state.epsilonTransitions)

	for len(traversal) > 0 {
		next := traversal[0]
		traversal = traversal[1:]
		if _, found := closureSet[next]; !found {
			closureSet[next] = true
			for _, next2 := range next.epsilonTransitions {
				traversal = append(traversal, next2)
			}
		}
	}

	return closureSet
}

//func getAlphabet(fa []State) (transitionChars []byte) {
//	for _, state := range finiteAutomata {
//		for k := range state.nextStates {
//			transitionChars = append(transitionChars, k)
//		}
//	}
//	return
//}