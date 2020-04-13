package nfa

import (
	"reflect"
	"testing"
)

func TestGetEpsilonClosureNoTransitions(t *testing.T) {
	state := State{}
	epsilonClosure := getEpsilonClosure(&state)
	expectedClosure := map[*State]bool {&state: true}

	if !reflect.DeepEqual(epsilonClosure, expectedClosure) {
		t.Errorf("Incorrect epsilon closure, expected %v, got %v", expectedClosure, epsilonClosure)
	}
}

func TestGetEpsilonClosureOneTransition(t *testing.T) {
	state1 := State{}
	state2 := State{}
	state1.epsilonTransitions = append(state1.epsilonTransitions, &state2)
	epsilonClosure := getEpsilonClosure(&state1)
	expectedClosure := map[*State]bool {&state1: true, &state2: true}

	if !reflect.DeepEqual(epsilonClosure, expectedClosure) {
		t.Errorf("Incorrect epsilon closure, expected %v, got %v", expectedClosure, epsilonClosure)
	}
}

func TestGetEpsilonClosureFanOut(t *testing.T) {
	state1 := State{}
	state2 := State{}
	state3 := State{}
	state1.epsilonTransitions = append(state1.epsilonTransitions, &state2)
	state1.epsilonTransitions = append(state1.epsilonTransitions, &state3)

	epsilonClosure := getEpsilonClosure(&state1)
	expectedClosure := map[*State]bool {&state1: true, &state2: true, &state3: true}

	if !reflect.DeepEqual(epsilonClosure, expectedClosure) {
		t.Errorf("Incorrect epsilon closure, expected %v, got %v", expectedClosure, epsilonClosure)
	}
}

func TestGetEpsilonClosureTwoLevels(t *testing.T) {
	state1 := State{}
	state2 := State{}
	state3 := State{}
	state1.epsilonTransitions = append(state1.epsilonTransitions, &state2)
	state2.epsilonTransitions = append(state2.epsilonTransitions, &state3)

	epsilonClosure := getEpsilonClosure(&state1)
	expectedClosure := map[*State]bool {&state1: true, &state2: true, &state3: true}

	if !reflect.DeepEqual(epsilonClosure, expectedClosure) {
		t.Errorf("Incorrect epsilon closure, expected %v, got %v", expectedClosure, epsilonClosure)
	}
}

func TestGetEpsilonClosureCycle(t *testing.T) {
	state1 := State{}
	state2 := State{}
	state1.epsilonTransitions = append(state1.epsilonTransitions, &state2)
	state2.epsilonTransitions = append(state2.epsilonTransitions, &state1)

	epsilonClosure := getEpsilonClosure(&state1)
	expectedClosure := map[*State]bool {&state1: true, &state2: true}

	if !reflect.DeepEqual(epsilonClosure, expectedClosure) {
		t.Errorf("Incorrect epsilon closure, expected %v, got %v", expectedClosure, epsilonClosure)
	}
}