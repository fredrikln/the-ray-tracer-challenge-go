package raytracer

type Triangle struct {
	P1        Point
	P2        Point
	P3        Point
	E1        Vec
	E2        Vec
	Normal    Vec
	Transform *Matrix
	Material  *Material
	Parent    Intersectable
}

func NewTriangle(p1, p2, p3 Point) *Triangle {
	return &Triangle{
		P1:        p1,
		P2:        p2,
		P3:        p3,
		E1:        p2.Sub(p1),
		E2:        p3.Sub(p1),
		Normal:    (p3.Sub(p1).Cross(p2.Sub(p1))).Norm(),
		Transform: NewIdentityMatrix(),
		Material:  NewMaterial(),
	}
}

func (tr *Triangle) Intersect(r Ray) []Intersection {
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
func (tr *Triangle) NormalAt(p Point, i Intersection) Vec {
	return tr.Normal
}

func (tr *Triangle) SetMaterial(m *Material) Intersectable {
	tr.Material = m

	return tr
}
func (tr *Triangle) GetMaterial() *Material {
	return tr.Material
}

func (tr *Triangle) SetTransform(t *Matrix) Intersectable {
	tr.Transform = t

	return tr
}
func (tr *Triangle) GetTransform() *Matrix {
	return tr.Transform
}

func (tr *Triangle) SetParent(g Intersectable) Intersectable {
	tr.Parent = g

	return tr
}
func (tr *Triangle) GetParent() Intersectable {
	return tr.Parent
}

func (tr *Triangle) WorldToObject(p Point) Point {
	parent := tr.GetParent()

	if parent != nil {
		p = parent.WorldToObject(p)
	}

	return tr.GetTransform().Inverse().MulPoint(p)
}

func (tr *Triangle) NormalToWorld(n Vec) Vec {
	normal := tr.GetTransform().Inverse().Transpose().MulVec(n).Norm()

	parent := tr.GetParent()

	if parent != nil {
		normal = parent.NormalToWorld(normal)
	}

	return normal
}

func (tr *Triangle) Bounds() *BoundingBox {
	bb := NewBoundingBox()

	bb.Add(tr.P1)
	bb.Add(tr.P2)
	bb.Add(tr.P3)

	return bb.Transform(tr.Transform)
}

func (tr *Triangle) Divide(int) {
	// Does nothing
}
