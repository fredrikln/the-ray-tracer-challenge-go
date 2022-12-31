package raytracer

type SmoothTriangle struct {
	P1     Point
	P2     Point
	P3     Point
	N1     Vec
	N2     Vec
	N3     Vec
	E1     Vec
	E2     Vec
	Normal Vec
	*object
}

func NewSmoothTriangle(p1, p2, p3 Point, n1, n2, n3 Vec) *SmoothTriangle {
	o := newObject()

	st := SmoothTriangle{
		P1:     p1,
		P2:     p2,
		P3:     p3,
		N1:     n1,
		N2:     n2,
		N3:     n3,
		E1:     p2.Sub(p1),
		E2:     p3.Sub(p1),
		Normal: (p3.Sub(p1).Cross(p2.Sub(p1))).Norm(),
		object: &o,
	}

	o.parentObject = &st

	return &st
}

func (tr *SmoothTriangle) LocalIntersect(r Ray) []Intersection {
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
func (tr *SmoothTriangle) LocalNormalAt(p Point, hit Intersection) Vec {
	a := tr.N2.Mul(*hit.U)
	b := tr.N3.Mul(*hit.V)
	c := tr.N1.Mul(1 - *hit.U - *hit.V)

	return a.Add(b).Add(c).Norm()
}

func (st *SmoothTriangle) Bounds() *BoundingBox {
	bb := NewBoundingBox()

	bb.Add(st.P1)
	bb.Add(st.P2)
	bb.Add(st.P3)

	return bb.Transform(st.GetTransform())
}
