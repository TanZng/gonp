package components

import "image/color"

type Color struct {
	Value color.RGBA
}

func (s *Color) Mask() uint64 {
	return MaskColor
}

func NewColor(color color.RGBA) *Color {
	return &Color{color}
}
