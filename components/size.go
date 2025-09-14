package components

import (
	"image/color"
)

type Size struct {
	Width  float32
	Height float32
	Color  color.RGBA
}

func (a *Size) Mask() uint64 {
	return MaskSize
}

func NewSize(w, h float32, col color.RGBA) *Size {
	return &Size{w, h, col}
}
