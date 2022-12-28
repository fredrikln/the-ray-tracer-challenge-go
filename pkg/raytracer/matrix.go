package raytracer

import (
	"math"
)

type Matrix struct {
	data [4][4]float64
	inv  *Matrix
}

func NewMatrix(a1, a2, a3, a4, b1, b2, b3, b4, c1, c2, c3, c4, d1, d2, d3, d4 float64) *Matrix {
	return &Matrix{
		data: [4][4]float64{
			{a1, a2, a3, a4},
			{b1, b2, b3, b4},
			{c1, c2, c3, c4},
			{d1, d2, d3, d4},
		},
	}
}

func NewIdentityMatrix() *Matrix {
	return NewMatrix(
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)
}

func NewTranslation(x, y, z float64) *Matrix {
	return NewMatrix(
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	)
}

func NewScaling(x, y, z float64) *Matrix {
	return NewMatrix(
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	)
}

func NewRotationX(radians float64) *Matrix {
	return NewMatrix(
		1, 0, 0, 0,
		0, math.Cos(radians), -math.Sin(radians), 0,
		0, math.Sin(radians), math.Cos(radians), 0,
		0, 0, 0, 1,
	)
}

func NewRotationY(radians float64) *Matrix {
	return NewMatrix(
		math.Cos(radians), 0, math.Sin(radians), 0,
		0, 1, 0, 0,
		-math.Sin(radians), 0, math.Cos(radians), 0,
		0, 0, 0, 1,
	)
}

func NewRotationZ(radians float64) *Matrix {
	return NewMatrix(
		math.Cos(radians), -math.Sin(radians), 0, 0,
		math.Sin(radians), math.Cos(radians), 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)
}

func NewShearing(xy, xz, yx, yz, zx, zy float64) *Matrix {
	return NewMatrix(
		1, xy, xz, 0,
		yx, 1, yz, 0,
		zx, zy, 1, 0,
		0, 0, 0, 1,
	)
}

func ViewTransform(from, to Point, up Vec) *Matrix {
	forward := to.Sub(from).Norm()
	upn := up.Norm()

	left := forward.Cross(upn)
	true_up := left.Cross(forward)

	orientation := NewMatrix(
		left.X, left.Y, left.Z, 0,
		true_up.X, true_up.Y, true_up.Z, 0,
		-forward.X, -forward.Y, -forward.Z, 0,
		0, 0, 0, 1,
	)

	return orientation.Mul(NewTranslation(-from.X, -from.Y, -from.Z))
}

func (a *Matrix) Eq(b *Matrix) bool {
	for i := range a.data {
		for j := range b.data {
			if !WithinTolerance(a.data[i][j], b.data[i][j], 1e-5) {
				return false
			}
		}
	}

	return true
}

func (a *Matrix) Mul(b *Matrix) *Matrix {
	m := &Matrix{}

	for r := 0; r < 4; r += 1 {
		for c := 0; c < 4; c += 1 {
			m.data[r][c] = a.data[r][0]*b.data[0][c] +
				a.data[r][1]*b.data[1][c] +
				a.data[r][2]*b.data[2][c] +
				a.data[r][3]*b.data[3][c]
		}
	}

	return m
}

func (a *Matrix) MulVec(b Vec) Vec {
	return NewVec(
		a.data[0][0]*b.X+a.data[0][1]*b.Y+a.data[0][2]*b.Z,
		a.data[1][0]*b.X+a.data[1][1]*b.Y+a.data[1][2]*b.Z,
		a.data[2][0]*b.X+a.data[2][1]*b.Y+a.data[2][2]*b.Z,
	)
}

func (a *Matrix) MulPoint(b Point) Point {
	return NewPoint(
		a.data[0][0]*b.X+a.data[0][1]*b.Y+a.data[0][2]*b.Z+a.data[0][3],
		a.data[1][0]*b.X+a.data[1][1]*b.Y+a.data[1][2]*b.Z+a.data[1][3],
		a.data[2][0]*b.X+a.data[2][1]*b.Y+a.data[2][2]*b.Z+a.data[2][3],
	)
}

func (a *Matrix) Transpose() *Matrix {
	m := Matrix{}

	for r := 0; r < 4; r += 1 {
		for c := 0; c < 4; c += 1 {
			m.data[c][r] = a.data[r][c]
		}
	}

	return &m
}

func (a *Matrix) Determinant() float64 {
	return determinant4(a.data)
}

func (a *Matrix) Invertible() bool {
	return a.Determinant() != 0.0
}

func (a *Matrix) Inverse() *Matrix {
	if a.inv != nil {
		return a.inv
	}

	if !a.Invertible() {
		panic("Matrix not invertible")
	}

	i := &Matrix{}

	determinant := determinant4(a.data)

	for r := 0; r < 4; r += 1 {
		for c := 0; c < 4; c += 1 {
			co := cofactor4(a.data, r, c)

			i.data[c][r] = co / determinant
		}
	}

	a.inv, i.inv = i, a

	return i
}

func (a *Matrix) Translate(x, y, z float64) *Matrix {
	return a.Mul(NewTranslation(x, y, z))
}
func (a *Matrix) Scale(x, y, z float64) *Matrix {
	return a.Mul(NewScaling(x, y, z))
}
func (a *Matrix) RotateX(radians float64) *Matrix {
	return a.Mul(NewRotationX(radians))
}
func (a *Matrix) RotateY(radians float64) *Matrix {
	return a.Mul(NewRotationY(radians))
}
func (a *Matrix) RotateZ(radians float64) *Matrix {
	return a.Mul(NewRotationZ(radians))
}
func (a *Matrix) Shear(xy, xz, yx, yz, zx, zy float64) *Matrix {
	return a.Mul(NewShearing(xy, xz, yx, yz, zx, zy))
}

// Helpers below

func determinant2(matrix [2][2]float64) float64 {
	return (matrix[0][0] * matrix[1][1]) - (matrix[0][1] * matrix[1][0])
}

func submatrix3(matrix [3][3]float64, row, col int) [2][2]float64 {
	out := [2][2]float64{}

	rowCounter := 0
	for i := 0; i < 3; i++ {
		if i == row {
			continue
		}

		colCounter := 0
		for j := 0; j < 3; j++ {
			if j == col {
				continue
			}

			out[rowCounter][colCounter] = matrix[i][j]

			colCounter += 1
		}

		rowCounter += 1
	}

	return out
}

func minor3(matrix [3][3]float64, row, col int) float64 {
	return determinant2(submatrix3(matrix, row, col))
}

func cofactor3(matrix [3][3]float64, row, col int) float64 {
	out := minor3(matrix, row, col)

	if (row+col)%2 == 0 {
		return out
	}

	return -out
}

func determinant3(matrix [3][3]float64) float64 {
	out := 0.0

	for i := 0; i < 3; i += 1 {
		out += cofactor3(matrix, 0, i) * matrix[0][i]
	}

	return out
}

func submatrix4(matrix [4][4]float64, row, col int) [3][3]float64 {
	out := [3][3]float64{}

	rowCounter := 0
	for i := 0; i < 4; i++ {
		if i == row {
			continue
		}

		colCounter := 0
		for j := 0; j < 4; j++ {
			if j == col {
				continue
			}

			out[rowCounter][colCounter] = matrix[i][j]

			colCounter += 1
		}

		rowCounter += 1
	}

	return out
}

func minor4(matrix [4][4]float64, row, col int) float64 {
	return determinant3(submatrix4(matrix, row, col))
}

func cofactor4(matrix [4][4]float64, row, col int) float64 {
	out := minor4(matrix, row, col)

	if (row+col)%2 == 0 {
		return out
	}

	return -out
}

func determinant4(matrix [4][4]float64) float64 {
	out := 0.0

	for i := 0; i < 4; i += 1 {
		out += cofactor4(matrix, 0, i) * matrix[0][i]
	}

	return out
}
