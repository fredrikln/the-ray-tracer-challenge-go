package geom

import "math"

type Vec struct {
	X float64
	Y float64
	Z float64
}

func NewVec(x, y, z float64) Vec {
	return Vec{
		x, y, z,
	}
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
		0 - a.X,
		0 - a.Y,
		0 - a.Z,
	}
}

func (a Vec) Mul(b float64) Vec {
	return Vec{
		a.X * b,
		a.Y * b,
		a.Z * b,
	}
}

func (a Vec) Div(b float64) Vec {
	return Vec{
		a.X / b,
		a.Y / b,
		a.Z / b,
	}
}

func (a Vec) Mag() float64 {
	return math.Sqrt(math.Pow(a.X, 2.0) + math.Pow(a.Y, 2.0) + math.Pow(a.Z, 2.0))
}

func (a Vec) Norm() Vec {
	mag := a.Mag()

	return Vec{
		a.X / mag,
		a.Y / mag,
		a.Z / mag,
	}
}
