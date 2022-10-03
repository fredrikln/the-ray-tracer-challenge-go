package geom

import (
	"math"

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
	if math.Abs(a.X-b.X) > c.EPSILON {
		return false
	}
	if math.Abs(a.Y-b.Y) > c.EPSILON {
		return false
	}
	if math.Abs(a.Z-b.Z) > c.EPSILON {
		return false
	}

	return true
}
