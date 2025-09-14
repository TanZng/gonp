package components

import (
	"testing"
)

func TestNewPositionAndMethods(t *testing.T) {
	tests := []struct {
		name     string
		x        float32
		y        float32
		wantX    float32
		wantY    float32
		wantMask uint64
	}{
		{
			name:     "DefaultPosition",
			x:        0,
			y:        0,
			wantX:    0,
			wantY:    0,
			wantMask: MaskPosition,
		},
		{
			name:     "SetXOnly",
			x:        10.5,
			y:        0,
			wantX:    10.5,
			wantY:    0,
			wantMask: MaskPosition,
		},
		{
			name:     "SetYOnly",
			x:        0,
			y:        -3.2,
			wantX:    0,
			wantY:    -3.2,
			wantMask: MaskPosition,
		},
		{
			name:     "SetBoth",
			x:        7.7,
			y:        8.8,
			wantX:    7.7,
			wantY:    8.8,
			wantMask: MaskPosition,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pos := NewPosition().WithX(tt.x).WithY(tt.y)

			if pos.X != tt.wantX {
				t.Errorf("X = %v, want %v", pos.X, tt.wantX)
			}
			if pos.Y != tt.wantY {
				t.Errorf("Y = %v, want %v", pos.Y, tt.wantY)
			}

			gotMask := pos.Mask()
			if gotMask != tt.wantMask {
				t.Errorf("Mask() = %v, want %v", gotMask, tt.wantMask)
			}
		})
	}
}
