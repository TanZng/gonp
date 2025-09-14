package components

import (
	"image/color"
)

type Size struct {
	Height float32
	Width  float32
	Color  color.RGBA
}

func (a *Size) Mask() uint64 {
	return MaskSize
}
