package geom

import (
	"math"

	c "github.com/fredrikln/the-ray-tracer-challenge-go/common"
)

type Vec struct {
	X float64
	Y float64
	Z float64
}

func NewVec(x, y, z float64) *Vec {
	return &Vec{
		x, y, z,
	}
}

func (a *Vec) Eq(b *Vec) bool {
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

func (a *Vec) Add(b *Vec) *Vec {
	return &Vec{
		a.X + b.X,
		a.Y + b.Y,
		a.Z + b.Z,
	}
}

func (a *Vec) Sub(b *Vec) *Vec {
	return &Vec{
		a.X - b.X,
		a.Y - b.Y,
		a.Z - b.Z,
	}
}

func (a *Vec) Neg() *Vec {
	return &Vec{
		-a.X,
		-a.Y,
		-a.Z,
	}
}

func (a *Vec) Mul(b float64) *Vec {
	return &Vec{
		a.X * b,
		a.Y * b,
		a.Z * b,
	}
}

func (a *Vec) MulMat(b *Matrix) *Vec {
	return b.MulVec(a)
}

func (a *Vec) Div(b float64) *Vec {
	return &Vec{
		a.X / b,
		a.Y / b,
		a.Z / b,
	}
}

func (a *Vec) Mag() float64 {
	return math.Sqrt(math.Pow(a.X, 2.0) + math.Pow(a.Y, 2.0) + math.Pow(a.Z, 2.0))
}

func (a *Vec) Norm() *Vec {
	mag := a.Mag()

	return &Vec{
		a.X / mag,
		a.Y / mag,
		a.Z / mag,
	}
}

func (a *Vec) Dot(b *Vec) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (a *Vec) Cross(b *Vec) *Vec {
	return &Vec{
		a.Y*b.Z - a.Z*b.Y,
		a.Z*b.X - a.X*b.Z,
		a.X*b.Y - a.Y*b.X,
	}
}
