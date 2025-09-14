package components

import (
	"testing"
)

func TestNewScoreMaskAndIncrease(t *testing.T) {
	tests := []struct {
		name       string
		startX     float32
		startY     float32
		startValue int8
		increments int
		wantValue  int8
		wantX      float32
		wantY      float32
		wantMask   uint64
	}{
		{
			name:       "InitialZeroValue",
			startX:     1.0,
			startY:     2.0,
			startValue: 0,
			increments: 0,
			wantValue:  0,
			wantX:      1.0,
			wantY:      2.0,
			wantMask:   MaskScore,
		},
		{
			name:       "IncrementOnce",
			startX:     0,
			startY:     0,
			startValue: 0,
			increments: 1,
			wantValue:  1,
			wantX:      0,
			wantY:      0,
			wantMask:   MaskScore,
		},
		{
			name:       "IncrementMultiple",
			startX:     -1.5,
			startY:     3.4,
			startValue: 5,
			increments: 3,
			wantValue:  8,
			wantX:      -1.5,
			wantY:      3.4,
			wantMask:   MaskScore,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := NewScore(tt.startX, tt.startY)
			score.Value = tt.startValue

			for i := 0; i < tt.increments; i++ {
				score.Increase()
			}

			if score.Value != tt.wantValue {
				t.Errorf("Value = %v, want %v", score.Value, tt.wantValue)
			}
			if score.X != tt.wantX {
				t.Errorf("X = %v, want %v", score.X, tt.wantX)
			}
			if score.Y != tt.wantY {
				t.Errorf("Y = %v, want %v", score.Y, tt.wantY)
			}

			gotMask := score.Mask()
			if gotMask != tt.wantMask {
				t.Errorf("Mask() = %v, want %v", gotMask, tt.wantMask)
			}
		})
	}
}
