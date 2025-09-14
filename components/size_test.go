package components

import (
	"testing"
)

func TestNewSizeAndMask(t *testing.T) {
	tests := []struct {
		name       string
		width      float32
		height     float32
		wantWidth  float32
		wantHeight float32
		wantMask   uint64
	}{
		{
			name:       "ZeroSize",
			width:      0,
			height:     0,
			wantWidth:  0,
			wantHeight: 0,
			wantMask:   MaskSize,
		},
		{
			name:       "SquareSize",
			width:      10,
			height:     10,
			wantWidth:  10,
			wantHeight: 10,
			wantMask:   MaskSize,
		},
		{
			name:       "RectangleSize",
			width:      20.5,
			height:     15.3,
			wantWidth:  20.5,
			wantHeight: 15.3,
			wantMask:   MaskSize,
		},
		{
			name:       "FloatValues",
			width:      1.1,
			height:     2.2,
			wantWidth:  1.1,
			wantHeight: 2.2,
			wantMask:   MaskSize,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			size := NewSize(tt.width, tt.height)

			if size.Width != tt.wantWidth {
				t.Errorf("Width = %v, want %v", size.Width, tt.wantWidth)
			}
			if size.Height != tt.wantHeight {
				t.Errorf("Height = %v, want %v", size.Height, tt.wantHeight)
			}

			gotMask := size.Mask()
			if gotMask != tt.wantMask {
				t.Errorf("Mask() = %v, want %v", gotMask, tt.wantMask)
			}
		})
	}
}
