package raytracer

import (
	"math"
)

type Direction Vec

type Vec struct {
	X, Y, Z float64
}

func NewVec(x, y, z float64) Vec {
	return Vec{
		x, y, z,
	}
}

func (a Vec) Eq(b Vec) bool {
	if !WithinTolerance(a.X, b.X, 1e-5) {
		return false
	}
	if !WithinTolerance(a.Y, b.Y, 1e-5) {
		return false
	}
	if !WithinTolerance(a.Z, b.Z, 1e-5) {
		return false
	}

	return true
}

func (a Vec) Add(b Vec) Vec {
	return Vec{
		a.X + b.X,
		a.Y + b.Y,
		a.Z + b.Z,
	}
}

func (a Vec) Sub(b Vec) Vec {
	return Vec{
		a.X - b.X,
		a.Y - b.Y,
		a.Z - b.Z,
	}
}

func (a Vec) Neg() Vec {
	return Vec{
		-a.X,
		-a.Y,
		-a.Z,
	}
}

func (a Vec) Mul(b float64) Vec {
	return Vec{
		a.X * b,
		a.Y * b,
		a.Z * b,
	}
}

func (a Vec) MulMat(b *Matrix) Vec {
	return b.MulVec(a)
}

func (a Vec) Div(b float64) Vec {
	return Vec{
		a.X / b,
		a.Y / b,
		a.Z / b,
	}
}

func (a Vec) Mag() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

func (a Vec) Norm() Vec {
	mag := a.Mag()

	return Vec{
		a.X / mag,
		a.Y / mag,
		a.Z / mag,
	}
}

func (a Vec) Dot(b Vec) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (a Vec) Cross(b Vec) Vec {
	return Vec{
		a.Y*b.Z - a.Z*b.Y,
		a.Z*b.X - a.X*b.Z,
		a.X*b.Y - a.Y*b.X,
	}
}

func (a Vec) Reflect(n Vec) Vec {
	return a.Sub(n.Mul(2).Mul(a.Dot(n)))
}

func (a Vec) NearZero() bool {
	return WithinTolerance(a.X, 0, 1e-5) && WithinTolerance(a.Y, 0, 1e-5) && WithinTolerance(a.Z, 0, 1e-5)
}

func (a Vec) LengthSquared() float64 {
	return a.X*a.X + a.Y*a.Y + a.Z*a.Z
}
