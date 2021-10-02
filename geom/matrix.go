package geom

import (
	"errors"
	"math"

	c "github.com/fredrikln/the-ray-tracer-challenge-go/common"
)

type Matrix struct {
	data [4][4]float64
}

func NewMatrix(a1, a2, a3, a4, b1, b2, b3, b4, c1, c2, c3, c4, d1, d2, d3, d4 float64) Matrix {
	return Matrix{
		data: [4][4]float64{
			{a1, a2, a3, a4},
			{b1, b2, b3, b4},
			{c1, c2, c3, c4},
			{d1, d2, d3, d4},
		},
	}
}

func NewIdentityMatrix() Matrix {
	return NewMatrix(
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)
}

func (a Matrix) Eq(b Matrix) bool {
	for i := range a.data {
		for j := range b.data {
			if math.Abs(a.data[i][j]-b.data[i][j]) > c.EPSILON {
				return false
			}
		}
	}

	return true
}

func (a Matrix) Mul(b Matrix) Matrix {
	m := NewMatrix(
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	)

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

func (a Matrix) MulVec(b Vec) Vec {
	return NewVec(
		a.data[0][0]*b.X+a.data[0][1]*b.Y+a.data[0][2]*b.Z+a.data[0][3],
		a.data[1][0]*b.X+a.data[1][1]*b.Y+a.data[1][2]*b.Z+a.data[1][3],
		a.data[2][0]*b.X+a.data[2][1]*b.Y+a.data[2][2]*b.Z+a.data[2][3],
	)
}

func (a Matrix) Transpose() Matrix {
	m := NewMatrix(
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	)

	for r := 0; r < 4; r += 1 {
		for c := 0; c < 4; c += 1 {
			m.data[c][r] = a.data[r][c]
		}
	}

	return m
}

func (a Matrix) Determinant() float64 {
	return determinant4(a.data)
}

func (a Matrix) Invertible() bool {
	return a.Determinant() != 0.0
}

func (a Matrix) Inverse() (Matrix, error) {
	if !a.Invertible() {
		return Matrix{}, errors.New("Matrix is not invertible")
	}

	out := NewIdentityMatrix()

	determinant := determinant4(a.data)

	for r := 0; r < 4; r += 1 {
		for c := 0; c < 4; c += 1 {
			co := cofactor4(a.data, r, c)

			out.data[c][r] = co / determinant
		}
	}

	return out, nil // Todo: Fix
}

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
