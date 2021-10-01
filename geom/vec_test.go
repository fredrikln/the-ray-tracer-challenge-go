package geom

import (
	"math"
	"testing"
)

func TestVec(t *testing.T) {
	t.Run("Test", func(t *testing.T) {
		vector := Vec{1.0, 2.0, 3.0}

		if vector.X != 1.0 && vector.Y != 2.0 && vector.Z != 3.0 {
			t.Error("Vector coordinates not correct")
		}
	})
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a    Vec
		b    Vec
		want Vec
	}{
		{"Test", Vec{3.0, -2.0, 5.0}, Vec{-2.0, 3.0, 1.0}, Vec{1.0, 1.0, 6.0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.a.Add(tt.b); got != tt.want {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		name string
		a    Vec
		b    Vec
		want Vec
	}{
		{"Test 1", Vec{3.0, 2.0, 1.0}, Vec{5.0, 6.0, 7.0}, Vec{-2.0, -4.0, -6.0}},
		{"Test 2", Vec{0.0, 0.0, 0.0}, Vec{1.0, -2.0, 3.0}, Vec{-1.0, 2.0, -3.0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.a.Sub(tt.b); got != tt.want {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNeg(t *testing.T) {
	tests := []struct {
		name string
		a    Vec
		want Vec
	}{
		{"Test", Vec{1.0, -2.0, 3.0}, Vec{-1.0, 2.0, -3.0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.a.Neg(); got != tt.want {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		name string
		a    Vec
		b    float64
		want Vec
	}{
		{"Test 1", Vec{1.0, -2.0, 3.0}, 3.5, Vec{3.5, -7.0, 10.5}},
		{"Test 2", Vec{1.0, -2.0, 3.0}, 0.5, Vec{0.5, -1.0, 1.5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.a.Mul(tt.b); got != tt.want {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiv(t *testing.T) {
	tests := []struct {
		name string
		a    Vec
		b    float64
		want Vec
	}{
		{"Test", Vec{1.0, -2.0, 3.0}, 2.0, Vec{0.5, -1.0, 1.5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.a.Div(tt.b); got != tt.want {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMag(t *testing.T) {
	tests := []struct {
		name string
		a    Vec
		want float64
	}{
		{"Test 1", Vec{1.0, 0.0, 0.0}, 1.0},
		{"Test 2", Vec{0.0, 1.0, 0.0}, 1.0},
		{"Test 3", Vec{0.0, 0.0, 1.0}, 1.0},
		{"Test 4", Vec{1.0, 2.0, 3.0}, math.Sqrt(14.0)},
		{"Test 5", Vec{-1.0, -2.0, -3.0}, math.Sqrt(14.0)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.a.Mag(); got != tt.want {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNorm(t *testing.T) {
	tests := []struct {
		name string
		a    Vec
		want Vec
	}{
		{"Test 1", Vec{1.0, 0.0, 0.0}, Vec{1.0, 0.0, 0.0}},
		{"Test 2", Vec{1.0, 2.0, 3.0}, Vec{1 / math.Sqrt(14.0), 2 / math.Sqrt(14.0), 3 / math.Sqrt(14.0)}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.a.Norm(); got != tt.want {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("Test 3", func(t *testing.T) {
		vec := Vec{1.0, 2.0, 3.0}
		mag := vec.Norm().Mag()

		if mag != 1 {
			t.Errorf("Normalized vector does not have magnitude 1")
		}
	})
}

func TestDot(t *testing.T) {
	tests := []struct {
		name string
		a    Vec
		b    Vec
		want float64
	}{
		{"Test 1", Vec{1.0, 2.0, 3.0}, Vec{2.0, 3.0, 4.0}, 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.a.Dot(tt.b); got != tt.want {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCross(t *testing.T) {
	tests := []struct {
		name string
		a    Vec
		b    Vec
		want Vec
	}{
		{"Test 1", Vec{1.0, 2.0, 3.0}, Vec{2.0, 3.0, 4.0}, Vec{-1, 2, -1}},
		{"Test 2", Vec{2.0, 3.0, 4.0}, Vec{1.0, 2.0, 3.0}, Vec{1, -2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.a.Cross(tt.b); got != tt.want {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}
