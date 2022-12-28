package raytracer

import (
	"math"
)

type Sphere struct {
	Transform *Matrix
	Material  *Material
	Parent    *Group
}

func NewSphere() *Sphere {
	return &Sphere{
		Transform: NewIdentityMatrix(),
		Material:  NewMaterial(),
	}
}

func NewGlassSphere() *Sphere {
	return &Sphere{
		Transform: NewIdentityMatrix(),
		Material:  NewMaterial().SetTransparency(1.0).SetRefractiveIndex(1.5),
	}
}

func (s *Sphere) GetMaterial() *Material {
	return s.Material
}
func (s *Sphere) SetMaterial(m *Material) Intersectable {
	s.Material = m

	return s
}

func (s *Sphere) GetTransform() *Matrix {
	return s.Transform
}
func (s *Sphere) SetTransform(m *Matrix) Intersectable {
	s.Transform = m

	return s
}

func (s *Sphere) GetParent() *Group {
	return s.Parent
}
func (s *Sphere) SetParent(g *Group) Intersectable {
	s.Parent = g

	return s
}

func (s *Sphere) Intersect(worldRay Ray) []Intersection {
	localRay := worldRay.Mul(s.Transform.Inverse())

	sphereToRay := localRay.Origin.Sub(NewPoint(0, 0, 0))

	a := localRay.Direction.Dot(localRay.Direction)
	b := 2 * localRay.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return []Intersection{}
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	i1 := NewIntersection(t1, s)

	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)
	i2 := NewIntersection(t2, s)

	return []Intersection{
		i1,
		i2,
	}
}

func (s *Sphere) NormalAt(worldPoint Point) Vec {
	objectPoint := worldPoint.MulMat(s.Transform.Inverse())

	objectNormal := objectPoint.Sub(NewPoint(0, 0, 0))
	worldNormal := objectNormal.MulMat(s.Transform.Inverse().Transpose())

	return worldNormal.Norm()
}
