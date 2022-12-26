package raytracer

import (
	"math"

	c "github.com/fredrikln/the-ray-tracer-challenge-go/common"
)

type Plane struct {
	Transform *Matrix
	Material  *Material
}

func NewPlane() *Plane {
	return &Plane{
		NewIdentityMatrix(),
		NewMaterial(),
	}
}

func (p *Plane) GetMaterial() *Material {
	return p.Material
}

func (p *Plane) SetMaterial(m *Material) *Plane {
	p.Material = m

	return p
}

func (p *Plane) SetTransform(m *Matrix) *Plane {
	p.Transform = m

	return p
}

func (p *Plane) Intersect(r Ray) []Intersection {
	r2 := r.Mul(p.Transform.Inverse())

	if c.WithinTolerance(0, math.Abs(r2.Direction.Y), 1e-5) {
		return []Intersection{}
	}

	t := (-r2.Origin.Y) / r2.Direction.Y

	return []Intersection{NewIntersection(t, p)}
}

func (p *Plane) NormalAt(point Point) Vec {
	objectNormal := NewVec(0, 1, 0)
	worldNormal := objectNormal.MulMat(p.Transform.Inverse().Transpose())

	return worldNormal
}
