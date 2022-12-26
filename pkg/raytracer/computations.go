package raytracer

type Computations struct {
	Time       float64
	Object     *Intersectable
	Point      Point
	Eyev       Vec
	Normalv    Vec
	Inside     bool
	OverPoint  Point
	UnderPoint Point
}

func PrepareComputations(i Intersection, r Ray) Computations {
	p := r.Position(i.Time)

	eyev := r.Direction.Neg()
	normalv := (*i.Object).NormalAt(p)
	inside := normalv.Dot(eyev) < 0

	if inside {
		normalv = normalv.Neg()
	}

	c := Computations{
		Time:       i.Time,
		Object:     i.Object,
		Point:      p,
		Eyev:       eyev,
		Normalv:    normalv,
		Inside:     inside,
		OverPoint:  p.AddVec(normalv.Mul((1e-5) / 2)),
		UnderPoint: p.AddVec(normalv.Mul(-(1e-5) / 2)),
	}

	return c
}
