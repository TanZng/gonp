// states/machine.go
package states

type GameState uint64

const (
	StateWelcome GameState = 1 << iota
	StatePlaying
	StatePause
	StateGameOver
)

type StateMachine struct {
	current GameState
}

func NewStateMachine(initial GameState) *StateMachine {
	return &StateMachine{current: initial}
}

func (s *StateMachine) Is(state GameState) bool {
	return s.current&state != 0
}

func (s *StateMachine) Set(state GameState) {
	s.current = state
}

func (s *StateMachine) Current() GameState {
	return s.current
}
