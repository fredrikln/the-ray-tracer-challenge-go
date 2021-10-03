package geom

type Ray struct {
	Origin    Point
	Direction Vec
}

func NewRay(origin Point, direction Vec) Ray {
	return Ray{
		origin,
		direction,
	}
}

func (r Ray) Position(time float64) Point {
	return r.Origin.AddVec(r.Direction.Mul(time))
}

func (r Ray) Mul(matrix *Matrix) Ray {
	return Ray{
		r.Origin.MulMat(matrix),
		r.Direction.MulMat(matrix),
	}
}
