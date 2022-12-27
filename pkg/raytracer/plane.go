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

func (p *Plane) SetMaterial(m *Material) Intersectable {
	p.Material = m

	return p
}

func (p *Plane) GetTransform() *Matrix {
	return p.Transform
}

func (p *Plane) SetTransform(m *Matrix) Intersectable {
	p.Transform = m

	return p
}

func (p *Plane) Intersect(worldRay Ray) []Intersection {
	localRay := worldRay.Mul(p.Transform.Inverse())

	if c.WithinTolerance(0, math.Abs(localRay.Direction.Y), 1e-5) {
		return []Intersection{}
	}

	t := (-localRay.Origin.Y) / localRay.Direction.Y

	return []Intersection{NewIntersection(t, p)}
}

func (p *Plane) NormalAt(worldPoint Point) Vec {
	objectNormal := NewVec(0, 1, 0)
	worldNormal := objectNormal.MulMat(p.Transform.Inverse().Transpose())

	return worldNormal
}
