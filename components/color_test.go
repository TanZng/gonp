package components

import (
	"image/color"
	"testing"
)

func TestNewColorAndMask(t *testing.T) {
	tests := []struct {
		name      string
		input     color.RGBA
		wantValue color.RGBA
		wantMask  uint64
	}{
		{
			name:      "Black",
			input:     color.RGBA{0, 0, 0, 255},
			wantValue: color.RGBA{0, 0, 0, 255},
			wantMask:  MaskColor,
		},
		{
			name:      "Red",
			input:     color.RGBA{255, 0, 0, 255},
			wantValue: color.RGBA{255, 0, 0, 255},
			wantMask:  MaskColor,
		},
		{
			name:      "Transparent",
			input:     color.RGBA{0, 0, 0, 0},
			wantValue: color.RGBA{0, 0, 0, 0},
			wantMask:  MaskColor,
		},
		{
			name:      "RandomColor",
			input:     color.RGBA{10, 20, 30, 40},
			wantValue: color.RGBA{10, 20, 30, 40},
			wantMask:  MaskColor,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			col := NewColor(tt.input)

			if col == nil {
				t.Fatalf("NewColor(%v) returned nil", tt.input)
			}

			if col.Value != tt.wantValue {
				t.Errorf("Color.Value = %v, want %v", col.Value, tt.wantValue)
			}

			gotMask := col.Mask()
			if gotMask != tt.wantMask {
				t.Errorf("Mask() = %v, want %v", gotMask, tt.wantMask)
			}
		})
	}
}
