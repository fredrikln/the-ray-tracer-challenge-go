package raytracer

import (
	"math"
	"testing"
)

func TestNewVec(t *testing.T) {
	t.Run("Test", func(t *testing.T) {
		vector := NewVec(1.0, 2.0, 3.0)

		if vector.X != 1.0 || vector.Y != 2.0 || vector.Z != 3.0 {
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
		{"Test", NewVec(3.0, -2.0, 5.0), NewVec(-2.0, 3.0, 1.0), NewVec(1.0, 1.0, 6.0)},
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
		a    Vec
		b    Vec
		want Vec
	}{
		{"Test 1", NewVec(3.0, 2.0, 1.0), NewVec(5.0, 6.0, 7.0), NewVec(-2.0, -4.0, -6.0)},
		{"Test 2", NewVec(0.0, 0.0, 0.0), NewVec(1.0, -2.0, 3.0), NewVec(-1.0, 2.0, -3.0)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.a.Sub(tt.b); !got.Eq(tt.want) {
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
		{"Test", NewVec(1.0, -2.0, 3.0), NewVec(-1.0, 2.0, -3.0)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.a.Neg(); !got.Eq(tt.want) {
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
		{"Test 1", NewVec(1.0, -2.0, 3.0), 3.5, NewVec(3.5, -7.0, 10.5)},
		{"Test 2", NewVec(1.0, -2.0, 3.0), 0.5, NewVec(0.5, -1.0, 1.5)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.a.Mul(tt.b); !got.Eq(tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestMulMat(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		a    Vec
// 		b    *Matrix
// 		want Vec
// 	}{
// 		{"Test 1",
// 			NewVec(1, 2, 3),
// 			NewMatrix(
// 				1, 2, 3, 4,
// 				2, 4, 4, 2,
// 				8, 6, 4, 1,
// 				0, 0, 0, 1,
// 			),
// 			NewVec(18, 24, 33),
// 		},
// 		{"Test 2",
// 			NewVec(1, 2, 3),
// 			NewIdentityMatrix(),
// 			NewVec(1, 2, 3),
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tt.a.MulMat(tt.b); !got.Eq(tt.want) {
// 				t.Errorf("Got %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestDiv(t *testing.T) {
	tests := []struct {
		name string
		a    Vec
		b    float64
		want Vec
	}{
		{"Test", NewVec(1.0, -2.0, 3.0), 2.0, NewVec(0.5, -1.0, 1.5)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.a.Div(tt.b); !got.Eq(tt.want) {
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
		{"Test 1", NewVec(1.0, 0.0, 0.0), 1.0},
		{"Test 2", NewVec(0.0, 1.0, 0.0), 1.0},
		{"Test 3", NewVec(0.0, 0.0, 1.0), 1.0},
		{"Test 4", NewVec(1.0, 2.0, 3.0), math.Sqrt(14.0)},
		{"Test 5", NewVec(-1.0, -2.0, -3.0), math.Sqrt(14.0)},
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
		{"Test 1", NewVec(1.0, 0.0, 0.0), NewVec(1.0, 0.0, 0.0)},
		{"Test 2", NewVec(1.0, 2.0, 3.0), NewVec(1/math.Sqrt(14.0), 2/math.Sqrt(14.0), 3/math.Sqrt(14.0))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.a.Norm(); !got.Eq(tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("Test 3", func(t *testing.T) {
		vec := NewVec(1.0, 2.0, 3.0)
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
		{"Test 1", NewVec(1.0, 2.0, 3.0), NewVec(2.0, 3.0, 4.0), 20},
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
		{"Test 1", NewVec(1.0, 2.0, 3.0), NewVec(2.0, 3.0, 4.0), NewVec(-1, 2, -1)},
		{"Test 2", NewVec(2.0, 3.0, 4.0), NewVec(1.0, 2.0, 3.0), NewVec(1, -2, 1)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.a.Cross(tt.b); !got.Eq(tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReflect(t *testing.T) {
	tests := []struct {
		name string
		v    Vec
		n    Vec
		want Vec
	}{
		{
			"Test 1",
			NewVec(1, -1, 0),
			NewVec(0, 1, 0),
			NewVec(1, 1, 0),
		},
		{
			"Test 2",
			NewVec(0, -1, 0),
			NewVec(math.Sqrt(2)/2, math.Sqrt(2)/2, 0),
			NewVec(1, 0, 0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Reflect(tt.n); !got.Eq(tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}
