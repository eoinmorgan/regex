package dfa

type State struct {
	accepting bool
	transitions map[byte]*State
}

func (state State) accepts (chars []byte) bool {
	if len(chars) == 0 {
		return state.accepting
	}

	if nextState, found := state.transitions[chars[0]]; found{
		return nextState.accepts(chars[1:])
	} else {
		return false
	}
}