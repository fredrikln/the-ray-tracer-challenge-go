package geom

type Ray struct {
	Origin    Vec
	Direction Vec
}

func NewRay(origin Vec, direction Vec) Ray {
	return Ray{
		origin,
		direction,
	}
}

func (r Ray) Position(time float64) Vec {
	return r.Origin.Add(r.Direction.Mul(time))
}
