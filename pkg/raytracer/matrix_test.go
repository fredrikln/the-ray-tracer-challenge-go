package raytracer

import (
	"math"
	"testing"

	c "github.com/fredrikln/the-ray-tracer-challenge-go/common"
)

func TestMatrix4(t *testing.T) {
	t.Run("Test", func(t *testing.T) {
		m := NewMatrix(
			1.0, 2.0, 3.0, 4.0,
			5.5, 6.5, 7.5, 8.5,
			9.0, 10.0, 11.0, 12.0,
			13.5, 14.5, 15.5, 16.5,
		)

		if m.data[0][0] != 1.0 || m.data[0][3] != 4.0 || m.data[1][0] != 5.5 || m.data[1][2] != 7.5 || m.data[2][2] != 11.0 || m.data[3][0] != 13.5 || m.data[3][2] != 15.5 {
			t.Error("Invalid matrix content")
		}
	})

}

func TestEq(t *testing.T) {
	tests := []struct {
		name string
		a    *Matrix
		b    *Matrix
		want bool
	}{
		{
			"Test 1",
			NewMatrix(
				1, 2, 3, 4,
				5, 6, 7, 8,
				9, 8, 7, 6,
				5, 4, 3, 2,
			),
			NewMatrix(
				1, 2, 3, 4,
				5, 6, 7, 8,
				9, 8, 7, 6,
				5, 4, 3, 2,
			),
			true,
		},
		{
			"Test 2",
			NewMatrix(
				1, 2, 3, 4,
				5, 6, 7, 8,
				9, 8, 7, 6,
				5, 4, 3, 2,
			),
			NewMatrix(
				2, 3, 4, 5,
				6, 7, 8, 9,
				8, 7, 6, 5,
				4, 3, 2, 1,
			),
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Eq(tt.b); got != tt.want {
				t.Errorf("Got %t, want %t", got, tt.want)
			}
		})
	}
}

func TestMatrixMul(t *testing.T) {
	tests := []struct {
		name string
		a    *Matrix
		b    *Matrix
		want *Matrix
	}{
		{
			"Test 1",
			NewMatrix(
				1, 2, 3, 4,
				5, 6, 7, 8,
				9, 8, 7, 6,
				5, 4, 3, 2,
			),
			NewMatrix(
				-2, 1, 2, 3,
				3, 2, 1, -1,
				4, 3, 6, 5,
				1, 2, 7, 8,
			),
			NewMatrix(
				20, 22, 50, 48,
				44, 54, 114, 108,
				40, 58, 110, 102,
				16, 26, 46, 42,
			),
		},
		{
			"Test 2",
			NewMatrix(
				1, 2, 3, 4,
				5, 6, 7, 8,
				9, 8, 7, 6,
				5, 4, 3, 2,
			),
			NewIdentityMatrix(),
			NewMatrix(
				1, 2, 3, 4,
				5, 6, 7, 8,
				9, 8, 7, 6,
				5, 4, 3, 2,
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Mul(tt.b); !got.Eq(tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrixMulPoint(t *testing.T) {
	tests := []struct {
		name string
		a    *Matrix
		b    Point
		want Point
	}{
		{
			"Test 1",
			NewMatrix(
				1, 2, 3, 4,
				2, 4, 4, 2,
				8, 6, 4, 1,
				0, 0, 0, 1,
			),
			NewPoint(1, 2, 3),
			NewPoint(18, 24, 33),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.MulPoint(tt.b); !got.Eq(tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTranspose(t *testing.T) {
	tests := []struct {
		name string
		a    *Matrix
		want *Matrix
	}{
		{
			"Test 1",
			NewMatrix(
				0, 9, 3, 0,
				9, 8, 0, 8,
				1, 8, 5, 3,
				0, 0, 5, 8,
			),
			NewMatrix(
				0, 9, 1, 0,
				9, 8, 8, 0,
				3, 0, 5, 5,
				0, 8, 3, 8,
			),
		},
		{
			"Test 2",
			NewIdentityMatrix(),
			NewIdentityMatrix(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Transpose(); !got.Eq(tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeterminant(t *testing.T) {
	tests := []struct {
		name string
		a    [2][2]float64
		want float64
	}{
		{
			"Test 1",
			[2][2]float64{
				{1, 5},
				{-3, 2},
			},
			17,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := determinant2(tt.a); got != tt.want {
				t.Errorf("Got %f, want %f", got, tt.want)
			}
		})
	}
}

func TestSubmatrix3(t *testing.T) {
	tests := []struct {
		name string
		a    [3][3]float64
		want [2][2]float64
	}{
		{
			"Test 1",
			[3][3]float64{
				{1, 5, 0},
				{-3, 2, 7},
				{0, 6, -3},
			},
			[2][2]float64{
				{-3, 2},
				{0, 6},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := submatrix3(tt.a, 0, 2); got != tt.want {
				t.Errorf("Got %f, want %f", got, tt.want)
			}
		})
	}
}

func TestSubmatrix4(t *testing.T) {
	tests := []struct {
		name string
		a    [4][4]float64
		want [3][3]float64
	}{
		{
			"Test 1",
			[4][4]float64{
				{-6, 1, 1, 6},
				{-8, 5, 8, 6},
				{-1, 0, 8, 2},
				{-7, 1, -1, 1},
			},
			[3][3]float64{
				{-6, 1, 6},
				{-8, 8, 6},
				{-7, -1, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := submatrix4(tt.a, 2, 1); got != tt.want {
				t.Errorf("Got %f, want %f", got, tt.want)
			}
		})
	}
}

func TestMinor3(t *testing.T) {
	tests := []struct {
		name string
		a    [3][3]float64
		row  int
		col  int
		want float64
	}{
		{
			"Test 1",
			[3][3]float64{
				{3, 5, 0},
				{2, -1, -7},
				{6, -1, 5},
			},
			1,
			0,
			25,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minor3(tt.a, tt.row, tt.col); got != tt.want {
				t.Errorf("Got %f, want %f", got, tt.want)
			}
			if got := determinant2(submatrix3(tt.a, tt.row, tt.col)); got != tt.want {
				t.Errorf("Got %f, want %f", got, tt.want)
			}
		})
	}
}

func TestCofactor3(t *testing.T) {
	tests := []struct {
		name     string
		a        [3][3]float64
		row      int
		col      int
		minor    float64
		cofactor float64
	}{
		{
			"Test 1",
			[3][3]float64{
				{3, 5, 0},
				{2, -1, -7},
				{6, -1, 5},
			},
			0,
			0,
			-12,
			-12,
		},
		{
			"Test 2",
			[3][3]float64{
				{3, 5, 0},
				{2, -1, -7},
				{6, -1, 5},
			},
			1,
			0,
			25,
			-25,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cofactor3(tt.a, tt.row, tt.col); got != tt.cofactor {
				t.Errorf("Got %f, want %f", got, tt.cofactor)
			}
			if got := minor3(tt.a, tt.row, tt.col); got != tt.minor {
				t.Errorf("Got %f, want %f", got, tt.minor)
			}
		})
	}
}

func TestDeterminant3(t *testing.T) {
	tests := []struct {
		name string
		a    [3][3]float64
		want float64
	}{
		{
			"Test 1",
			[3][3]float64{
				{1, 2, 6},
				{-5, 8, -4},
				{2, 6, 4},
			},
			-196,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := determinant3(tt.a); got != tt.want {
				t.Errorf("Got %f, want %f", got, tt.want)
			}
		})
	}
}

func TestDeterminant4(t *testing.T) {
	tests := []struct {
		name string
		a    [4][4]float64
		want float64
	}{
		{
			"Test 1",
			[4][4]float64{
				{-2, -8, 3, 5},
				{-3, 1, 7, 3},
				{1, 2, -9, 6},
				{-6, 7, 7, -9},
			},
			-4071,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := determinant4(tt.a); got != tt.want {
				t.Errorf("Got %f, want %f", got, tt.want)
			}
		})
	}
}

func TestInvertible(t *testing.T) {
	tests := []struct {
		name string
		a    *Matrix
		want bool
	}{
		{
			"Test 1",
			NewMatrix(
				6, 4, 4, 4,
				5, 5, 7, 8,
				4, -9, 3, -7,
				9, 1, 7, -6,
			),
			true,
		},
		{
			"Test 2",
			NewMatrix(
				-4, 2, -2, 3,
				9, 6, 2, 6,
				0, -5, 1, -5,
				0, 0, 0, 0,
			),
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Invertible(); got != tt.want {
				t.Errorf("Got %t, want %t", got, tt.want)
			}
		})
	}
}

func TestInverse(t *testing.T) {
	tests := []struct {
		name string
		a    *Matrix
		want *Matrix
	}{
		{
			"Test 1",
			NewMatrix(
				-5, 2, 6, -8,
				1, -5, 1, 8,
				7, 7, -6, -7,
				1, -3, 7, 4,
			),
			NewMatrix(
				0.218045, 0.451127, 0.240601, -0.0451127,
				-0.808270, -1.456766, -0.443609, 0.520676,
				-0.078947, -0.223684, -0.0526315, 0.197368,
				-0.522556, -0.813909, -0.300751, 0.306390,
			),
		},
		{
			"Test 2",
			NewMatrix(
				8.0, -5.0, 9.0, 2.0,
				7.0, 5.0, 6.0, 1.0,
				-6.0, 0.0, 9.0, 6.0,
				-3.0, 0.0, -9.0, -4.0,
			),
			NewMatrix(
				-0.153846, -0.153846, -0.28205, -0.53846,
				-0.076923, 0.123076, 0.0256410, 0.0307692,
				0.3589743, 0.3589743, 0.43590, 0.92308,
				-0.69231, -0.69231, -0.76923, -1.9230769,
			),
		},
		{
			"Test 3",
			NewMatrix(
				9.0, 3.0, 0.0, 9.0,
				-5.0, -2.0, -6.0, -3.0,
				-4.0, 9.0, 6.0, 4.0,
				-7.0, 6.0, 6.0, 2.0,
			),
			NewMatrix(
				-0.0407407, -0.077777777, 0.144444444, -0.222222222,
				-0.07777777, 0.033333333, 0.36666666, -0.33333333,
				-0.02901234, -0.146296, -0.10926, 0.12963,
				0.1777778, 0.0666666, -0.266666, 0.3333333,
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Inverse(); !got.Eq(tt.want) {
				CompareMatrixForTest(got, tt.want, t)
			}
		})
	}
}

func CompareMatrixForTest(got, want *Matrix, t *testing.T) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			got := got.data[i][j]
			want := want.data[i][j]

			if !c.WithinTolerance(got, want, 1e-5) {
				t.Errorf("[%v][%v] Got %v, want %v, ", i, j, got, want)
			}
		}
	}
}

func TestInverseMul(t *testing.T) {
	tests := []struct {
		name string
		a    *Matrix
		b    *Matrix
	}{
		{
			"Test 1",
			NewMatrix(
				3, -9, 7, 3,
				3, -8, 2, -9,
				-4, 4, 4, 1,
				-6, 5, -1, 1,
			),
			NewMatrix(
				8, 2, 2, 2,
				3, -1, 7, 0,
				7, 0, 5, 4,
				6, -2, 0, 5,
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.a.Mul(tt.b)
			bInv := tt.b.Inverse()

			if !c.Mul(bInv).Eq(tt.a) {
				t.Error("Inverse of a times c does not equal a")
			}
		})
	}
}

func TestMulTranslation(t *testing.T) {
	tests := []struct {
		name string
		a    Point
		b    *Matrix
		want Point
	}{
		{
			"Test 1",
			NewPoint(-3, 4, 5),
			NewTranslation(5, -3, 2),
			NewPoint(2, 1, 7),
		},
		{
			"Test 2",
			NewPoint(-3, 4, 5),
			NewTranslation(5, -3, 2).Inverse(),
			NewPoint(-8, 7, 3),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.MulMat(tt.b); !got.Eq(tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMulTranslationVec(t *testing.T) {
	tests := []struct {
		name string
		a    Vec
		b    *Matrix
		want Vec
	}{
		{
			"Test 1",
			NewVec(-3, 4, 5),
			NewTranslation(5, -3, 2),
			NewVec(-3, 4, 5),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.MulMat(tt.b); !got.Eq(tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMulScaling(t *testing.T) {
	tests := []struct {
		name string
		a    Vec
		b    *Matrix
		want Vec
	}{
		{
			"Test 1",
			NewVec(-4, 6, 8),
			NewScaling(2, 3, 4),
			NewVec(-8, 18, 32),
		},
		{
			"Test 2",
			NewVec(-4, 6, 8),
			NewScaling(2, 3, 4).Inverse(),
			NewVec(-2, 2, 2),
		},
		{
			"Test 3",
			NewVec(2, 3, 4),
			NewScaling(-1, 1, 1).Inverse(),
			NewVec(-2, 3, 4),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.MulMat(tt.b); !got.Eq(tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMulScalingPoint(t *testing.T) {
	tests := []struct {
		name string
		a    Point
		b    *Matrix
		want Point
	}{
		{
			"Test 1",
			NewPoint(-4, 6, 8),
			NewScaling(2, 3, 4),
			NewPoint(-8, 18, 32),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.MulMat(tt.b); !got.Eq(tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRotationX(t *testing.T) {
	tests := []struct {
		name string
		a    Vec
		b    *Matrix
		want Vec
	}{
		{
			"Test 1",
			NewVec(0, 1, 0),
			NewRotationX(math.Pi / 4),
			NewVec(0, math.Sqrt(2)/2, math.Sqrt(2)/2),
		},
		{
			"Test 2",
			NewVec(0, 1, 0),
			NewRotationX(math.Pi / 2),
			NewVec(0, 0, 1),
		},
		{
			"Test 3",
			NewVec(0, 1, 0),
			NewRotationX(math.Pi / 4).Inverse(),
			NewVec(0, math.Sqrt(2)/2, -math.Sqrt(2)/2),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.MulMat(tt.b); !got.Eq(tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRotationY(t *testing.T) {
	tests := []struct {
		name string
		a    Vec
		b    *Matrix
		want Vec
	}{
		{
			"Test 1",
			NewVec(0, 0, 1),
			NewRotationY(math.Pi / 4),
			NewVec(math.Sqrt(2)/2, 0, math.Sqrt(2)/2),
		},
		{
			"Test 2",
			NewVec(0, 0, 1),
			NewRotationY(math.Pi / 2),
			NewVec(1, 0, 0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.MulMat(tt.b); !got.Eq(tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRotationZ(t *testing.T) {
	tests := []struct {
		name string
		a    Vec
		b    *Matrix
		want Vec
	}{
		{
			"Test 1",
			NewVec(0, 1, 0),
			NewRotationZ(math.Pi / 4),
			NewVec(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0),
		},
		{
			"Test 2",
			NewVec(0, 1, 0),
			NewRotationZ(math.Pi / 2),
			NewVec(-1, 0, 0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.MulMat(tt.b); !got.Eq(tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewShearing(t *testing.T) {
	tests := []struct {
		name string
		a    Vec
		b    *Matrix
		want Vec
	}{
		{
			"Test 1",
			NewVec(2, 3, 4),
			NewShearing(1, 0, 0, 0, 0, 0),
			NewVec(5, 3, 4),
		},
		{
			"Test 1",
			NewVec(2, 3, 4),
			NewShearing(0, 1, 0, 0, 0, 0),
			NewVec(6, 3, 4),
		},
		{
			"Test 1",
			NewVec(2, 3, 4),
			NewShearing(0, 0, 1, 0, 0, 0),
			NewVec(2, 5, 4),
		},
		{
			"Test 1",
			NewVec(2, 3, 4),
			NewShearing(0, 0, 0, 1, 0, 0),
			NewVec(2, 7, 4),
		},
		{
			"Test 1",
			NewVec(2, 3, 4),
			NewShearing(0, 0, 0, 0, 1, 0),
			NewVec(2, 3, 6),
		},
		{
			"Test 1",
			NewVec(2, 3, 4),
			NewShearing(0, 0, 0, 0, 0, 1),
			NewVec(2, 3, 7),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.MulMat(tt.b); !got.Eq(tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChaining(t *testing.T) {
	rotationX := NewRotationX(math.Pi / 2)
	scaling := NewScaling(5, 5, 5)
	translation := NewTranslation(10, 5, 7)

	tests := []struct {
		name string
		a    Point
		b    *Matrix
		want Point
	}{
		{
			"Test 1",
			NewPoint(1, 0, 1),
			rotationX,
			NewPoint(1, -1, 0),
		},
		{
			"Test 2",
			NewPoint(1, 0, 1).MulMat(rotationX),
			scaling,
			NewPoint(5, -5, 0),
		},
		{
			"Test 3",
			NewPoint(1, 0, 1).MulMat(rotationX).MulMat(scaling),
			translation,
			NewPoint(15, 0, 7),
		},
		{
			"Test 4",
			NewPoint(1, 0, 1),
			translation.Mul(scaling).Mul(rotationX),
			NewPoint(15, 0, 7),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.MulMat(tt.b); !got.Eq(tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFluent(t *testing.T) {
	transform := NewIdentityMatrix().Translate(10, 5, 7).Scale(5, 5, 5).RotateX(math.Pi / 2)

	tests := []struct {
		name string
		a    Point
		b    *Matrix
		want Point
	}{
		{
			"Test 1",
			NewPoint(1, 0, 1),
			transform,
			NewPoint(15, 0, 7),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.MulMat(tt.b); !got.Eq(tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestViewTransform(t *testing.T) {
	testCases := []struct {
		desc string
		from Point
		to   Point
		up   Vec
		want *Matrix
	}{
		{
			desc: "Test 1",
			from: NewPoint(0, 0, 0),
			to:   NewPoint(0, 0, -1),
			up:   NewVec(0, 1, 0),
			want: NewIdentityMatrix(),
		},
		{
			desc: "Test 2",
			from: NewPoint(0, 0, 0),
			to:   NewPoint(0, 0, 1),
			up:   NewVec(0, 1, 0),
			want: NewScaling(-1, 1, -1),
		},
		{
			desc: "Test 3",
			from: NewPoint(0, 0, 8),
			to:   NewPoint(0, 0, 0),
			up:   NewVec(0, 1, 0),
			want: NewTranslation(0, 0, -8),
		},
		{
			desc: "Test 4",
			from: NewPoint(1, 3, 2),
			to:   NewPoint(4, -2, 8),
			up:   NewVec(1, 1, 0),
			want: NewMatrix(
				-0.50709, 0.50709, 0.676122, -2.36643,
				0.76772, 0.60609, 0.121218, -2.82843,
				-0.35857, 0.59761, -0.71714, 0.000000,
				0.00000, 0.00000, 0.00000, 1.000000,
			),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := ViewTransform(tC.from, tC.to, tC.up)

			if !got.Eq(tC.want) {
				CompareMatrixForTest(got, tC.want, t)
			}
		})
	}
}
