package components

type Size struct {
	Height float32
	Width  float32
}

func (a *Size) Mask() uint64 {
	return MaskSize
}
