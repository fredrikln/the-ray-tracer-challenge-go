package raytracer

import (
	c "github.com/fredrikln/the-ray-tracer-challenge-go/common"
)

type Point struct {
	X, Y, Z float64
}

func NewPoint(x, y, z float64) Point {
	return Point{
		x,
		y,
		z,
	}
}

func (a Point) AddVec(b Vec) Point {
	return Point{
		a.X + b.X,
		a.Y + b.Y,
		a.Z + b.Z,
	}
}

func (a Point) Sub(b Point) Vec {
	return Vec{
		a.X - b.X,
		a.Y - b.Y,
		a.Z - b.Z,
	}
}

func (a Point) SubVec(b Vec) Point {
	return Point{
		a.X - b.X,
		a.Y - b.Y,
		a.Z - b.Z,
	}
}

func (a Point) MulMat(b *Matrix) Point {
	return b.MulPoint(a)
}

func (a Point) Eq(b Point) bool {
	if !c.WithinTolerance(a.X, b.X, 1e-5) {
		return false
	}
	if !c.WithinTolerance(a.Y, b.Y, 1e-5) {
		return false
	}
	if !c.WithinTolerance(a.Z, b.Z, 1e-5) {
		return false
	}

	return true
}
