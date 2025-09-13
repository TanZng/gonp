package components

import "time"

type StateControls uint8

const (
	No StateControls = iota
	Up
	Down
)

type State struct {
	Duration time.Duration `json:"duration"`
	Next     uint8         `json:"next"`
	Start    time.Time     `json:"start"`
	Value    uint8         `json:"value"`
}

func (a *State) Mask() uint64 {
	return MaskState
}
