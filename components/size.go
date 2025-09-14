package components

type Size struct {
	Width  float32
	Height float32
}

func (a *Size) Mask() uint64 {
	return MaskSize
}

func NewSize(w, h float32) *Size {
	return &Size{w, h}
}
