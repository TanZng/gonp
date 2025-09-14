package components

type Score struct {
	Value int8
	X     float32
	Y     float32
}

func (s *Score) Mask() uint64 {
	return MaskScore
}

func (s *Score) Increase() {
	s.Value++
}

func NewScore(x, y float32) *Score {
	return &Score{X: x, Y: y}
}
