package raytracer

type SmoothTriangle struct {
	P1        Point
	P2        Point
	P3        Point
	N1        Vec
	N2        Vec
	N3        Vec
	E1        Vec
	E2        Vec
	Normal    Vec
	Transform *Matrix
	Material  *Material
	Parent    Intersectable
}

func NewSmoothTriangle(p1, p2, p3 Point, n1, n2, n3 Vec) *SmoothTriangle {
	return &SmoothTriangle{
		P1:        p1,
		P2:        p2,
		P3:        p3,
		N1:        n1,
		N2:        n2,
		N3:        n3,
		E1:        p2.Sub(p1),
		E2:        p3.Sub(p1),
		Normal:    (p3.Sub(p1).Cross(p2.Sub(p1))).Norm(),
		Transform: NewIdentityMatrix(),
		Material:  NewMaterial(),
	}
}

func (tr *SmoothTriangle) Intersect(r Ray) []Intersection {
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
func (tr *SmoothTriangle) NormalAt(p Point, hit Intersection) Vec {
	a := tr.N2.Mul(*hit.U)
	b := tr.N3.Mul(*hit.V)
	c := tr.N1.Mul(1 - *hit.U - *hit.V)

	return a.Add(b).Add(c).Norm()
}

func (tr *SmoothTriangle) SetMaterial(m *Material) Intersectable {
	tr.Material = m

	return tr
}
func (tr *SmoothTriangle) GetMaterial() *Material {
	return tr.Material
}

func (tr *SmoothTriangle) SetTransform(t *Matrix) Intersectable {
	tr.Transform = t

	return tr
}
func (tr *SmoothTriangle) GetTransform() *Matrix {
	return tr.Transform
}

func (tr *SmoothTriangle) SetParent(p Intersectable) Intersectable {
	tr.Parent = p

	return tr
}
func (tr *SmoothTriangle) GetParent() Intersectable {
	return tr.Parent
}

func (tr *SmoothTriangle) WorldToObject(p Point) Point {
	parent := tr.GetParent()

	if parent != nil {
		p = parent.WorldToObject(p)
	}

	return tr.GetTransform().Inverse().MulPoint(p)
}

func (tr *SmoothTriangle) NormalToWorld(n Vec) Vec {
	normal := tr.GetTransform().Inverse().Transpose().MulVec(n).Norm()

	parent := tr.GetParent()

	if parent != nil {
		normal = parent.NormalToWorld(normal)
	}

	return normal
}
