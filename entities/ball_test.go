package entities

import (
	"math"
	"testing"

	"github.com/TanZng/gonp/components"
)

// Utility for approximate float comparison
func almostEqual(a, b, tolerance float32) bool {
	return float32(math.Abs(float64(a-b))) <= tolerance
}

func TestBall_IncreaseSpeed(t *testing.T) {
	tests := []struct {
		name       string
		startSpeed float32
		wantSpeed  float32
	}{
		{
			name:       "Increase within limit",
			startSpeed: 0.2,
			wantSpeed:  0.21,
		},
		{
			name:       "At max speed",
			startSpeed: 0.3,
			wantSpeed:  0.3,
		},
		{
			name:       "Above max speed",
			startSpeed: 0.31,
			wantSpeed:  0.31,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ball := &Ball{
				Velocity: &components.Velocity{Speed: tt.startSpeed},
			}

			ball.IncreaseSpeed()

			if !almostEqual(ball.Speed, tt.wantSpeed, 0.0001) {
				t.Errorf("Speed = %v; want %v", ball.Speed, tt.wantSpeed)
			}
		})
	}
}

func TestBall_Reset(t *testing.T) {
	tests := []struct {
		name       string
		initialX   float32
		initialY   float32
		initialDx  float32
		initialDy  float32
		resetX     float32
		resetY     float32
		wantSpeed  float32
		wantDxFlip bool
	}{
		{
			name:       "Reset position and flip direction",
			initialX:   1,
			initialY:   2,
			initialDx:  1.0,
			initialDy:  0.5,
			resetX:     10,
			resetY:     20,
			wantSpeed:  0.1,
			wantDxFlip: true,
		},
		{
			name:       "Reset position and flip negative Dx",
			initialX:   4,
			initialY:   5,
			initialDx:  -0.5,
			initialDy:  0.0,
			resetX:     -3,
			resetY:     7,
			wantSpeed:  0.1,
			wantDxFlip: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ball := &Ball{
				Position: &components.Position{X: tt.initialX, Y: tt.initialY},
				Velocity: &components.Velocity{Dx: tt.initialDx, Dy: tt.initialDy},
				Size:     &components.Size{},
			}

			ball.Reset(tt.resetX, tt.resetY)

			if ball.X != tt.resetX || ball.Y != tt.resetY {
				t.Errorf("Position reset failed: got (%v, %v), want (%v, %v)", ball.X, ball.Y, tt.resetX, tt.resetY)
			}

			if !almostEqual(ball.Speed, tt.wantSpeed, 0.0001) {
				t.Errorf("Speed = %v, want %v", ball.Speed, tt.wantSpeed)
			}

			if tt.wantDxFlip && ball.Dx != -tt.initialDx {
				t.Errorf("Dx flip failed: got %v, want %v", ball.Dx, -tt.initialDx)
			}

			if ball.Dy == 0 {
				t.Error("Dy is 0 after reset; expected random non-zero value")
			}
		})
	}
}
