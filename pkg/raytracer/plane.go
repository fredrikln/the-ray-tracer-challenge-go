package raytracer

import (
	"math"
)

type Plane struct {
	Transform *Matrix
	Material  *Material
	Parent    *Group
}

func NewPlane() *Plane {
	return &Plane{
		Transform: NewIdentityMatrix(),
		Material:  NewMaterial(),
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

func (p *Plane) GetParent() *Group {
	return p.Parent
}
func (p *Plane) SetParent(g *Group) Intersectable {
	p.Parent = g

	return p
}

func (p *Plane) Intersect(worldRay Ray) []Intersection {
	localRay := worldRay.Mul(p.Transform.Inverse())

	if WithinTolerance(0, math.Abs(localRay.Direction.Y), 1e-5) {
		return []Intersection{}
	}

	t := (-localRay.Origin.Y) / localRay.Direction.Y

	return []Intersection{NewIntersection(t, p)}
}

func (p *Plane) NormalAt(worldPoint Point) Vec {
	// objectPoint := p.WorldToObject(worldPoint)
	objectNormal := NewVec(0, 1, 0)
	worldNormal := p.NormalToWorld(objectNormal)

	return worldNormal
}

func (pl *Plane) WorldToObject(p Point) Point {
	parent := pl.GetParent()

	if parent != nil {
		p = parent.WorldToObject(p)
	}

	return pl.GetTransform().Inverse().MulPoint(p)
}

func (pl *Plane) NormalToWorld(n Vec) Vec {
	normal := pl.GetTransform().Inverse().Transpose().MulVec(n).Norm()

	parent := pl.GetParent()

	if parent != nil {
		normal = parent.NormalToWorld(normal)
	}

	return normal
}
