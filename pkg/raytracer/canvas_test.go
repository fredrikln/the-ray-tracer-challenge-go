package raytracer

import (
	"strings"
	"testing"
)

func TestNewCanvas(t *testing.T) {
	t.Run("Test", func(t *testing.T) {
		canvas := NewCanvas(10, 20)

		if canvas.Width != 10 || canvas.Height != 20 {
			t.Error("Canvas width and height wrong")
		}

		for i, pixel := range canvas.Pixels {
			if pixel.R != 0.0 || pixel.G != 0.0 || pixel.B != 0.0 {
				t.Errorf("Pixel %d not initialized with black pixels", i)
			}
		}
	})
}

func TestSetPixel(t *testing.T) {
	t.Run("Test", func(t *testing.T) {
		canvas := NewCanvas(10, 20)
		red := NewColor(1.0, 0.0, 0.0)
		want := NewColor(1.0, 0.0, 0.0)

		canvas.SetPixel(2, 3, red)

		pixel := canvas.GetPixel(2, 3)

		// Get pixel from raw pixel array
		// row 3    = 3 x width (10) = 30
		// column 2 = third pixel on column
		// = 32
		pixel2 := canvas.Pixels[32]

		if !pixel.Eq(want) {
			t.Errorf("Got %v, want %v", pixel, want)
		}
		if !pixel2.Eq(want) {
			t.Errorf("Got %v, want %v", pixel, want)
		}
	})
}

func TestSavePPM(t *testing.T) {
	canvas := NewCanvas(5, 3)
	c1 := NewColor(1.5, 0.0, 0.0)
	c2 := NewColor(0.0, 0.5, 0.0)
	c3 := NewColor(-0.5, 0.0, 1.0)

	canvas.SetPixel(0, 0, c1)
	canvas.SetPixel(2, 1, c2)
	canvas.SetPixel(4, 2, c3)

	ppm := canvas.GetPPMString()
	lines := strings.Split(ppm, "\n")

	tests := []struct {
		name   string
		lineNo int
		want   string
	}{
		{"Line 0", 0, "P3"},
		{"Line 1", 1, "5 3"},
		{"Line 2", 2, "255"},
		{"Line 3", 3, "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0"},
		{"Line 4", 4, "0 0 0 0 0 0 0 128 0 0 0 0 0 0 0"},
		{"Line 5", 5, "0 0 0 0 0 0 0 0 0 0 0 0 0 0 255"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lines[tt.lineNo]

			if got != tt.want {
				t.Errorf("Line %d, got %s, want %s", tt.lineNo, got, tt.want)
			}
		})
	}
}

func TestSavePPM2(t *testing.T) {
	canvas := NewCanvas(10, 2)
	c := NewColor(1.0, 0.8, 0.6)

	for y := 0; y < canvas.Height; y += 1 {
		for x := 0; x < canvas.Width; x += 1 {
			canvas.SetPixel(x, y, c)
		}
	}

	ppm := canvas.GetPPMString()
	lines := strings.Split(ppm, "\n")

	tests := []struct {
		name   string
		lineNo int
		want   string
	}{
		{"Line 3", 3, "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204"},
		{"Line 4", 4, "153 255 204 153 255 204 153 255 204 153 255 204 153"},
		{"Line 5", 5, "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204"},
		{"Line 6", 6, "153 255 204 153 255 204 153 255 204 153 255 204 153"},
		{"Line 7", 7, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lines[tt.lineNo]

			if got != tt.want {
				t.Errorf("Line %d, got %s, want %s", tt.lineNo, got, tt.want)
			}
		})
	}
}
