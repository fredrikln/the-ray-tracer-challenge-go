package material_test

import (
	"testing"

	m "github.com/fredrikln/the-ray-tracer-challenge-go/material"
)

func TestColor(t *testing.T) {
	t.Run("Test", func(t *testing.T) {
		color := m.NewColor(-0.5, 0.4, 1.7)

		if color.R != -0.5 || color.G != 0.4 || color.B != 1.7 {
			t.Error("Color struct invalid")
		}
	})
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a    m.Color
		b    m.Color
		want m.Color
	}{
		{"Test", m.NewColor(0.9, 0.6, 0.75), m.NewColor(0.7, 0.1, 0.25), m.NewColor(1.6, 0.7, 1.0)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Add(tt.b); !got.Eq(tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		name string
		a    m.Color
		b    m.Color
		want m.Color
	}{
		{"Test", m.NewColor(0.9, 0.6, 0.75), m.NewColor(0.7, 0.1, 0.25), m.NewColor(0.2, 0.5, 0.5)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Sub(tt.b); !got.Eq(tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMulFloat(t *testing.T) {
	tests := []struct {
		name string
		a    m.Color
		b    float64
		want m.Color
	}{
		{"Test", m.NewColor(0.2, 0.3, 0.4), 2, m.NewColor(0.4, 0.6, 0.8)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.MulFloat(tt.b); !got.Eq(tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		name string
		a    m.Color
		b    m.Color
		want m.Color
	}{
		{"Test", m.NewColor(1.0, 0.2, 0.4), m.NewColor(0.9, 1.0, 0.1), m.NewColor(0.9, 0.2, 0.04)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Mul(tt.b); !got.Eq(tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}
