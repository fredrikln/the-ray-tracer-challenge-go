package raytracer

type Triangle struct {
	P1     Point
	P2     Point
	P3     Point
	E1     Vec
	E2     Vec
	Normal Vec
	*object
}

func NewTriangle(p1, p2, p3 Point) *Triangle {
	o := newObject()

	t := Triangle{
		P1:     p1,
		P2:     p2,
		P3:     p3,
		E1:     p2.Sub(p1),
		E2:     p3.Sub(p1),
		Normal: (p3.Sub(p1).Cross(p2.Sub(p1))).Norm(),
		object: &o,
	}

	o.parentObject = &t

	return &t
}

func (tr *Triangle) LocalIntersect(r Ray) []Intersection {
	dirCrossE2 := r.Direction.Cross(tr.E2)
	det := tr.E1.Dot(dirCrossE2)

	if WithinTolerance(det, 0, 1e-5) {
		return []Intersection{}
	}

	f := 1 / det

	p1ToOrigin := r.Origin.Sub(tr.P1)
	u := f * p1ToOrigin.Dot(dirCrossE2)

	if u < 0 || u > 1 {
		return []Intersection{}
	}

	originCrossE1 := p1ToOrigin.Cross(tr.E1)

	v := f * r.Direction.Dot(originCrossE1)

	if v < 0 || u+v > 1 {
		return []Intersection{}
	}

	t := f * tr.E2.Dot(originCrossE1)

	return []Intersection{NewIntersectionWithUV(t, tr, u, v)}
}
func (tr *Triangle) LocalNormalAt(p Point, i Intersection) Vec {
	return tr.Normal
}

func (tr *Triangle) Bounds() *BoundingBox {
	bb := NewBoundingBox()

	bb.Add(tr.P1)
	bb.Add(tr.P2)
	bb.Add(tr.P3)

	return bb.Transform(tr.GetTransform())
}
