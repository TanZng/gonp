package components

import (
	"testing"
)

func TestNewVelocityAndMethods(t *testing.T) {
	tests := []struct {
		name      string
		dx        float32
		dy        float32
		speed     float32
		wantDx    float32
		wantDy    float32
		wantSpeed float32
		wantMask  uint64
	}{
		{
			name:      "DefaultVelocity",
			dx:        0,
			dy:        0,
			speed:     0,
			wantDx:    0,
			wantDy:    0,
			wantSpeed: 0,
			wantMask:  MaskVelocity,
		},
		{
			name:      "SetDx",
			dx:        5,
			dy:        0,
			speed:     0,
			wantDx:    5,
			wantDy:    0,
			wantSpeed: 0,
			wantMask:  MaskVelocity,
		},
		{
			name:      "SetDy",
			dx:        0,
			dy:        -3.2,
			speed:     0,
			wantDx:    0,
			wantDy:    -3.2,
			wantSpeed: 0,
			wantMask:  MaskVelocity,
		},
		{
			name:      "SetSpeed",
			dx:        0,
			dy:        0,
			speed:     9.81,
			wantDx:    0,
			wantDy:    0,
			wantSpeed: 9.81,
			wantMask:  MaskVelocity,
		},
		{
			name:      "SetAll",
			dx:        1.5,
			dy:        2.5,
			speed:     3.5,
			wantDx:    1.5,
			wantDy:    2.5,
			wantSpeed: 3.5,
			wantMask:  MaskVelocity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vel := NewVelocity().
				WithDx(tt.dx).
				WithDy(tt.dy).
				WithSpeed(tt.speed)

			if vel.Dx != tt.wantDx {
				t.Errorf("Dx = %v, want %v", vel.Dx, tt.wantDx)
			}
			if vel.Dy != tt.wantDy {
				t.Errorf("Dy = %v, want %v", vel.Dy, tt.wantDy)
			}
			if vel.Speed != tt.wantSpeed {
				t.Errorf("Speed = %v, want %v", vel.Speed, tt.wantSpeed)
			}

			gotMask := vel.Mask()
			if gotMask != tt.wantMask {
				t.Errorf("Mask() = %v, want %v", gotMask, tt.wantMask)
			}
		})
	}
}
