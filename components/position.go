package components

type Position struct {
	X float32
	Y float32
}

func (a *Position) Mask() uint64 {
	return MaskPosition
}
