package dfa

import "testing"

func TestStateAcceptsEmpty(t *testing.T) {
	emptyAcceptingState := State{accepting: true}
	if !emptyAcceptingState.accepts([]byte{}) {
		t.Error("Expected empty string to be accepted.")
	}
}

func TestStateRejectsEmpty(t *testing.T) {
	emptyAcceptingState := State{accepting: false}
	if emptyAcceptingState.accepts([]byte{}) {
		t.Error("Expected empty string to be rejected.")
	}
}

func TestStateAcceptsOne(t *testing.T) {
	acceptingState := State{accepting: true}
	startState := State{
		accepting: false,
		transitions: map[byte]*State{'a': &acceptingState},
	}

	if !startState.accepts([]byte{'a'}) {
		t.Error("Expected state to accept input 'a'")
	}
}

func TestStateRejectsMissing(t *testing.T) {
	acceptingState := State{accepting: true}
	startState := State{
		accepting: false,
		transitions: map[byte]*State{'a': &acceptingState},
	}

	if startState.accepts([]byte{'b'}) {
		t.Error("Expected state to reject input 'b'")
	}
}

func TestStateRejectsRemainingChars(t *testing.T) {
	acceptingState := State{accepting: true}
	startState := State{
		accepting: false,
		transitions: map[byte]*State{'a': &acceptingState},
	}

	if startState.accepts([]byte{'a', 'b'}) {
		t.Error("Expected state to reject trailing input 'b'")
	}
}

func TestStateComplex(t *testing.T) {
	// test the DFA representation of the regex "[a|b]c" as a sanity check

	acceptingState := State{accepting: true}
	bState := State{accepting: false, transitions: map[byte]*State{'c': &acceptingState}}
	aState := State{accepting: false, transitions: map[byte]*State{'c': &acceptingState}}
	startState := State{
		accepting: false,
		transitions: map[byte]*State{
			'a': &aState,
			'b': &bState,
		},
	}

	if !startState.accepts([]byte{'a', 'c'}) {
		t.Error("Expected state to accept input 'ac'")
	}

	if !startState.accepts([]byte{'b', 'c'}) {
		t.Error("Expected state to accept input 'bc'")
	}

	if startState.accepts([]byte{'a', 'b'}) {
		t.Error("Expected state to reject input 'ab'")
	}

	if startState.accepts([]byte{'a'}) {
		t.Error("Expected state to reject input 'a'")
	}

	if startState.accepts([]byte{'a', 'c', 'z'}) {
		t.Error("Expected state to reject input 'a'")
	}
}