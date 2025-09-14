package components

type Position struct {
	X float32
	Y float32
}

func (a *Position) Mask() uint64 {
	return MaskPosition
}

func (a *Position) WithX(x float32) *Position {
	a.X = x
	return a
}

func (a *Position) WithY(y float32) *Position {
	a.Y = y
	return a
}

func NewPosition() *Position {
	return &Position{}
}
